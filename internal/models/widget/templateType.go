package widget

import "encoding/json"

type TemplateType int32

const (
	TEMPLATE_TYPE_INVALID       TemplateType = 0
	TEMPLATE_TYPE_DYNAMIC_SHELF TemplateType = 1
)

var templateTypeName = map[TemplateType]string{
	0: "TEMPLATE_TYPE_INVALID",
	1: "TEMPLATE_TYPE_DYNAMIC_SHELF",
}

var templateTypeValue = map[string]TemplateType{
	"TEMPLATE_TYPE_INVALID":       0,
	"TEMPLATE_TYPE_DYNAMIC_SHELF": 1,
}

func (t TemplateType) String() string {
	if name := templateTypeName[t]; name != "" {
		return name
	}
	return templateTypeName[TEMPLATE_TYPE_INVALID]
}

func GetTemplateType(templateTypeString string) TemplateType {
	return templateTypeValue[templateTypeString]
}

func (t *TemplateType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = GetTemplateType(j)
	return nil
}

func (t TemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
