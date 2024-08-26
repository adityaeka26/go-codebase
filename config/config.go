package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppEnv                string `mapstructure:"APP_ENV"`
	AppName               string `mapstructure:"APP_NAME"`
	AppVersion            string `mapstructure:"APP_VERSION"`
	RestPort              string `mapstructure:"REST_PORT"`
	GrpcPort              string `mapstructure:"GRPC_PORT"`
	GracefulPeriod        int    `mapstructure:"GRACEFUL_PERIOD"`
	JwtPrivateKey         string `mapstructure:"JWT_PRIVATE_KEY"`
	JwtPublicKey          string `mapstructure:"JWT_PUBLIC_KEY"`
	PostgresHost          string `mapstructure:"POSTGRES_HOST"`
	PostgresPort          string `mapstructure:"POSTGRES_PORT"`
	PostgresDb            string `mapstructure:"POSTGRES_DB"`
	PostgresUsername      string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword      string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSslEnabled    bool   `mapstructure:"POSTGRES_SSL_ENABLED"`
	KafkaSasl             bool   `mapstructure:"KAFKA_SASL"`
	KafkaHosts            string `mapstructure:"KAFKA_HOSTS"`
	KafkaUsername         string `mapstructure:"KAFKA_USERNAME"`
	KafkaPassword         string `mapstructure:"KAFKA_PASSWORD"`
	ElasticsearchUsername string `mapstructure:"ELASTICSEARCH_USERNAME"`
	ElasticsearchPassword string `mapstructure:"ELASTICSEARCH_PASSWORD"`
	ElasticsearchHost     string `mapstructure:"ELASTICSEARCH_HOST"`
	RedisHost             string `mapstructure:"REDIS_HOST"`
	RedisPassword         string `mapstructure:"REDIS_PASSWORD"`
}

func Load(filename string) (*EnvConfig, error) {
	var envCfg EnvConfig

	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&envCfg); err != nil {
		return nil, err
	}

	return &envCfg, nil
}
