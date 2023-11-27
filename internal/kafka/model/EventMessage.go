package model

type EventMessage struct {
	Timestamp      int64  `json:"timestamp"`
	AccountMail    string `json:"accountMail"`
	ItemType       string `json:"itemType"`
	EventType      string `json:"eventType"`
	Sender         string `json:"sender"`
	Subject        string `json:"subject"`
	Fragment       string `json:"fragment"`
	SenderName     string `json:"senderName"`
	ItemId         int64  `json:"itemId"`
	FolderId       int64  `json:"folderId"`
	FolderName     string `json:"folderName"`
	ConversationId int64  `json:"conversationId"`
	IsHighPriority bool   `json:"isHighPriority"`
	IsLowPriority  bool   `json:"isLowPriority"`
	IsFromMe       bool   `json:"isFromMe"`
}
