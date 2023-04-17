// Interface Adapters層/Gateways(Repository)
package interfaces

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/domain/models"
)

type IDBConnectionRepository interface {
	Save(conn *models.Connection) error
	Delete(connectionId string) error
	WhereByRoomId(roomId string) (models.Connections, error)
	FindByRoomIdAndUserId(roomId string, userId string) (models.Connection, error)
}
