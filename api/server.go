package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"git.apache.org/thrift.git/lib/go/thrift"

	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/member/memberimp"
	"github.com/gorilla/mux"
)

func main() {
	state := flag.String("state", "develop", "startup message")
	flag.Parse()

	router := mux.NewRouter()

	memberService := memberimp.MemberService{}
	memberService.State = *state

	processor := member.NewMemberServiceProcessor(memberService)
	protocolFactory := thrift.NewTJSONProtocolFactory()
	//server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	handler := NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)
	router.HandleFunc("/api", handler)

	log.Println("Server Start ...")

	log.Fatal(http.ListenAndServe(":8080", router))

}

// NewThriftHandlerFunc is a function that create a ready to use Apache Thrift Handler function
func NewThriftHandlerFunc(processor thrift.TProcessor,
	inPfactory, outPfactory thrift.TProtocolFactory) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v", r.Body)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		w.Header().Add("Content-Type", "application/x-thrift")

		transport := thrift.NewStreamTransport(r.Body, w)
		bool, err := processor.Process(inPfactory.GetProtocol(transport), outPfactory.GetProtocol(transport))

		if !bool {
			fmt.Printf("%v", err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	fmt.Fprintf(w, "Hello, %v", addrs)
}
