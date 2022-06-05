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
	Options *NATSMessageSubscriberOptions
	Conn    *nats.Conn
	Logger  *zerolog.Logger
}

func (s *NATSMessageSubscriber) Subscribe(ctx context.Context) error {
	js, err := s.Conn.JetStream()
	if err != nil {
		return err
	}

	sub, err := js.PullSubscribe(s.Options.Subject, s.Options.Durable)
	if err != nil {
		return err
	}

	go func() {
		s.Logger.Info().Msgf("message subscriber started")

		for {
			select {
			case <-ctx.Done():
				s.Logger.Info().Msgf("message subscriber stopped")
				return
			default:
			}

			messages, err := sub.Fetch(s.Options.BatchSize)
			if err != nil {
				continue
			}

			for _, m := range messages {
				s.Logger.Info().Msgf("message received (%s)", m.Subject)

				handler, ok := s.Options.Handlers[m.Subject]
				if ok {
					if err := handler.Handle(ctx, m); err != nil {
						s.Logger.Error().Err(err).Send()
					}
				} else {
					s.Logger.Info().Msgf("no handler found for subject '%s'", m.Subject)
				}

				s.Logger.Info().Msgf("message handled (%s)", m.Subject)
			}
		}
	}()

	return nil
}
