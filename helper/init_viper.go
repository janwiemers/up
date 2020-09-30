package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// InitViperConfig initializes the config
func InitViperConfig() {
	envFolder := findFolder("env")
	viper.SetConfigName("default")

	viper.SetConfigType("yml")
	viper.AddConfigPath(envFolder)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %s ", err))
	}
	viper.SetConfigName("config-" + getCurrentProfile())
	_ = viper.MergeInConfig()
	viper.AutomaticEnv() // https://github.com/spf13/viper#working-with-environment-variables
}

func findFolder(folder string) string {
	path := folder
	for {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path
		}
		path = fmt.Sprintf("../%v", path)
	}
}

func getCurrentProfile() string {
	envVar := os.Getenv("PROFILE")
	if envVar != "" {
		return envVar
	}
	return "development"
}
