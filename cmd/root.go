package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-cli-cobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func initConfig() {
	configFile = ".golang-cli-cobra.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("COBRACLISAMPLES")
	HandleError(viper.BindEnv("API_KEY"))
	HandleError(viper.BindEnv("API_SECRET"))
	HandleError(viper.BindEnv("USERNAME"))
	HandleError(viper.BindEnv("PASSWORD"))

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
