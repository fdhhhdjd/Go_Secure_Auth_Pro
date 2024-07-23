package models

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	Gmail    GmailConfig
	Telegram TelegramConfig
	RabbitMQ RabbitMQConfig
	Cors     CorsConfig
}

type CorsConfig struct {
	AllowedOrigins []string
}

type ServerConfig struct {
	Host         string
	Port         string
	PortFrontend string
	KeyPassword  string
}

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

type CacheConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

type GmailConfig struct {
	Host     string
	Port     string
	Password string
	Service  string
	Mail     string
}

type TelegramConfig struct {
	BotToken string
	ChatID   string
}

type RabbitMQConfig struct {
	Username string
	Password string
	URL      string
}
