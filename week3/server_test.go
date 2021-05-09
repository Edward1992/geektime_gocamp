package week3

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestServer(t *testing.T) {
	s := http.Server{
		Addr: ":8080",
	}

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		return s.Shutdown(ctx)
	})

	g.Go(func() error {
		return s.ListenAndServe()
	})

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	g.Go(func() error {
		for {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case <- signal_chan:
				return s.Shutdown(ctx)
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		t.Error(err)
	}
}