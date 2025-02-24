package config

type Config struct {
	AddrHttp string
	DbConfig DBConfig
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
