package config

type DatabaseEnv struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	User     string `default:"postgres"`
	Password string `default:"postgres"`
	Database string `default:"postgres"`
	SSLMode  string `default:"disable"`
	TimeZone string `default:"Asia/Shanghai" envconfig:"TZ"`
}

type ConsulEnv struct {
	Addr   string `default:"localhost:8500"`
	Scheme string `default:"http"`
	Token  string `default:""`
}

type HashidsEnv struct {
	Salt      string `default:"salt"`
	MinLength int    `default:"8" envconfig:"HASHIDS_MIN_LENGTH"`
}

type JwtEnv struct {
	Secret string `default:"secret"`
	Expire int    `default:"3600"`
}

type LoggerEnv struct {
	Level     string `default:"info"`
	Formatter string `default:"text"`
}

type MqEnv struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5672"`
	User     string `default:"guest"`
	Password string `default:"guest"`
	Vhost    string `default:"/"`
}

type RedisEnv struct {
	Network  string `default:"tcp"`
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Username string `default:""`
	Password string `default:""`
	DB       int    `default:"0"`
}

type TraceEnv struct {
	Endpoint string `default:"http://localhost:14268/api/traces"`
	Exporter string `default:"jaeger"`
}

type ServiceEnv struct {
	Network string `default:"tcp"`
	Address string `default:"0.0.0.0:8888"`
}
