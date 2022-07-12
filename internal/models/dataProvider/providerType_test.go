package dataProvider

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input = `{"type": "PROVIDER_TYPE_SHOPIFY"}`

func Test_UnmarshalJSON(t *testing.T) {
	type test struct {
		ProviderType ProviderType `json:"type"`
	}

	t.Run("invalid provider type", func(t *testing.T) {
		var test test
		invalidInput := `{"type": 1}`
		err := json.Unmarshal([]byte(invalidInput), &test)
		assert.NotNil(t, err)
		assert.Equal(t, test.ProviderType, PROVIDER_TYPE_INVALID)
	})

	t.Run("valid provider type", func(t *testing.T) {
		var test test
		err := json.Unmarshal([]byte(input), &test)
		assert.Nil(t, err)
		assert.Equal(t, test.ProviderType, PROVIDER_TYPE_SHOPIFY)
	})
}

func Test_GetTemplateType(t *testing.T) {
	t.Run("invalid provider type", func(t *testing.T) {
		assert.Equal(t, PROVIDER_TYPE_INVALID, GetProviderType("PROVIDER_TYPE_RANDOM"))
	})

	t.Run("valid provider type", func(t *testing.T) {
		// each new entry should be asserted here
		assert.Equal(t, PROVIDER_TYPE_INVALID, GetProviderType("TEMPLATE_TYPE_INVALID"))
		assert.Equal(t, PROVIDER_TYPE_NO_OP, GetProviderType("PROVIDER_TYPE_NO_OP"))
		assert.Equal(t, PROVIDER_TYPE_SHOPIFY, GetProviderType("PROVIDER_TYPE_SHOPIFY"))
	})
}
