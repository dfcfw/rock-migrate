package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flow struct {
	ID            primitive.ObjectID `json:"id"             bson:"_id,omitempty"`  // 数据库 ID
	Addr          string             `json:"addr"           bson:"addr"`           // 客户端IP
	App           string             `json:"app"            bson:"app"`            // 应用分组
	AppKey        string             `json:"app_key"        bson:"app_key"`        // ?
	ASI           string             `json:"asi"            bson:"asi"`            // 大数据设备指纹	大数据指纹埋点字段
	Body          string             `json:"body"           bson:"body"`           // 请求体	需要手动开启才会记录，默认不记录
	CDN           string             `json:"cdn"            bson:"cdn"`            // CDN 回源标识，对接cdn的header头
	Cluster       string             `json:"cluster"        bson:"cluster"`        // WAF 集群名称
	Conn          string             `json:"conn"           bson:"conn"`           // 当前连接总数
	ContentLength int                `json:"content-length" bson:"content-length"` // 请求内容长度
	ContentType   string             `json:"content_type"   bson:"content_type"`   // 请求数据类型
	Cookie        string             `json:"cookie"         bson:"cookie"`         // 用户请求cookie，默认不记录
	CToken        string             `json:"ctoken"         bson:"ctoken"`         // 通行证登录ctoken，hash存储
	CustomerNo    string             `json:"customerno"     bson:"customerno"`     // customer number	基金用户标识 ID
	Debug         string             `json:"debug"          bson:"debug"`          // 调试信息，需要手动开启
	Domain        string             `json:"domain"         bson:"domain"`         // 原始主机头
	EmChl         string             `json:"em-chl"         bson:"em-chl"`         // 大数据设备指纹	大数据指纹埋点字段
	EmOS          string             `json:"em-os"          bson:"em-os"`          // 大数据设备指纹	大数据指纹埋点字段
	EmPkg         string             `json:"em-pkg"         bson:"em-pkg"`         // 大数据设备指纹	大数据指纹埋点字段
	EmVer         string             `json:"em-ver"         bson:"em-ver"`         // 大数据设备指纹	大数据指纹埋点字段
	EmpID         string             `json:"empId"          bson:"empId"`          // 内部员工工号
	GToken        string             `json:"gtoken"         bson:"gtoken"`         // 通行证登录 gtoken
	Host          string             `json:"host"           bson:"host"`           // 主机头
	Inirurl       string             `json:"inirurl"        bson:"inirurl"`        // 大数据设备指纹	大数据指纹埋点字段
	JA3           string             `json:"ja3"            bson:"ja3"`            // SSL JA3 指纹
	Local         string             `json:"local"          bson:"local"`          // 本地时间
	Method        string             `json:"method"         bson:"method"`         // 请求方法
	Pass          string             `json:"pass"           bson:"pass"`           // 白名单
	PI            string             `json:"pi"             bson:"pi"`             // 通行证认证字段
	PluginChain   string             `json:"plugin_chain"   bson:"plugin_chain"`   // 插件调用链
	Port          string             `json:"port"           bson:"port"`           // 客户端请求端口	客户端连接请求的端口（不一定准确）
	PPAddr        string             `json:"ppaddr"         bson:"ppaddr"`         // proxy_protocol_addr	原始客户端的IP地址（开启tcp_proxy）
	Proto         string             `json:"proto"          bson:"proto"`          // ?
	Proxy         string             `json:"proxy"          bson:"proxy"`          // 请求的WAF配置id，格式：代理名称_应用分组
	PSI           string             `json:"psi"            bson:"psi"`            // 大数据设备指纹	大数据指纹埋点字段
	PVI           string             `json:"pvi"            bson:"pvi"`            // 设备指纹字段	从cookie字段中获取
	Query         string             `json:"query"          bson:"query"`          // query 请求参数
	RawHeader     string             `json:"raw_header"     bson:"raw_header"`     // 原始请求 Header
	RawURI        string             `json:"raw_uri"        bson:"raw_uri"`        // 原始请求路径	未改写之前的请求路径
	Ref           string             `json:"ref"            bson:"ref"`            // Referer
	Remote        string             `json:"remote"         bson:"remote"`         // remote addr字段
	Reqs          string             `json:"reqs"           bson:"reqs"`           // 同一个tcp连接中的请求个数
	Response      string             `json:"response"       bson:"response"`       // 响应报文
	Risk          string             `json:"risk"           bson:"risk"`           // 请求风险
	RT            float64            `json:"rt"             bson:"rt"`             // 请求耗时	从WAF接收到用户请求到处理完成的时间，包含后端响应时间，单位:ms
	Rule          string             `json:"rule"           bson:"rule"`           // 拦截规则	rule为down时，表示所有后端节点均不在线，请求被转发到peer(返回403)
	SAddr         string             `json:"saddr"          bson:"saddr"`          // Server Addr	处理该请求的WAF节点地址
	Scheme        string             `json:"scheme"         bson:"scheme"`         // http 请求协议：http/https
	Security      string             `json:"security"       bson:"security"`       // 安全策略返回内容	部分安全插件返回的详细命中信息
	SI            string             `json:"si"             bson:"si"`             // 大数据设备指纹	大数据指纹埋点字段
	Size          int                `json:"size"           bson:"size"`           // 响应报文大小
	SN            int                `json:"sn"             bson:"sn"`             // 大数据设备指纹	大数据指纹埋点字段
	SP            string             `json:"sp"             bson:"sp"`             // 大数据设备指纹	大数据指纹埋点字段
	SPort         string             `json:"sport"          bson:"sport"`          // Server Port 即：WAF服务器端口
	SSLProtocol   string             `json:"ssl_protocol"   bson:"ssl_protocol"`   // SSL 协议
	Status        int                `json:"status"         bson:"status"`         // 响应状态码
	Time          int64              `json:"time"           bson:"time"`           // 时间戳
	Trace         string             `json:"trace"          bson:"trace"`          // ?
	UA            string             `json:"ua"             bson:"ua"`             // User-Agent字段
	UAddr         string             `json:"uaddr"          bson:"uaddr"`          // upstream addr	后端业务服务器
	Uidal         string             `json:"uidal"          bson:"uidal"`          // 大数据设备指纹	大数据指纹埋点字段
	Uname         string             `json:"uname"          bson:"uname"`          // upstream 名称，后端服务器分组名称（上游名）
	URI           string             `json:"uri"            bson:"uri"`            // 改写后的请求路径
	URT           float64            `json:"urt"            bson:"urt"`            // upstream 响应时间，后端服务器响应时间，单位:ms
	UserID        string             `json:"userid"         bson:"userid"`         // 请求的用户 ID，在基金活动中用到
	UToken        string             `json:"utoken"         bson:"utoken"`         // 通行证登录 utoken
	WARN          string             `json:"warn"           bson:"warn"`           // 请求风险字段，命中的标记模式安全规则
	XFF           string             `json:"xff"            bson:"xff"`            // X-Forwarded-For字段
}
