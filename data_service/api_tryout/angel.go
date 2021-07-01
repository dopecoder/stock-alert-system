package api_tryout

import (
	"fmt"
	"time"

	SmartApi "github.com/angelbroking-github/smartapigo"
	"github.com/angelbroking-github/smartapigo/websocket"
)

var socketClient *websocket.SocketClient

// Triggered when any error is raised
func onErrorAngel(err error) {
	fmt.Println("Error: ", err)
}

// Triggered when websocket connection is closed
func onCloseAngel(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnectAngel() {
	fmt.Println("Connected")
	err := socketClient.Subscribe()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when a message is received
func onMessageAngel(message []map[string]interface{}) {
	for _, m := range message {
		fmt.Printf("Received ticker for %s -> %v\n", m["tk"], m)
	}
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnectAngel(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnectAngel(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d\n", attempt)
}

func angel() {
	// var wg sync.WaitGroup

	// Create New Angel Broking Client
	ABClient := SmartApi.New("Y41983", "Yashas711811", "pVKn0AoU")

	// User Login and Generate User Session
	session, err := ABClient.GenerateSession()

	fmt.Println(session)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Get User Profile
	session.UserProfile, err = ABClient.GetUserProfile()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func(index int) {
	// 		res, err := ABClient.GetLTP(SmartApi.LTPParams{Exchange: "NSE", TradingSymbol: "SBIN-EQ", SymbolToken: "3045"})
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			wg.Done()
	// 			return
	// 		}
	// 		fmt.Printf("%d -> %v\n", index, res)
	// 		wg.Done()
	// 	}(i)
	// }

	// New Websocket Client
	socketClient = websocket.New(session.ClientCode, session.FeedToken, "nse_cm|3045")

	// Assign callbacks
	socketClient.OnError(onErrorAngel)
	socketClient.OnClose(onCloseAngel)
	socketClient.OnMessage(onMessageAngel)
	socketClient.OnConnect(onConnectAngel)
	socketClient.OnReconnect(onReconnectAngel)
	socketClient.OnNoReconnect(onNoReconnectAngel)

	// Start Consuming Data
	socketClient.Serve()

	// wg.Wait()
}
