package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hugdubois/svc-fizzbuzz/service"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	svc     = service.NewService()
	svcName = fmt.Sprintf("%s-%s", svc.Name, svc.Version)

	rootCmd = &cobra.Command{
		Use:   "svc-fizzbuzz",
		Short: "A simple fizzbuzz microservice",
		Long: fmt.Sprintf(`To get started run the serve subcommand which will start a server:

  $ svc-fizzbuzz serve

Curl examples:
  $ curl -X GET    http://localhost%[1]s/
  $ curl -X GET    http://localhost%[1]s/version
  $ curl -X GET    http://localhost%[1]s/metrics
  $ curl -X GET    http://localhost%[1]s/status`, DefautAddress),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.svc-fizzbuzz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".svc-fizzbuzz" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".svc-fizzbuzz")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
