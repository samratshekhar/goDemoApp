package nosql

import (
	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"goDemoApp/internal/config"
	"goDemoApp/internal/datastore/dao"
	"goDemoApp/internal/logger"
	"goDemoApp/internal/models/widget"
	"time"
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

func (w DynamoWidgetDAO) Create(data []byte, newWidget *widget.Widget, userId string) error {
	dynamoWidget := &widget.DynamoWidget{
		Id:             newWidget.Id,
		WidgetTemplate: newWidget.TemplateType.String(),
		Data:           string(data),
	}

	av, err := dynamodbattribute.MarshalMap(dynamoWidget)
	log := logger.GetLogger()
	if err != nil {
		log.Error("Got error marshalling map:" + err.Error())
		return err
	}

	t := time.Now()
	// 2006 is placeholder for year, 01/1 is placeholder for month, and 02/2 is placeholder for date.
	dateString := t.Format("2006-1")
	dynamoWidgetAudit := &widget.DynamoWidgetAudit{
		WidgetId:     newWidget.Id,
		Version:      1,
		Data:         "",
		Timestamp:    int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond),
		Comments:     "Widget Created",
		UserId:       userId,
		EventType:    "create-" + dateString,
		TemplateName: newWidget.TemplateType.String(),
	}

	auditAv, err := dynamodbattribute.MarshalMap(dynamoWidgetAudit)
	if err != nil {
		log.Error("Got error marshalling map:" + err.Error())
		return err
	}

	input := &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					Item:      av,
					TableName: aws.String(w.TableName),
				},
			},
			{
				Put: &dynamodb.Put{
					Item:      auditAv,
					TableName: aws.String(w.AuditTableName),
				},
			},
		},
	}

	_, err = w.Client.TransactWriteItems(input)
	if err != nil {
		log.Error("Got error calling PutItem:" + err.Error())
	}
	return err
}

func (w DynamoWidgetDAO) GetById(string) (*widget.Widget, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetWidgetJsonById(widgetId string, widgetTemplateType string) (*widget.DynamoWidget, error) {
	result, err := w.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(w.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(widgetId),
			}, "widget-partition": {
				S: aws.String(widgetTemplateType),
			},
		},
	})
	log := logger.GetLogger()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var dynamoWidget = new(widget.DynamoWidget)

	err = dynamodbattribute.UnmarshalMap(result.Item, &dynamoWidget)
	if err != nil {
		return nil, err
	}
	return dynamoWidget, err
}

func (w DynamoWidgetDAO) SaveWidgetAndAudit(data []byte, newWidget *widget.Widget, auditData []byte, comments string, userId string) error {
	log := logger.GetLogger()
	dynamoWidget := &widget.DynamoWidget{
		Id:             newWidget.Id,
		WidgetTemplate: newWidget.TemplateType.String(),
		Data:           string(data),
	}
	av, err := dynamodbattribute.MarshalMap(dynamoWidget)
	if err != nil {
		log.Error("Got error marshalling map:" + err.Error())
		return err
	}

	version, err := w.fetchMaxAuditVersionByWidgetId(newWidget.Id)

	if err != nil {
		log.Error("Unable to fetch last audit version :" + err.Error())
		return err
	}

	t := time.Now()
	// 2006 is placeholder for year, 01/1 is placeholder for month, and 02/2 is placeholder for date.
	dateString := t.Format("2006-1")
	dynamoWidgetAudit := &widget.DynamoWidgetAudit{
		WidgetId:     newWidget.Id,
		Version:      version + 1,
		Data:         string(auditData),
		Timestamp:    int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond),
		Comments:     comments,
		UserId:       userId,
		EventType:    "edit-" + dateString,
		TemplateName: newWidget.TemplateType.String(),
	}
	auditAv, err := dynamodbattribute.MarshalMap(dynamoWidgetAudit)
	if err != nil {
		log.Error("Got error marshalling map:" + err.Error())
		return err
	}
	input := &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					Item:      av,
					TableName: aws.String(w.TableName),
				},
			},
			{
				Put: &dynamodb.Put{
					Item:      auditAv,
					TableName: aws.String(w.AuditTableName),
				},
			},
		},
	}

	_, err = w.Client.TransactWriteItems(input)
	if err != nil {
		log.Error("SaveWidgetAndAudit :: Got error calling TransactWriteItems:" + err.Error())
	}
	return err
}

func (w DynamoWidgetDAO) fetchMaxAuditVersionByWidgetId(widgetId string) (int32, error) {

	dynamoWidgetAudits, err := w.GetAuditsByLimit(widgetId, 1)
	if err != nil {
		return -1, err
	}

	if len(dynamoWidgetAudits) == 0 {
		return 0, nil
	}

	return dynamoWidgetAudits[0].Version, nil

}

func (w DynamoWidgetDAO) GetAuditsByLimit(widgetId string, limit int64) ([]widget.DynamoWidgetAudit, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(w.AuditTableName),
		KeyConditionExpression: aws.String("#key = :id"),
		ExpressionAttributeNames: map[string]*string{
			"#key": aws.String("widget-id"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(widgetId),
			},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int64(limit),
	}

	result, err := w.DDBClient.Query(input)

	if err != nil {
		return nil, err
	}

	return parseWidgetAuditsFromResult(result), nil
}

func parseWidgetAuditsFromResult(result *dynamodb.QueryOutput) []widget.DynamoWidgetAudit {
	if result == nil || len(result.Items) == 0 {
		return nil
	}
	log := logger.GetLogger()
	var audits []widget.DynamoWidgetAudit
	for _, item := range result.Items {
		var dynamoWidgetAudit = new(widget.DynamoWidgetAudit)
		err := dynamodbattribute.UnmarshalMap(item, &dynamoWidgetAudit)
		if err != nil {
			log.Error("Got error unmarshalling:")
			log.Error(err.Error())
			continue
		}
		audits = append(audits, *dynamoWidgetAudit)
	}
	return audits
}

func (w DynamoWidgetDAO) GetAuditsByVersion(widgetId string, version string) ([]widget.DynamoWidgetAudit, error) {
	panic("implement me")
	return nil, nil
}

func (w DynamoWidgetDAO) GetAuditsByDateRange(startDate string, endDate string, eventType string) ([]widget.DynamoWidgetAudit, error) {
	panic("implement me")
	return nil, nil
}
