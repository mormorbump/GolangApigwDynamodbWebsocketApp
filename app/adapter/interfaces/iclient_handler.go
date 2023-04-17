package interfaces

type IClientHandler interface {
	PostToConnection(connectionId string, response interface{}) error
	GetConnection(connectionId string) (bool, error)
}
