package config

type Environment struct {
	Env       string
	ApiPort   string
	ApiUrl    string
	SecretKey string
	DB        Postgres
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
