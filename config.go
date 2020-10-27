package nmc_typhoon_db_client

type DatabaseConfig struct {
	Host         string
	DatabaseName string `yaml:"database_name"`
	TableName    string `yaml:"table_name"`

	Auth struct {
		User     string
		Password string
	}
}

type Config struct {
	Database DatabaseConfig
}
