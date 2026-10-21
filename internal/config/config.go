package config

// HTTP server configurations
type ServerConfig struct {
	Port            string
	ReadTimeoutSec  int
	WriteTimeoutSec int
}

// Database configurations
type DatabaseConfig struct {
	DSN string // PostgreSQL Data Source Name
}

// JWT configurations
type JWTConfig struct {
	Secret        string
	ExpiryMinutes int
}

// Config holds all application configs, read from environment variables
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}
