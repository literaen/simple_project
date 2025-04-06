package config

type DB_CREDS struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Name     string `env:"POSTGRES_DB"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
}

type REDIS_CREDS struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
}
