package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"git.apache.org/thrift.git/lib/go/thrift"

	"code.olipicus.com/bsd_shift_for_clean/api/member/gen-go/member"
	"code.olipicus.com/bsd_shift_for_clean/api/member/memberimp"
	"code.olipicus.com/go_rest_api/api/person"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//REST API For Person
	router.HandleFunc("/person/{id}", person.Handler.GetDataByID).Methods("GET")
	router.HandleFunc("/person", person.Handler.InsertData).Methods("POST")
	router.HandleFunc("/person/{id}", person.Handler.UpdateByID).Methods("PUT")
	router.HandleFunc("/person/{id}", person.Handler.RemoveByID).Methods("DELETE")

	//router.HandleFunc("/member/random/{id}", member.Handler.Random).Methods("GET")
	//router.HandleFunc("/member/{id}", member.Handler.GetDataByID).Methods("GET")
	//router.HandleFunc("/", member.Handler.Result).Methods("GET")

	memberService := memberimp.MemberService{}

	processor := member.NewMemberServiceProcessor(memberService)
	protocolFactory := thrift.NewTJSONProtocolFactory()
	//server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	handler := NewThriftHandlerFunc(processor, protocolFactory, protocolFactory)
	router.HandleFunc("/", handler)

	log.Println("Server Start ...")

	log.Fatal(http.ListenAndServe(":8081", router))

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
