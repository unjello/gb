package cmd

import (
	"fmt"
	"os"

	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unjello/gb/core"
)

const (
	exitCodeTestsFailed = 1
)

var cfgFile string

var (
	verbose bool
	debug   bool
	quiet   bool
)

var rootCmd = &cobra.Command{
	Use:   "gb",
	Short: "gb: Great Builder, The",
	Long:  core.Desc,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gb.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "no output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gb")
	}

	viper.SetEnvPrefix("GB")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	log.Level = log.ERROR
	if viper.GetBool("verbose") {
		log.Level = log.INFO
	}
	if viper.GetBool("debug") {
		log.Level = log.DEBUG
	}
	if viper.GetBool("quiet") {
		log.Level = log.FATAL
	}

	if err == nil {
		log.Info("Using config file: " + tui.Dim(viper.ConfigFileUsed()))
	} else {
		log.Info(err.Error())
	}
}
