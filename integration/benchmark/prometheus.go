package benchmark

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rs/zerolog"
)

type statsPusherImpl struct {
	pusher *push.Pusher
	cancel context.CancelFunc
	done   chan struct{}
}

// NewStatsPusher creates a new stats pusher that sends metrics to the given pushgateway.
func NewStatsPusher(
	ctx context.Context,
	log zerolog.Logger,
	pushgateway string,
	job string,
	gatherer prometheus.Gatherer,
) *statsPusherImpl {

	localCtx, cancel := context.WithCancel(ctx)
	sp := &statsPusherImpl{
		pusher: push.New(pushgateway, job).Gatherer(gatherer),
		cancel: cancel,
	}

	go func() {
		// Push stats to the gateway at regular intervals
		t := time.NewTicker(10 * time.Second)

		// Stop and flush stats on exit
		defer func() {
			t.Stop()

			_ = sp.pusher.Push()
			close(sp.done)
		}()

		for {
			select {
			case <-t.C:
				err := sp.pusher.Push()
				if err != nil {
					log.Warn().Err(err).Str("uri", pushgateway).Msg("failed to push metrics to pushgateway")
				}
			case <-localCtx.Done():
				return
			}
		}
	}()

	return sp
}

// Close stops the stats pusher and waits for it to finish.
func (sp *statsPusherImpl) Close() {
	sp.cancel()
	<-sp.done
}
