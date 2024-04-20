package kar_test

import (
	"context"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/testcontainers/testcontainers-go"
	TestKafka "github.com/testcontainers/testcontainers-go/modules/kafka"
)

const (
	TestClusterID = "test-cluster"
	KafkaImage    = "confluentinc/confluent-local:7.5.0"
	TestTopic     = "kar-test-topic"
)

func TestKafkaIsUp(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Start a Kafka container
	kafkaContainer, err := TestKafka.RunContainer(ctx,
		TestKafka.WithClusterID(TestClusterID),
		testcontainers.WithImage(KafkaImage),
	)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	// Ensure the container is terminated after the test
	defer func() {
		if err := kafkaContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Get the Kafka brokers from the container
	kBrokers, err := kafkaContainer.Brokers(ctx)
	if err != nil {
		t.Fatalf("failed to get host: %s", err)
	}

	// Connect to the Kafka broker
	conn, err := kafka.Dial("tcp", kBrokers[0])
	if err != nil {
		t.Fatalf("failed to dial Kafka broker: %s", err)
	}

	// Ensure the connection is closed after the test
	defer func() {
		if err := conn.Close(); err != nil {
			t.Errorf("failed to close connection: %s", err)
		}
	}()

	t.Log("Kafka is up and running")

	// Create a Kafka topic
	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             TestTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		t.Fatalf("failed to create topic: %s", err)
	}

	// Create a Kafka writer
	w := &kafka.Writer{
		Addr:     kafka.TCP(kBrokers...),
		Topic:    TestTopic,
		Balancer: &kafka.LeastBytes{},
	}

	// Write a message to Kafka
	err = w.WriteMessages(ctx, kafka.Message{
		Key:   nil,
		Value: []byte("Hello, Kafka!"),
	})
	if err != nil {
		t.Fatalf("failed to write message: %s", err)
	}

	t.Log("Message sent to Kafka")
}
