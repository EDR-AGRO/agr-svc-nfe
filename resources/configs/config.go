package configs

import (
	"br/com/agr/nfe/infrastructure/apm"
	"br/com/agr/nfe/infrastructure/logger"
	"net/http"

	"github.com/spf13/viper"
)

var Cfg *conf

type conf struct {
	Port            string `mapstructure:"API_PORT"`
	ShowSwagger     int    `mapstructure:"SHOW_SWAGGER"`
	LogLevel        string `mapstructure:"LOG_LEVEL"`
	RequestTimeOut  int    `mapstructure:"REQUEST_TIMEOUT"`
	DB_PASSWORD     string `mapstructure:"DB_PASSWORD"`
	DB_USER         string `mapstructure:"DB_USER"`
	DB_PORT         int    `mapstructure:"DB_PORT"`
	DB_SERVER       string `mapstructure:"DB_SERVER"`
	DB_DATABASE     string `mapstructure:"DB_DATABASE"`
	OtlpEndpoint    string `mapstructure:"OPENTELEMETRY_OTLP_ENDPOINT"`
	OtlpServiceName string `mapstructure:"OTLP_SERVICE"`
	ClientRequest   *http.Client
}

func LoadConfig(apmt *apm.ApmTransaction) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("/app/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf(apmt.Ctx, "Error reading config file: %v", err.Error())
		logger.Infof(apmt.Ctx, "Trying to read local.env")
		viper.AddConfigPath(".")
		viper.SetConfigFile("resources/local.env")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Errorf(apmt.Ctx, "Error reading local config file: %v", err.Error())
			return nil, err
		}
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		logger.Errorf(apmt.Ctx, "Error unmarshalling config file: %v", err.Error())
		return nil, err
	}

	return Cfg, nil
}
