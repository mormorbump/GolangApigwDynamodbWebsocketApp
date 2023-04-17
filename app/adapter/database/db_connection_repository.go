// Interface Adapters層/Gateways(Repository)
package database

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/interfaces"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/domain/models"
)

type DBConnectionRepository struct {
	interfaces.IDBHandler
}

func (db DBConnectionRepository) Save(conn *models.Connection) error {
	err := db.DBPutItem(conn.LiveID, conn.UserID, conn.ConnectionID)
	if err != nil {
		return err
	}

	return nil
}

func (db DBConnectionRepository) Delete(connectionId string) error {
	err := db.DBDeleteItem(connectionId)
	if err != nil {
		return err
	}

	return nil
}

func (db DBConnectionRepository) FindByLiveIdAndUserId(liveId string, userId string) (models.Connection, error) {
	var out models.Connection
	err := db.DBGetItem(liveId, userId, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (db DBConnectionRepository) WhereByLiveId(liveId string) (models.Connections, error) {
	var out models.Connections
	err := db.DBQuery(liveId, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
