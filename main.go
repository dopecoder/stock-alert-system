package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"path"
	"runtime"
	"time"

	"github.com/revulcan/stock-alert-system/grpc/server"
	"github.com/revulcan/stock-alert-system/grpc/trigger_service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func main() {

	// file, err := os.OpenFile("trigger_service.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.StampNano,
		DisableSorting:  true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := server.NewTriggerSystemServer()

	stop := make(chan interface{})
	setupStarterStopper(server, stop)

	trigger_service.RegisterTriggerServiceServer(grpcServer, server)
	grpcServer.Serve(lis)

}

func setupStarterStopper(server *server.TriggerServiceServer, quit chan interface{}) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				server.StartService(context.Background(), &trigger_service.StartServiceReq{})
				server.StopService(context.Background(), &trigger_service.StopServiceReq{})
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
