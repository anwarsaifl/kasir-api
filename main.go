package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	viper "github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
	}

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on " + addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Server failed to run", err)
	}
}
