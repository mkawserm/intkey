package cmd

import (
	"fmt"
	"github.com/mkawserm/intkey/pkg/conf"
	"github.com/mkawserm/intkey/pkg/intkey"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() {
	if conf.ServiceConfigurationIns().TimeFormat == "unix" {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	log.Debug().Interface("conf", conf.ServiceConfigurationIns()).Msg("configuration")

	log.Info().
		Str("host", conf.ServiceConfigurationIns().Host).
		Uint16("port", conf.ServiceConfigurationIns().Port).
		Msg("Initializing IntKey server...")

	lis, err := net.Listen("tcp",
		fmt.Sprintf(conf.ServiceConfigurationIns().Host+":%d",
			conf.ServiceConfigurationIns().Port))

	if err != nil {
		log.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	// register server
	intkey.RegisterIntKeyRPCServer(grpcServer, &intkey.RPCServer{})

	// handle signal
	idleChan := make(chan struct{})
	go func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		sig := <-signChan

		log.Info().Str("signal", sig.String()).Msg("shutdown signal received")
		log.Info().Msg("preparing for shutdown")

		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		grpcServer.GracefulStop()
		// Actual shutdown trigger.
		close(idleChan)
	}()

	go func() {
		log.Info().
			Msg("IntKey server started...")
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal().Msg(err.Error())
			os.Exit(1)
		}
	}()

	// Blocking until the shutdown is complete
	<-idleChan
	log.Info().Msg("shutdown complete")
}
