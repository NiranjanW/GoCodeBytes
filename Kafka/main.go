package main

import (
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/ksuid"
	"log"
	"time"
)

var (
	brokers  string
	topic    string
	records  int
	certFile string
	caFile   string
)

func init() {
	flag.StringVar(&brokers, "brokers", "localhost:9092", "broker addresses, comma-separated")
	flag.StringVar(&topic, "topic", "topic", "topic to produce to")
	flag.IntVar(&records, "records", 250000, "number of records to read from kafka")
}
func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func confluent() {
	groupID := ksuid.New().String()
	cm := &kafka.ConfigMap{
		"session.timeout.ms":              6000,
		"metadata.broker.list":            brokers,
		"enable.auto.commit":              false,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"group.id":                        groupID,
		"default.topic.config": kafka.ConfigMap{
			"auto.offset.reset": "earliest",
		},
	}
	consumer, err := kafka.NewConsumer(cm)
	check(err)
	defer consumer.Close()
	check(consumer.Subscribe(topic, nil))
	var start time.Time
	count := 0
loop:
	for {
		select {
		case m, ok := <-consumer.Events():
			if !ok {
				panic("unexpected eof")
			}
			switch event := m.(type) {
			case kafka.AssignedPartitions:
				consumer.Assign(event.Partitions)

			case kafka.PartitionEOF:
				// nop

			case kafka.RevokedPartitions:
				consumer.Unassign()

			case *kafka.Message:
				count++
				if count == 1 {
					start = time.Now()
				}
				if count == records {
					break loop
				}

			default:
				panic(m)

			}
		}
		elapsed := time.Now().Sub(start)
		fmt.Printf("confluent: %v records, %v\n", count, elapsed)
	}

}
