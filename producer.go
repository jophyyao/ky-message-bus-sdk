package kymessage

import (
	"encoding/json"
	"fmt"

	"github.com/jophyyao/ky-message-bus-sdk/config"
	"github.com/jophyyao/ky-message-bus-sdk/utils"
	"github.com/siddontang/go/log"
)

type ProducerConfig struct {
	Platform    string `json:"platform"`
	Password    string `json:"password"`
	Producer    string `json:"producer"`
	PublishKey  string `json:"publish_key"`
	MessageJson string `json:"message_json"`
}

type ProducerApiResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewProducer(producer_name, publish_key, message_json string) (int64, string) {
	cf := config.InitConfig()
	url := fmt.Sprintf("%[1]s:%[2]d/%[3]s/message/producer", cf.Host, cf.Port, cf.UrlPrefix)
	httpArgs := ProducerConfig{
		Platform:    Platform,
		Password:    Password,
		Producer:    producer_name,
		PublishKey:  publish_key,
		MessageJson: message_json,
	}
	jsonStr, _ := json.Marshal(httpArgs)
	http_response, err := utils.HttpPostJson(
		url,
		jsonStr,
	)
	if err != nil {
		log.Fatal("producer http requset error: ", err.Error())
		return 101, err.Error()
	}
	var producer_api_response ProducerApiResponse
	json.Unmarshal([]byte(http_response), &producer_api_response)
	return producer_api_response.Code, producer_api_response.Message
}
