package config

type Basecfg struct {
	RunMode  string
	HttpPort int
}

type DBcfg struct {
	User   string
	Pwd    string
	Host   string
	Port   int
	DBName string
}

type Config struct {
	Base  Basecfg
	Mongo DBcfg
}
