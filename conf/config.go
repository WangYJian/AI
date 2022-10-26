package conf

type Config struct {
	Settings Settings
}

type Settings struct {
	ApiKey 	   string `json:"ApiKey"`
	SecreatKey string `json:"SecreatKey"`
}
