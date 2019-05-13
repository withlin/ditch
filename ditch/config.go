package ditch

import (
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)



//Config config for ditch
type Config struct {
	Port int `mapstructure:"port"`
	Level logrus.Level `mapstructure:"port"`
}

const DefaultPort int = 8080

func DefaultConfig() *Config {
	return &Config{
		Port: DefaultPort,
		Level:logrus.WarnLevel,
	}
}

func ConfigFlagSet() *flag.FlagSet {
	config := DefaultConfig()
	cmdFlags := flag.NewFlagSet("ditch flagset", flag.ContinueOnError)
	cmdFlags.Int("port",config.Port,"program started port")

	return  cmdFlags
}
