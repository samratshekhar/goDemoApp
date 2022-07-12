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

type DynamoWidget struct {
	Id             string `json:"id"`
	WidgetTemplate string `json:"widget-partition"`
	Data           string `json:"data"`
}

type DynamoWidgetAudit struct {
	WidgetId     string `json:"widget-id"`
	Version      int32  `json:"version"`
	Data         string `json:"data"`
	Timestamp    int64  `json:"updated_at"`
	Comments     string `json:"comments"`
	UserId       string `json:"userId"`
	EventType    string `json:"event_type_date"`
	TemplateName string `json:"templateName"`
}