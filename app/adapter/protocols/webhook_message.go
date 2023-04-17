package protocols

// カラムはそれぞれ変えてください。
type WebhookMessage struct {
	Type            string `json:"type"`
	Message         string `json:"message,omitempty"`
	LiveId          string `json:"liveId"`
	Name            string `json:"name"`
	ImageUrl        string `json:"imageUrl"`
	PlatformUserId  string `json:"platformUserId"`
	PlatformOwnerId string `json:"platformOwnerId"`
	SentAt          string `json:"sentAt,omitempty"`
	StartAt         string `json:"StartAt"`
}
