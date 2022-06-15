package common

type Config struct {
	Port       string             `toml:"port" default:"80" env:"PORT"`
	DB         DatabaseConnection `toml:"DB" env:"DATABASE"`
	CoinMarket CoinMarket         `toml:"CoinMarket"`
}

type DatabaseConnection struct {
	DriverName     string `toml:"driver_name" env:"DRIVER_NAME"`
	DataSourceName string `toml:"data_source_name" env:"DATA_SOURCE_NAME"`
	RootSourceName string `toml:"root_source_name" env:"ROOT_SOURCE_NAME"`
	MaxIdleConns   int    `toml:"max_idle_conns" env:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns   int    `toml:"max_open_conns" env:"DB_MAX_OPEN_CONNS"`
}

type CoinMarket struct {
	Host     string `toml:"host"`
	ClientID string `toml:"client_id"`
}
