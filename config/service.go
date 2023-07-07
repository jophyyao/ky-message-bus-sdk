package config

type ServiceConfig struct {
	Host      string `json:"host"`
	Port      int64  `json:"port"`
	UrlPrefix string `json:"url_prefix"`
}

func InitConfig() *ServiceConfig {
	GetEnv()

	cf := &ServiceConfig{}
	if Env == "dev" {
		cf.Host = "http://localhost"
		cf.Port = 5280
		cf.UrlPrefix = "api/v1"
	} else {
		// prod
	}
	return cf
}
