package config

import (
	"fmt"
	"log"


	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		HTTPPort int
		GRPCPort int
	}

	Database struct {
		PostgresDSN string
		RedisAddr   string
	}

	Kafka struct {
		Brokers      []string
		PaymentTopic string
	}

	Log struct {
		Level string
	}
}

var AppConfig *Config

func LoadConfig(env string) {
	v := viper.New()

	// Set defaults
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AutomaticEnv()

	// Load base config
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading base config: %v", err)
	}

	// Merge environment-specific overrides
	if env != "" {
		v.SetConfigName(fmt.Sprintf("app.%s", env))
		v.MergeInConfig() // Merge overrides
	}

	// Unmarshal to struct
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	AppConfig = &c
	fmt.Printf("Config loaded for env: %s\n", env)
}
