package entities

type ConfigFunc func() *Config

type Config struct {
	// SqlSettings          SqlSettings
	// ConfigSettings       ConfigSettings
	// NoSqlSettings        NoSqlSettings
	// GeneralSettings      GeneralSettings
	// LocalizationSettings LocalizationSettings
	// EmailSettings        EmailSettings
	// ServiceSettings      ServiceSettings
	// MailChimp            MailChimp
	Urls        Urls
	Environment string
}

type Urls struct {
	Nats     string
	Memcache string
	MySQL    string
	NoSQL    string
}
