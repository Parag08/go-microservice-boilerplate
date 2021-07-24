// package logrus handles creating logrus logger
package logrus

import (
	"github.com/parag08/go-microservice-boilerplate/config"
	"github.com/parag08/go-microservice-boilerplate/pkg/logger"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func RegisterLog(lc config.LogConfig) error {
	//standard configuration
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)
	//log.SetOutput(os.Stdout)
	//customize it from configuration file
	err := customizeLogFromConfig(log, lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	//This is for loggerWrapper implementation
	//logger.Logger(&loggerWrapper{log})

	logger.SetLogger(log)
	return nil
}

// customizeLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(log *logrus.Logger, lc config.LogConfig) error {
	log.SetReportCaller(lc.EnableCaller)
	//log.SetOutput(os.Stdout)
	l := &log.Level
	err := l.UnmarshalText([]byte(lc.Level))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.SetLevel(*l)
	return nil
}
