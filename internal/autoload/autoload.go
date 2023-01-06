package autoload

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("APP_ENV")
	envFile := ".env"
	if env != "" {
		envFile = fmt.Sprintf(".env.%s", env)
	}
	_, err := os.Stat(envFile)
	if err == nil {
		err = godotenv.Load(envFile)
	}
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
