package aimanager

import "context"

type Client interface {
	Send(ctx context.Context, messages []*Message) (string, error)
}
