package chat

type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	Username       string   `json:"username"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}
