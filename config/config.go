package config

type MysqlConf struct {
	User     string
	Pass     string
	Host     string
	Port     int
	Database string
}

type Config struct {
	MysqlConf *MysqlConf
}
