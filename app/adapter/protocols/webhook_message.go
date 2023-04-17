package protocols

// カラムはそれぞれ変えてください。
type WebhookMessage struct {
	Type            string `json:"type"`
	Message         string `json:"message,omitempty"`
	RoomId          string `json:"roomId"`
	Name            string `json:"name"`
	ImageUrl        string `json:"imageUrl"`
	PlatformUserId  string `json:"platformUserId"`
	PlatformOwnerId string `json:"platformOwnerId"`
	SentAt          string `json:"sentAt,omitempty"`
	StartAt         string `json:"StartAt"`
}
