package cmd

import (
	"fmt"

	"github.com/covexo/devspace/pkg/devspace/upgrade"
	"github.com/covexo/devspace/pkg/util/logutil"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logutil.GetLogger("default", true)
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devspace",
	Short: "Welcome to the DevSpace CLI!",
	Long: `With a DevSpace you can program, build and execute cloud-native applications directly inside a Kubernetes cluster. Start your DevSpace now with:

	devspace up`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if upgrade.GetVersion() != "" {
		rootCmd.Version = upgrade.GetVersion()
	}

	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Panic(err)
		}

		// Search config in home directory with name ".devspace" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".devspace")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
