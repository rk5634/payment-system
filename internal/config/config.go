package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		HTTPPort int
		GRPCPort int
	}

	Database struct {
		PostgresDSN    string
		RedisAddr      string
		PostgresPassword string
		RedisPassword    string
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
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, skipping...")
	}

	v := viper.New()

	// Set base config
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	// Automatically override with env vars
	v.AutomaticEnv()

	// Load base YAML config
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading base config: %v", err)
	}

	// Merge env-specific YAML
	if env != "" {
		v.SetConfigName(fmt.Sprintf("app.%s", env))
		if err := v.MergeInConfig(); err != nil {
			log.Fatalf("Error reading %s config: %v", env, err)
		}
	}

	// Bind specific env vars that don't directly map
	v.BindEnv("database.postgrespassword", "POSTGRES_PASSWORD")
	v.BindEnv("database.redispassword", "REDIS_PASSWORD")

	// Unmarshal into Config struct
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	AppConfig = &c
	fmt.Printf("Config loaded for env: %s\n", env)
}
