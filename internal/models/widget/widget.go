package widget

import "goDemoApp/internal/models/dataProvider"

type Widget struct {
	Id                        string                      `json:"id" validate:"required"`
	Name                      string                      `json:"name" validate:"required"`
	TemplateType              TemplateType                `json:"templateType" validate:"required"`
	DataProvider              []dataProvider.DataProvider `json:"dataProvider" validate:"required"`
	OffsetDataProviderMapping string                      `json:"widgetOffsetDataProviderMapping"`
	MetaData                  map[string]string           `json:"metaData"`
}
