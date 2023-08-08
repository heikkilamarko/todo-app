package internal

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NATSMessageSubscriberOptions struct {
	Stream    string
	Consumer  string
	BatchSize int
	Handlers  map[string]NATSMessageHandler
}

type NATSMessageSubscriber struct {
	Options *NATSMessageSubscriberOptions
	Conn    *nats.Conn
	Logger  *slog.Logger
}

func (s *NATSMessageSubscriber) Subscribe(ctx context.Context) error {
	js, err := jetstream.New(s.Conn)
	if err != nil {
		return err
	}

	con, err := js.Consumer(ctx, s.Options.Stream, s.Options.Consumer)
	if err != nil {
		return err
	}

	go func() {
		s.Logger.Info("message subscriber started")

		for {
			select {
			case <-ctx.Done():
				s.Logger.Info("message subscriber stopped")
				return
			default:
			}

			batch, err := con.Fetch(s.Options.BatchSize)
			if err != nil {
				continue
			}

			for msg := range batch.Messages() {
				s.Logger.Info("message received", "subject", msg.Subject())

				handler, ok := s.Options.Handlers[msg.Subject()]
				if ok {
					if err := handler.Handle(ctx, msg); err != nil {
						s.Logger.Error(err.Error())
					}
				} else {
					s.Logger.Info("handler not found", "subject", msg.Subject())
				}

				s.Logger.Info("message handled", "subject", msg.Subject())
			}
		}
	}()

	return nil
}
