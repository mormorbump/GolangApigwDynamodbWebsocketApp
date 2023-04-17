package models

type Connection struct {
	ConnectionID string `json:"connectionId"`
	RoomID       string `json:"roomId"`
	OwnerID      string `json:"-"`
	UserID       string `json:"userId"`
	TimeToExist  int64  `json:"-"`
}

type Connections []Connection
type ConnectionIds []string

func NewConnection(connectionId string, queryStringParameters map[string]string) *Connection {
	conn := new(Connection)
	conn.ConnectionID = connectionId
	conn.RoomID = queryStringParameters["roomId"]
	conn.OwnerID = queryStringParameters["OwnerId"]
	conn.UserID = queryStringParameters["UserId"]
	return conn
}
