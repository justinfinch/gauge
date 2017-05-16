package conf

import (
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

// ConfigureLogging will take the logging configuration and also adds
// a few default parameters
func ConfigureLogging(config *Config) (*logrus.Entry, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	level, err := logrus.ParseLevel(strings.ToUpper(config.LogLevel))
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(level)

	// always use the fulltimestamp
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		DisableTimestamp: false,
	})

	return logrus.StandardLogger().WithField("hostname", hostname), nil
}
