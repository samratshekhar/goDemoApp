package dataProvider

type DataProvider struct {
	Type        ProviderType           `json:"type"`
	InputParams map[string]interface{} `json:"inputParams"`
	Mandatory   bool                   `json:"isMandatory"`
}
