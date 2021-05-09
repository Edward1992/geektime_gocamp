package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serverConfig struct {
	Addr string
	ReadTimeout time.Duration
}

func NewServer(options ...ServerOption) http.Server {
	conf := &serverConfig{
		Addr: ":8080", // default to 127.0.0.1:8080
		ReadTimeout: time.Second,
	}

	for _, opt := range options {
		opt(conf)
	}

	return http.Server{
		Addr: conf.Addr,
		ReadTimeout: conf.ReadTimeout,
	}
}

type ServerOption func(conf *serverConfig)

func AddrOption(addr string) ServerOption {
	return func(conf *serverConfig) {
		conf.Addr = addr
	}
}

func ReadTimeoutOption(timeout time.Duration) ServerOption {
	return func(conf *serverConfig) {
		conf.ReadTimeout = timeout
	}
}

func serve(addr string) *http.Server {
	s := NewServer(AddrOption(addr))
	return &s
}

func StartServer(server *http.Server, stop <-chan bool) func() error {
	return func() error {
		go func() {
			<-stop
			server.Shutdown(context.Background())
		}()

		return server.ListenAndServe()
	}
}

func OnSignal(stop chan bool, sig ...os.Signal) func() error {
	return func() error {
		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}

		done := make(chan os.Signal, len(sig))
		signal.Notify(done, sig...)
		<-done
		signal.Stop(done)
		close(done)
		close(stop)
		return nil
	}
}

func main() {
	var group errgroup.Group
	stop := make(chan bool)
	group.Go(StartServer(serve(":8001"), stop))
	group.Go(StartServer(serve(":8002"), stop))
	group.Go(OnSignal(stop, syscall.SIGINT,syscall.SIGTERM, syscall.SIGABRT))

	err := group.Wait()
	fmt.Println("Errgroup finished with: ", err.Error())
}