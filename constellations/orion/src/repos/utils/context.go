package utils

import (
	"context"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func LogWithContext(ctx context.Context, label string, fields logger.Fields) {
	requestUuid := ctx.Value(domains.REQUEST_UUID)
	extraFields := logger.CombineFields(fields, logger.Fields{
		"requestUuid": requestUuid,
	})
	logger.Debug(label, extraFields)
}
