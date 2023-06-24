package config

// Connection is the database connection config
type Connection struct {
	URL     string  `yaml:"url"`
	Logging Logging `yaml:"logging"`
}
