package config

type (
	Configuration struct {
		PostgreSqlDbSettings DbSettings `json:"PostgreSqlDbSettings"`
	}

	DbSettings struct {
		Host string `json:"host"`
		User string `json:"user"`
		Password string `json:"password"`
		DbName string `json:"dbname"`
		Port string `json:"port"`
		SslMode string `json:"sslmode"`
	}
)