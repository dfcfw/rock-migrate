package launch

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/business/service"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"github.com/dfcfw/rock-migrate/handler/middle"
	"github.com/dfcfw/rock-migrate/handler/restapi"
	"github.com/dfcfw/rock-migrate/handler/shipx"
	"github.com/dfcfw/rock-migrate/library/cronv3"
	"github.com/dfcfw/rock-migrate/library/dynwriter"
	"github.com/dfcfw/rock-migrate/logger"
	"github.com/dfcfw/rock-migrate/profile"
	"github.com/xgfone/ship/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/connstring"
)

func Run(ctx context.Context, fp string) error {
	cfg, err := profile.JSONC(fp)
	if err != nil {
		return err
	}

	return Exec(ctx, cfg)
}

func Exec(ctx context.Context, cfg *profile.Config) error {
	logWriter := dynwriter.New()
	logCfg := cfg.Logger
	if logCfg.Console {
		logWriter.Attach(os.Stdout)
	}
	if lumber := logCfg.Lumber(); lumber != nil {
		//goland:noinspection GoUnhandledErrorResult
		defer lumber.Close()
		logWriter.Attach(lumber)
	}

	logLevel := new(slog.LevelVar)
	if err := logLevel.UnmarshalText([]byte(logCfg.Level)); err != nil {
		logLevel.Set(slog.LevelDebug)
	}
	logOpt := &slog.HandlerOptions{AddSource: true, Level: logLevel}
	logHandler := slog.NewJSONHandler(logWriter, logOpt)
	log := slog.New(logHandler)

	// 连接数据库
	mongoLogOpt := options.Logger().
		SetSink(logger.NewSink(logHandler, 13)).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)
	var sourceDB, targetDB *mongo.Database
	dbCfg := cfg.Database
	{
		log.Info("开始连接源数据库")
		mongoURI := dbCfg.Source
		mongoURL, err := connstring.ParseAndValidate(mongoURI)
		if err != nil {
			return err
		}
		mongoOpt := options.Client().
			ApplyURI(mongoURI).
			SetLoggerOptions(mongoLogOpt)
		cli, err := mongo.Connect(mongoOpt)
		if err != nil {
			log.Error("连接源数据库出错", slog.Any("error", err))
			return err
		}
		defer disconnectDB(cli, 10*time.Second)
		sourceDB = cli.Database(mongoURL.Database)
		log.Info("连接源数据库成功")
	}
	{
		log.Info("开始连接目的数据库")
		mongoURI := dbCfg.Target
		mongoURL, err := connstring.ParseAndValidate(mongoURI)
		if err != nil {
			return err
		}
		mongoOpt := options.Client().
			ApplyURI(mongoURI).
			SetLoggerOptions(mongoLogOpt)
		cli, err := mongo.Connect(mongoOpt)
		if err != nil {
			log.Error("连接目的数据库出错", slog.Any("error", err))
			return err
		}
		defer disconnectDB(cli, 10*time.Second)
		targetDB = cli.Database(mongoURL.Database)
		log.Info("连接目的数据库成功")
	}

	cronLog := slog.New(logger.Skip(logHandler, 5))
	crontab := cronv3.New(cronLog)
	crontab.Start()
	defer crontab.Stop()

	sourceThreatIP := repository.NewThreatIP(sourceDB)
	targetThreatIP := repository.NewThreatIP(targetDB)
	sourceStatus := repository.NewStatus(sourceDB)
	targetStatus := repository.NewStatus(targetDB)

	indexes := []repository.IndexCreator{
		targetThreatIP,
		targetStatus,
	}
	log.Info("开始创建索引")
	if err := repository.CreateIndex(ctx, indexes); err != nil {
		return err
	}
	log.Info("开始创建完毕")

	threatIPBiz := service.NewThreatIP(sourceThreatIP, targetThreatIP, log)
	statusBiz := service.NewStatus(sourceStatus, targetStatus, log)

	executor := execute.New(crontab, log)
	executor.Add(ctx, threatIPBiz, statusBiz)

	logBiz := service.NewLog(logWriter, log)
	shipRoutes := []shipx.RouteRegister{
		restapi.NewLog(logBiz),
	}

	srvCfg := cfg.Server
	httpHandler := ship.Default()
	httpHandler.Use(middle.AccessLog)
	if dir := srvCfg.Static; dir != "" {
		httpHandler.Route("/").Static(dir)
	}
	httpHandler.Logger = logger.NewShip(logHandler, 6)

	baseGroup := httpHandler.Group("/api")
	if err := shipx.RegisterRoutes(baseGroup, shipRoutes); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    srvCfg.Addr,
		Handler: httpHandler,
	}
	if vhosts := srvCfg.Vhosts; len(vhosts) != 0 {
		manager := ship.NewHostManager(nil)
		for _, vhost := range vhosts {
			if _, err := manager.AddHost(vhost, httpHandler); err != nil {
				return err
			}
		}
		srv.Handler = ship.NewHostManagerHandler(manager)
	}
	errs := make(chan error, 1)
	go serveHTTP(errs, srv)

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errs:
	}

	return err
}

func serveHTTP(errs chan<- error, srv *http.Server) {
	errs <- srv.ListenAndServe()
}

func disconnectDB(cli *mongo.Client, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_ = cli.Disconnect(ctx)
}
