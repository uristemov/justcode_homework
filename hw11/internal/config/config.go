package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTP  HTTPServerConfig `yaml:"http"`
	DB    DBConfig         `yaml:"db"`
	Token TokenConfig      `yaml:"token"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name" mapstructure:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"db_password" mapstructure:"db_password"`
}

type TokenConfig struct {
	SecretKey  string        `yaml:"secret_key" mapstructure:"secret_key"`
	TimeToLive time.Duration `yaml:"time_to_live" mapstructure:"time_to_live"`
}

type HTTPServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

func NewCleanEnvConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	return cfg, nil
}

func NewViperConfig() (*Config, error) {
	cfg := Config{}

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./hw11/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
