package config

const (
	DefaultPort = "8080"
)

/* configuration definition */
type ConfigDefn struct {
	Addr     string
	Port     string
	Security SecurityConfigDefn
	DescPath string
	Verbose  bool
	Quiet    bool
}

var Config = ConfigDefn{Addr: "", Port: DefaultPort}
