package config

type Environment struct {
	Env     string
	ApiPort string
	ApiUrl  string
	DB      Mysql
}

type Mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
