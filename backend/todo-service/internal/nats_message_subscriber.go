package internal

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type NATSMessageSubscriberOptions struct {
	Subject   string
	Durable   string
	BatchSize int
	Handlers  map[string]NATSMessageHandler
}

type NATSMessageSubscriber struct {
	options *NATSMessageSubscriberOptions
	nc      *nats.Conn
	logger  *zerolog.Logger
}

func (s *NATSMessageSubscriber) Subscribe(ctx context.Context) error {
	js, err := s.nc.JetStream()

	if err != nil {
		return err
	}

	sub, err := js.PullSubscribe(s.options.Subject, s.options.Durable)

	if err != nil {
		return err
	}

	go func() {
		s.logger.Info().Msgf("message subscriber started")

		for {
			select {
			case <-ctx.Done():
				s.logger.Info().Msgf("message subscriber stopped")
				return
			default:
			}

			messages, err := sub.Fetch(s.options.BatchSize)
			if err != nil {
				continue
			}

			for _, m := range messages {
				s.logger.Info().Msgf("message received (%s)", m.Subject)

				handler, ok := s.options.Handlers[m.Subject]
				if ok {
					if err := handler.Handle(ctx, m); err != nil {
						s.logger.Error().Err(err).Send()
					}
				} else {
					s.logger.Info().Msgf("no handler found for subject '%s'", m.Subject)
				}

				s.logger.Info().Msgf("message handled (%s)", m.Subject)
			}
		}
	}()

	return nil
}
