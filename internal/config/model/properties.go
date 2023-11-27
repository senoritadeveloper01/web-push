package model

type Properties struct {
	Mysql struct {
		ConnectionUrl string `json:"connectionUrl"`
	} `json:"mysql"`
	Kafka struct {
		Servers               string `json:"servers"`
		Topic                 string `json:"topic"`
		Id                    string `json:"id"`
		GroupId               string `json:"groupId"`
		SessionTimeoutMs      int    `json:"sessionTimeoutMs"`
		DefaultApiTimeoutMs   int    `json:"defaultApiTimeoutMs"`
		RequestTimeoutMs      int    `json:"requestTimeoutMs"`
		ReconnectBackoffMaxMs int    `json:"reconnectBackoffMaxMs"`
		ReconnectBackoffMs    int    `json:"reconnectBackoffMs"`
		RetryBackoffMs        int    `json:"retryBackoffMs"`
	} `json:"kafka"`
	Firebase struct {
		ConfigFilePath string `json:"configFilePath"`
		LogoUrl        string `json:"logoUrl"`
	} `json:"firebase"`
}
