package loggerImpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	assert.NotNil(t, logger)
	assert.EqualValues(t, "debug", logger.GetLevel().String())
}

