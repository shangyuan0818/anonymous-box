package config

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"config",
		fx.Invoke(LoadDotEnv),
		fx.Provide(NewEnvConfig("DB", DatabaseEnv{})),
		fx.Provide(NewEnvConfig("CONSUL", ConsulEnv{})),
		fx.Provide(NewEnvConfig("HASHIDS", HashidsEnv{})),
		fx.Provide(NewEnvConfig("JWT", JwtEnv{})),
		fx.Provide(NewEnvConfig("MQ", MqEnv{})),
		fx.Provide(NewEnvConfig("REDIS", RedisEnv{})),
		fx.Provide(NewEnvConfig("TRACE", TraceEnv{})),
		fx.Provide(NewEnvConfig("LOGGER", LoggerEnv{})),
		fx.Provide(NewEnvConfig("SERVICE", ServiceEnv{})),
		fx.Provide(NewEnvConfig("EMAIL", EmailEnv{})),
		fx.Provide(NewEnvConfig("SERVER", ServerEnv{})),
	)
}
