// Interface Adapters層/Gateways(Repository)
package usecase

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/domain/models"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/usecase/interfaces"
)

type WebSocketInteractor struct {
	DBConnectionRepository interfaces.IDBConnectionRepository
}

func (i *WebSocketInteractor) OnConnect(conn *models.Connection) error {
	return i.DBConnectionRepository.Save(conn)
}

func (i *WebSocketInteractor) OnDisconnect(connectionId string) error {
	return i.DBConnectionRepository.Delete(connectionId)
}

func (i *WebSocketInteractor) FetchConnection(liveId string, userId string) (models.Connection, error) {
	connection, err := i.DBConnectionRepository.FindByLiveIdAndUserId(liveId, userId)
	if err != nil {
		return connection, err
	}
	return connection, nil
}

func (i *WebSocketInteractor) FetchConnections(liveId string) (models.Connections, error) {
	connections, err := i.DBConnectionRepository.WhereByLiveId(liveId)
	if err != nil {
		return nil, err
	}
	return connections, nil
}
