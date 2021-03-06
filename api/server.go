package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"

	"code.olipicus.com/bsd_shift_for_clean/api/line"
	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/member/memberimp"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var connections map[*websocket.Conn]bool
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	state := flag.String("state", "develop", "startup message")
	flag.Parse()

	connections = make(map[*websocket.Conn]bool)
	router := mux.NewRouter()

	memberService := memberimp.MemberService{}
	memberService.State = *state

	processor := member.NewMemberServiceProcessor(memberService)
	protocolFactory := thrift.NewTJSONProtocolFactory()
	//server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	handler := NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)

	app, err := line.NewLineApp(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), &memberService)

	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/api", handler)
	router.HandleFunc("/ws", wsHandler)
	router.HandleFunc("/linebot", app.CallbackHandler)

	log.Println("Server Start ...")

	log.Fatal(http.ListenAndServe(":8802", router))

}

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully upgraded connection")
	connections[conn] = true

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			conn.Close()
			return
		}
		log.Println(string(msg))
		sendAll(msg)
	}

}

// NewThriftHandlerFunc is a function that create a ready to use Apache Thrift Handler function
func NewThriftHandlerFunc(processor thrift.TProcessor,
	inPfactory, outPfactory thrift.TProtocolFactory) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("%v", r.Body)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		w.Header().Add("Content-Type", "application/x-thrift")

		transport := thrift.NewStreamTransport(r.Body, w)
		ctx := context.Background()
		bool, err := processor.Process(ctx, inPfactory.GetProtocol(transport), outPfactory.GetProtocol(transport))

		if !bool {
			fmt.Printf("%v", err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	fmt.Fprintf(w, "Hello, %v", addrs)
}
