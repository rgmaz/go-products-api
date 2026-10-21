package config

const (
	jwtSecretEnv = "JWT_SECRET"
	jwtExpiryEnv = "JWT_EXPIRY_MINUTES"

	dsnEnv = "DATABASE_URL"

	readTimeoutEnv  = "SERVER_READ_TIMEOUT_SEC"
	writeTimeoutEnv = "SERVER_WRITE_TIMEOUT_SEC"
	portEnv         = "SERVER_PORT"
)

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

// Load reads environment variables and returns a populated Config struct.
func Load() (*Config, error) {

	// JWT configurations
	jwtSecret, err := readRequiredString(jwtSecretEnv)
	if err != nil {
		return nil, err
	}

	jwtExpiry, err := readRequiredInt(jwtExpiryEnv)
	if err != nil {
		return nil, err
	}

	jwtConf := JWTConfig{
		Secret:        jwtSecret,
		ExpiryMinutes: jwtExpiry,
	}

	// Database configurations
	dsn, err := readRequiredString(dsnEnv)
	if err != nil {
		return nil, err
	}

	dbConf := DatabaseConfig{
		DSN: dsn,
	}

	//Server configurations
	readTimeout, err := readOptionalInt(readTimeoutEnv, 10)
	if err != nil {
		return nil, err
	}

	writeTimeout, err := readOptionalInt(writeTimeoutEnv, 30)
	if err != nil {
		return nil, err
	}

	port := readOptionalString(portEnv, "8080")

	serverConf := ServerConfig{
		ReadTimeoutSec:  readTimeout,
		WriteTimeoutSec: writeTimeout,
		Port:            port,
	}

	return &Config{
		Server:   serverConf,
		Database: dbConf,
		JWT:      jwtConf,
	}, nil
}
