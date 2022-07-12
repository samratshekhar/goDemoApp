package widget

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input = `{"type": "TEMPLATE_TYPE_DYNAMIC_SHELF"}`

func Test_UnmarshalJSON(t *testing.T) {
	type test struct {
		TemplateType TemplateType `json:"type"`
	}

	t.Run("invalid template type", func(t *testing.T) {
		var test test
		invalidInput := `{"type": 1}`
		err := json.Unmarshal([]byte(invalidInput), &test)
		assert.NotNil(t, err)
		assert.Equal(t, test.TemplateType, TEMPLATE_TYPE_INVALID)
	})

	t.Run("valid template type", func(t *testing.T) {
		var test test
		err := json.Unmarshal([]byte(input), &test)
		assert.Nil(t, err)
		assert.Equal(t, test.TemplateType, TEMPLATE_TYPE_DYNAMIC_SHELF)
	})
}

func Test_GetTemplateType(t *testing.T) {
	t.Run("invalid template type", func(t *testing.T) {
		assert.Equal(t, TEMPLATE_TYPE_INVALID, GetTemplateType("TEMPLATE_TYPE_RANDOM"))
	})

	t.Run("valid template type", func(t *testing.T) {
		// each new entry should be asserted here
		assert.Equal(t, TEMPLATE_TYPE_INVALID, GetTemplateType("TEMPLATE_TYPE_INVALID"))
		assert.Equal(t, TEMPLATE_TYPE_DYNAMIC_SHELF, GetTemplateType("TEMPLATE_TYPE_DYNAMIC_SHELF"))
	})
}


