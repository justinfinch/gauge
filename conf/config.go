package conf

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the application's configuration
type Config struct {
	Port     int64
	LogLevel string
}

//LoadConfig loads a config struct from cmd line or env vars
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	// NOTE: the order of these blocks doesn't matter, the hierarchy is handled by the viper library

	// from the command itself
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	// from the environment
	viper.SetEnvPrefix("GAUGE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
