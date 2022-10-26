package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/shettyh/threadpool"
	"go.uber.org/zap"
)

var Pool *threadpool.ThreadPool

type KibanaMessage struct {
	Url          string      `json:"url"`
	ServiceName  string      `json:"mservice_name"`
	UserAgent    string      `json:"user_agent"`
	FuncName     string      `json:"fuc_name"`
	Token        string      `json:"token"`
	Input        interface{} `json:"input"`
	Output       interface{} `json:"output"`
	ExecutedTime float64     `json:"dt"`
	Version      string      `json:"version"`
}
type KibanaMessageAll struct {
	Phone          string `json:"phone"`
	CustomerId     string `json:"customerId"`
	IpAddress      string `json:"ipAddress"`
	IsCanhTo       string `json:"isCanhTo"`
	IsCustomer     string `json:"isCustomer"`
	Provider       string `json:"orovider"`
	DeviceId       string `json:"deviceId"`
	DevicePlatform string `json:"devicePlatform"`
	Lang           string `json:"lang"`
	AppVersion     string `json:"appVersion"`
	ContractNo     string `json:"contractNo"`
	LocationZone   string `json:"locationZone"`
	LocationCode   string `json:"locationCode"`
	BranchName     string `json:"branchName"`
	Status         int    `json:"status"`
	ServiceName    string `json:"serviceName"`
	FunctionName   string `json:"functionName"`
	ActionName     string `json:"actionName"`
	Url            string `json:"url"`
	DateAction     string `json:"dateAction"`
	PositionIcon   string `json:"positionIcon"`
	Referer        string `json:"referer"`
	Note           string `json:"note"`
	TypeLog        string `json:"typeLog"`
	ProcessTime    string `json:"processTime"`
	Topic_name     string `json:"topic_name"`
	ScreenId       string `json:"screenId"`
}
type SendLogToKibanaTask struct {
	Log             *zap.Logger
	Message         KibanaMessage
	BootstrapServer []string
	TopicName       string
}
type SendLogToKibanaAllTask struct {
	Log             *zap.Logger
	Message         KibanaMessage
	BootstrapServer []string
	TopicName       string
}

func (m KibanaMessage) ToByte() []byte {
	buffers := new(bytes.Buffer)
	json.NewEncoder(buffers).Encode(m)
	return buffers.Bytes()
}
func (t *SendLogToKibanaTask) Run() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(t.BootstrapServer...),
		Topic:    t.TopicName,
		Balancer: &kafka.LeastBytes{},
	}
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	err := writer.WriteMessages(ctxTimeout,
		kafka.Message{
			Key:   []byte(t.Message.ServiceName),
			Value: t.Message.ToByte(),
		},
	)
	if err != nil {
		fmt.Println("Error Send log to kibana", err)
	}
}
func (m KibanaMessageAll) ToByte() []byte {
	buffers := new(bytes.Buffer)
	json.NewEncoder(buffers).Encode(m)
	return buffers.Bytes()
}
func (t *SendLogToKibanaAllTask) Run() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(t.BootstrapServer...),
		Topic:    t.TopicName,
		Balancer: &kafka.LeastBytes{},
	}
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	err := writer.WriteMessages(ctxTimeout,
		kafka.Message{
			Key:   []byte(t.Message.ServiceName),
			Value: t.Message.ToByte(),
		},
	)
	if err != nil {
		t.Log.Error("Error Send log to kibana all", zap.Error(err))
	}

}
