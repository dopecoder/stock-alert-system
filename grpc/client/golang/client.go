package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
	"github.com/revulcan/stock-alert-system/grpc/trigger_service"
	"google.golang.org/grpc"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("server_addr", "localhost:80", "The server address in the format of host:port")
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

func CreateTrigger(client trigger_service.TriggerServiceClient, req *trigger_service.CreateTriggerReq) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := client.CreateTrigger(ctx, req)
	if err != nil {
		fmt.Println("CreateTrigger encountered error : ", err)
	} else {
		fmt.Println("CreateTrigger Response : ", res)
	}
}

func GetTrigger(client trigger_service.TriggerServiceClient, req *trigger_service.GetTriggerReq) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := client.GetTrigger(ctx, req)
	if err != nil {
		fmt.Println("GetTrigger encountered error : ", err)
	} else {
		fmt.Println("GetTrigger Response : ", res)
	}
}

func GetTriggerStatus(client trigger_service.TriggerServiceClient, req *trigger_service.TriggerStatusReq) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := client.GetTriggerStatus(ctx, req)
	if err != nil {
		fmt.Println("GetTriggerStatus encountered error : ", err)
	} else {
		fmt.Println("GetTriggerStatus Response : ", res)
	}
}

var wg sync.WaitGroup

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := trigger_service.NewTriggerServiceClient(conn)

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go create100TriggersFor1Instrument(client, fmt.Sprint(3044+i))
	}
	// CreateTrigger(client, &trigger_service.CreateTriggerReq{
	// 	Id: "1", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TNearPrice: 290, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE,
	// })
	// GetTriggerStatus(client, &trigger_service.TriggerStatusReq{
	// 	TriggerId: "1",
	// })
	wg.Wait()
}

func create100TriggersFor1Instrument(c trigger_service.TriggerServiceClient, exchngToken string) {
	for i := 1; i <= 100; i++ {
		id := exchngToken + fmt.Sprintf("-%d", i)
		CreateTrigger(c, &trigger_service.CreateTriggerReq{
			Id: id, TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TNearPrice: 10000.0 + float64(i), TPrice: 10000.0 + float64(i), Scrip: exchngToken, KiteToken: "", ExchangeToken: exchngToken, Exchange: trigger_service.Exchange_NSE,
		})
	}

	for i := 1; i <= 100; i++ {
		id := exchngToken + fmt.Sprintf("-%d", i)
		GetTriggerStatus(c, &trigger_service.TriggerStatusReq{
			TriggerId: id,
		})
	}
	wg.Done()
}
