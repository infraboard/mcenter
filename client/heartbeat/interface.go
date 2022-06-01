package heartbeat

import "context"

type Heartbeater interface {
	Heartbeat(ctx context.Context) error
}
