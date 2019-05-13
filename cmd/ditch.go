package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/withlin/ditch/ditch"
	"os"
)

var ditchCmd = &cobra.Command{
	Use:"ditch",
	Short: "mysql sync data system",
	Long: "ditch is a  mysql sync data system",
}
func Execute() {
	if err := ditchCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var config = &ditch.Config{}

func init(){
	cobra.OnInitialize(initConfig)
	ditchCmd.Flags().AddFlagSet(ditch.ConfigFlagSet())
	viper.BindPFlags(ditchCmd.Flags())
}

func initConfig(){
	if cfgFile == "" {
		viper.SetConfigFile(cfgFile)

	}else {
		// Find home directory.
		viper.AddConfigPath("./config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Info("No valid config found: Applying default values.")
	}

	if err := viper.Unmarshal(config); err != nil {
		logrus.WithError(err).Fatal("config: Error unmarshaling config")
	}

	logrus.SetLevel(config.Level)
}