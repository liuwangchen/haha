package impl

import (
	"context"

	"haha/conf"

	"github.com/Shopify/sarama"
	"github.com/reugn/go-streams"
	"github.com/reugn/go-streams/kafka"
)

type KafkaEntity struct {
	source *kafka.KafkaSource
	sink   *kafka.KafkaSink
}

func NewKafkaEntity() *KafkaEntity {
	return &KafkaEntity{}
}

func (ks *KafkaEntity) newSaramaConfig(config conf.EntityConfig) *sarama.Config {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Version, _ = sarama.ParseKafkaVersion(config.KafkaVersion)
	return saramaConfig
}

func (ks *KafkaEntity) ToSink(config conf.EntityConfig) streams.Sink {
	hosts := []string{config.Url}
	saramaConfig := ks.newSaramaConfig(config)
	ks.sink = kafka.NewKafkaSink(hosts, saramaConfig, config.KafkaTopic)
	return ks.sink
}

func (ks *KafkaEntity) ToSource(config conf.EntityConfig) streams.Source {
	hosts := []string{config.Url}
	ctx := context.Background()
	saramaConfig := ks.newSaramaConfig(config)
	groupID := config.KafkaGroup

	ks.source = kafka.NewKafkaSource(ctx, hosts, groupID, saramaConfig, config.KafkaTopic)
	return ks.source
}
