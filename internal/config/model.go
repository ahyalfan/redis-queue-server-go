package config

type Config struct {
	Server Server
	Mail   Email
	Redis  Redis
}

type Server struct {
	Port string
	Host string
}

type Email struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Redis struct {
	Addr     string
	Password string
	DB       string
}
