package service

import (
	"my-microservice/cmd/internal/model"

	"github.com/IBM/sarama"
)

type KafkaService struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func NewKafkaService() *KafkaService {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Consumer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		panic(err)
	}

	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		panic(err)
	}

	return &KafkaService{
		producer: producer,
		consumer: consumer,
	}
}

func (s *KafkaService) SendMessage(msg *model.Message) error {
	_, _, err := s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "messages",
		Value: sarama.StringEncoder(msg.Text),
	})
	return err
}
