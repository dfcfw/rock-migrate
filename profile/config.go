package profile

type Config struct {
	Active   string   `json:"active"`
	Server   Server   `json:"server"`
	Database Database `json:"database"`
	Logger   Logger   `json:"logger"`
}

type Database struct {
	Source string `json:"source" validate:"required"`
	Target string `json:"target" validate:"required"`
}

type Server struct {
	Addr   string `json:"addr"   validate:"lte=100"`
	Cert   string `json:"cert"   validate:"lte=255"`
	Pkey   string `json:"pkey"   validate:"lte=255"`
	Static string `json:"static" validate:"lte=255"`
}
