package dataProvider

import "encoding/json"

type ProviderType int32

const (
	PROVIDER_TYPE_INVALID ProviderType = 0
	PROVIDER_TYPE_NO_OP   ProviderType = 1
	PROVIDER_TYPE_SHOPIFY ProviderType = 2
)

var providerTypeName = map[ProviderType]string{
	0: "PROVIDER_TYPE_INVALID",
	1: "PROVIDER_TYPE_NO_OP",
	2: "PROVIDER_TYPE_SHOPIFY",
}

var providerTypeValue = map[string]ProviderType{
	"PROVIDER_TYPE_INVALID": 0,
	"PROVIDER_TYPE_NO_OP":   1,
	"PROVIDER_TYPE_SHOPIFY": 2,
}

func (p ProviderType) String() string {
	if name := providerTypeName[p]; name != "" {
		return name
	}
	return providerTypeName[PROVIDER_TYPE_INVALID]
}

func GetProviderType(providerTypeString string) ProviderType {
	return providerTypeValue[providerTypeString]
}

func (p *ProviderType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*p = GetProviderType(j)
	return nil
}

func (p ProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
