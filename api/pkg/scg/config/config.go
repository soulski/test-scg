package config

type MessageConfig struct {
	Secret      string
	AccessToken string
}

type PlaceConfig struct {
	ApiKey string
}

type Config struct {
	Message *MessageConfig
	Place   *PlaceConfig
}
