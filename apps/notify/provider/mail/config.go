package mail

func NewDefaultConfig() *Config {
	return &Config{}
}

type Config struct {
	Host     string `bson:"host" json:"host" env:"MAIL_HOST"`
	Port     int    `bson:"port" json:"port" env:"MAIL_HORT"`
	Username string `bson:"username" json:"username" env:"MAIL_USERNAME"`
	Password string `bson:"password" json:"password" env:"MAIL_PASSWORD"`
}
