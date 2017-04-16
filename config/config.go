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
}

var Config = ConfigDefn{Addr: "", Port: DefaultPort}
