package utils

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

// TODO: Implement logging with requestUuid retrieved from a context
// func LogWithContext(ctx context, label string, fields logger.Fields{}) {
// }

func LogWithContext(label string, fields logger.Fields) {
	// Have this for now
	logger.Debug(label, fields)
}
