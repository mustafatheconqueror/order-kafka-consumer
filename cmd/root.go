package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "OmsProjections",
	Short: "Oms Projection Application",
	Long:  "Command Line Interface for Oms Projection Applications",
}

func Execute() {
	InitEnvVariables()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	viper.AutomaticEnv()
}

func InitEnvVariables() {
	appName := os.Args[1]
	environmentName := strings.ToLower(os.Getenv("ENV"))

	fmt.Printf("{\"Message\": \"Environment: %s\"}\n", environmentName)
	if environmentName != "" {
		environmentFilePath := ""

		environmentFilePath = fmt.Sprintf("tools/%s/%s.env", appName, environmentName)

		fmt.Printf("{\"Message\": \"EnvironmentFilePath: %s\"}\n", environmentFilePath)

		absPath, _ := filepath.Abs(environmentFilePath)

		if _, err := os.Stat(absPath); err == nil {

			err = godotenv.Load(environmentFilePath)
			if err != nil {
				panic("Error loading .env file")
			}
			os.Setenv("APP_NAME", appName)
		}

	}
}
