package config

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		GRPC  `yaml:"grpc"`
		Log   `yaml:"logger"`
		MySQL `yaml:"mysql"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// gRPC
	GRPC struct {
		Port string `env-required:"true" yaml:"port" env:"GRPC_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// MySQL -.
	MySQL struct {
		URL          string `env-required:"true"                 env:"MYSQL_URL"`
		MaxIdleConns int    `env-required:"true" yaml:"max_idle_conns" env:"MAX_IDLE_CONNS"`
		MaxOpenConns int    `env-required:"true" yaml:"max_open_conns" env:"MAX_OPEN_CONNS"`
	}
)
