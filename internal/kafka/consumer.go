package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/bss-office-facade/internal/logger"
)

type consumeFunction func(ctx context.Context, message *sarama.ConsumerMessage) error

type consumer struct {
	fn consumeFunction
}

func (consumer *consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := consumer.fn(session.Context(), message)

		if err != nil {
			logger.ErrorKV(session.Context(), "cannot handle msg", "err", err, "msg", message)
		}

		session.MarkMessage(message, "")
	}

	return nil
}

// StartConsuming запускает консюмера
func StartConsuming(ctx context.Context, brokers []string, topic string, group string, consumeFunction consumeFunction) error {

	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)

	if err != nil {
		return err
	}

	c := consumer{
		fn: consumeFunction,
	}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &c); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}
