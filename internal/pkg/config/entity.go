package config

type HttpConfig struct {
	BindAddress string
	Mode        string
	Prefix      string
}

type CommonConfig struct {
	LogLevel int8
}

type RedisConfig struct {
	RedisAddresses string
	Password       string
	MasterName     string
}
