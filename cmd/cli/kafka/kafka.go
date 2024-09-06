package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"strings"
	"time"
)

var (
	kafkaProducer *kafka.Writer
)
var (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // cân bằng tải
	}
}

// for consumer
func getKafkaReader(KafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(KafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              //10KB
		MaxBytes:       10e6,              //10MB
		CommitInterval: time.Second,       //
		StartOffset:    kafka.FirstOffset, //
	})
}

type stockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan

func newStock(msg, typeMsg string) *stockInfo {
	s := stockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s
	jsonbody, _ := json.Marshal(body)
	// tao msg
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonbody),
	}
	// viet message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer hong mua ATC
func RegisterConsumerATC(id int) {
	// group consumer??
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) //consumer-group-
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Println("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer (%d) error: %v", id, err)
		}
		fmt.Println("Consumer (%d), hong topic:%v, partition:%v, offset:%v,time%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()
	r.POST("action/stock", actionStock)
	// dang ky 2 user de mau stock trong atc (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	r.Run(":8999")
}
