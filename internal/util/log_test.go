package util_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/sebastianmacias/clientify_livechat/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelFromString(t *testing.T) {
	res := util.LogLevelFromString("panic")
	assert.Equal(t, zerolog.PanicLevel, res)

	res = util.LogLevelFromString("warn")
	assert.Equal(t, zerolog.WarnLevel, res)

	res = util.LogLevelFromString("foo")
	assert.Equal(t, zerolog.DebugLevel, res)
}
