package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	properties "web-push/internal/config/model"
	"web-push/internal/utils"
)

func InitConfiguration(ctx context.Context) properties.Properties {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Users/tcnseri/dev/go-workspace/web-push/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			utils.LogFatalAndStop(ctx, utils.GetLogDetails("[VIPER_OPERATIONS]", "[INIT_CONFIGURATION]", fmt.Errorf("config file not found: %w", err).Error()))
		}
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[VIPER_OPERATIONS]", "[INIT_CONFIGURATION]", fmt.Errorf("config file was found but another error ocurred: %w", err).Error()))
	}

	var config properties.Properties

	err = viper.Unmarshal(&config)
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[VIPER_OPERATIONS]", "[INIT_CONFIGURATION]", fmt.Errorf("error mapping properties: %w", err).Error()))
	}

	return config
}
