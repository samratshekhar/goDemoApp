package nosql

import (
	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"goDemoApp/internal/config"
	"goDemoApp/internal/datastore/dao"
	"goDemoApp/internal/models/widget"
)

type DynamoWidgetDAO struct {
	Client         dynamodbiface.DynamoDBAPI
	DDBClient      dynamodbiface.DynamoDBAPI
	TableName      string
	AuditTableName string
	PageLimit      int64
}

func NewDaxClient(region, endpoint string) (dynamodbiface.DynamoDBAPI, error) {
	cfg := dax.DefaultConfig()
	cfg.HostPorts = []string{endpoint}
	cfg.Region = region
	return dax.New(cfg)
}

func NewDdbClient(region, endpoint string) (dynamodbiface.DynamoDBAPI, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region),
	},
	)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

func InitDDB(configuration config.DynamoConfiguration) dao.WidgetDAO {
	var dynamoDao DynamoWidgetDAO
	var ddb, svc dynamodbiface.DynamoDBAPI
	ddb, err1 := NewDdbClient(configuration.DBRegion, configuration.DBEndpoint)
	if err1 != nil {
		panic("failed to created DDB client")
	}
	if configuration.DaxEnabled {
		dax, err := NewDaxClient(configuration.DaxRegion, configuration.DaxEndpoint)
		if err != nil {
			panic("failed to create Dax Client")
		}
		svc = dax
	} else {
		svc = ddb
	}
	dynamoDao = DynamoWidgetDAO{
		Client:         svc,
		DDBClient:      ddb,
		TableName:      configuration.WidgetTable,
		AuditTableName: configuration.WidgetAuditTable,
		PageLimit:      configuration.PageLimit,
	}
	return dynamoDao
}

func (w DynamoWidgetDAO) GetAll() ([]*widget.Widget, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetAllWidgetIds() ([]string, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) Get([]string) ([]*widget.Widget, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) Create(data []byte, widget *widget.Widget, userId string) error {
	panic("implement me")
	return nil
}

func (w DynamoWidgetDAO) GetById(string) (*widget.Widget, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetWidgetJsonById(string) (*widget.DynamoWidget, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) SaveWidgetAndAudit(data []byte, widget *widget.Widget, auditData []byte, comments string, userId string) error {
	panic("implement me")
	return nil
}
func (w DynamoWidgetDAO) GetAuditsByLimit(widgetId string, limit int64) ([]widget.DynamoWidgetAudit, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetAuditsByVersion(widgetId string, version string) ([]widget.DynamoWidgetAudit, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetAuditsByDateRange(startDate string, endDate string, eventType string) ([]widget.DynamoWidgetAudit, error) {
	panic("implement me")
	return nil, nil
}
