package loggerfactory

import (
	"github.com/parag08/go-microservice-boilerplate/config"
	"github.com/parag08/go-microservice-boilerplate/pkg/loggerfactory/zap"
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
