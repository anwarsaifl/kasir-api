package main

import (
	"fmt"
	"kasir-api/database"
	"log"
	"net/http"
	"os"
	"strings"

	viper "github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DBConn"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	defer db.Close()

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product", productHandler.HandleProductByID)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on " + addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Server failed to run", err)
	}
}
