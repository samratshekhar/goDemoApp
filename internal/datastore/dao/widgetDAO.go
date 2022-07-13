package dao

import "goDemoApp/internal/models/widget"

type WidgetDAO interface {
	GetAll() ([]*widget.Widget, error)
	GetAllWidgetIds() ([]string, error)
	Get([]string) ([]*widget.Widget, error)
	Create(data []byte, widget *widget.Widget, userId string) error
	GetById(string) (*widget.Widget, error)
	GetWidgetJsonById(string, string) (*widget.DynamoWidget, error)
	SaveWidgetAndAudit(data []byte, widget *widget.Widget, auditData []byte, comments string, userId string) error
	GetAuditsByLimit (widgetId string, limit int64) ([]widget.DynamoWidgetAudit, error)
	GetAuditsByVersion (widgetId string, version string) ([]widget.DynamoWidgetAudit, error)
	GetAuditsByDateRange (startDate string, endDate string, eventType string) ([]widget.DynamoWidgetAudit, error)
}

