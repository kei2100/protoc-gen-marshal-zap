package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestTypes_MarshalLogObject(t *testing.T) {
	m := &Types{
		Double1: 0.1,
	}
	enc := zapcore.NewMapObjectEncoder()
	err := m.MarshalLogObject(enc)
	if !assert.NoError(t, err) {
		return
	}
	assert.EqualValues(t, enc.Fields, map[string]interface{}{
		"secret1": "[MASKED]",
		"double1": 0.1,
	})
}
