package main

import (
	"context"
	"flag"
	"io"
	"log"

	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
	"google.golang.org/grpc"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr    = flag.String("server_addr", "localhost:8000", "The server address in the format of host:port")
	exchangeToken = flag.String("et", "3045", "Exchange Token of instrument to fetch")
	// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// printFeatures lists all the features within the given bounding Rectangle.
func WatchLTPforInstruments(client data_service.LTPServiceClient, instruments *data_service.Instruments, done chan bool) {
	log.Printf("Looking for LTPs for %v", instruments)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.WatchLTPforInstruments(ctx, instruments)
	if err != nil {
		log.Fatalf("%v.WatchLTPforInstruments(_) = _, %v", client, err)
	}
	for {
		select {
		case <-done:
			return
		default:
			ltp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.WatchLTPforInstruments(_) = _, %v", client, err)
			}
			log.Printf("WatchLTPforInstruments received: %v", ltp)
		}

	}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	// if *tls {
	// 	if *caFile == "" {
	// 		*caFile = data.Path("x509/ca_cert.pem")
	// 	}
	// 	creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
	// 	if err != nil {
	// 		log.Fatalf("Failed to create TLS credentials %v", err)
	// 	}
	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// } else {
	opts = append(opts, grpc.WithInsecure())
	// }

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := data_service.NewLTPServiceClient(conn)
	ch := make(chan bool)
	WatchLTPforInstruments(client, &data_service.Instruments{
		Items: []*data_service.Instrument{
			{ExchangeToken: "3045", Exchange: "NSE"},
			{ExchangeToken: "2885", Exchange: "NSE"},
		},
	}, ch)

	// time.Sleep(time.Second * 5)
	// ch <- true
	// close(ch)

	// time.Sleep(time.Second * 20)
	// ch <- true

}
