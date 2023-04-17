// Interface Adapters層/Controllers
package controllers

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/database"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/interfaces"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/adapter/protocols"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/domain/models"
	"com.mormorbump/apigateway.dynamodb.websockets.golang/usecase"
	"log"
)

type WebSocketController struct {
	webSocketInteractor *usecase.WebSocketInteractor
	clientHandler       interfaces.IClientHandler
}

func NewWebSocketController(dbHandler interfaces.IDBHandler, clientHandler interfaces.IClientHandler) *WebSocketController {
	return &WebSocketController{
		webSocketInteractor: &usecase.WebSocketInteractor{
			DBConnectionRepository: &database.DBConnectionRepository{IDBHandler: dbHandler},
		},
		clientHandler: clientHandler,
	}
}

func (wsc WebSocketController) Create(connectionId string, queryStringParameters map[string]string) (*models.Connection, error) {
	conn := models.NewConnection(connectionId, queryStringParameters)
	if err := wsc.webSocketInteractor.OnConnect(conn); err != nil {
		return conn, err
	}
	return conn, nil
}

func (wsc WebSocketController) Destroy(connectionId string) error {
	if err := wsc.webSocketInteractor.OnDisconnect(connectionId); err != nil {
		return err
	}
	return nil
}

func (wsc WebSocketController) CheckConnect(liveId string, platformOwnerId string) (bool, error) {
	connection, err := wsc.webSocketInteractor.FetchConnection(liveId, platformOwnerId)
	if err != nil {
		return false, err
	}
	isConnect, err := wsc.clientHandler.GetConnection(connection.ConnectionID)
	if err != nil {
		return isConnect, err
	}
	return isConnect, nil
}

func (wsc WebSocketController) SendMessage(webhookMessage *protocols.WebhookMessage) error {
	connections, err := wsc.webSocketInteractor.FetchConnections(webhookMessage.LiveId)
	if err != nil {
		return err
	}

	wsc.postToConnections(connections, webhookMessage)
	return nil
}

func (wsc WebSocketController) postToConnections(connections models.Connections, body interface{}) {
	for _, connection := range connections {
		err := wsc.clientHandler.PostToConnection(connection.ConnectionID, body)
		if err != nil {
			log.Println(err)
		}
	}
}
