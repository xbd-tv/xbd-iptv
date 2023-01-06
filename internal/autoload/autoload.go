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
	if err != nil {
		return
	}
	if os.IsNotExist(err) {
		return
	}
	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
}
