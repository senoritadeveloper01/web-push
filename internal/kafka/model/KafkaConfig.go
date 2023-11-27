package model

type KafkaConfig struct {
	Brokers               string `json:"servers"`
	Topic                 string `json:"topic"`
	Id                    string `json:"id"`
	GroupId               string `json:"groupId"`
}
