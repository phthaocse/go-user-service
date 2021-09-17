package config
import (
	"github.com/phthaocse/user-service-go/utils"
)

type Config struct {
	ServerPort  string
	DbDriver string
	DbAddr string
	DbName string
	DbUser string
	DbPassword string
}


func SrvConfigSetup() (*Config, error) {
	port := utils.GetEnv("SERVER_PORT", "8080")
	port = ":" + port

	config := Config{
		ServerPort: port,
		DbDriver: "mysql",
		DbAddr: utils.GetEnv("DATABASE_ADDRESS", "localhost"),
		DbName: utils.GetEnv("DATABASE_NAME", "user_service"),
		DbUser: utils.GetEnv("DATABASE_USER", "root"),
		DbPassword: utils.GetEnv("DATABASE_PASSWORD", ""),
	}
	return &config, nil
}