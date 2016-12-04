// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"member"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "   assignDay(string id)")
	fmt.Fprintln(os.Stderr, "   getResults()")
	fmt.Fprintln(os.Stderr, "  Member getMember(string id)")
	fmt.Fprintln(os.Stderr, "  ResultDay getResultByDay(string day)")
	fmt.Fprintln(os.Stderr, "   getNotAssign()")
	fmt.Fprintln(os.Stderr, "  void addMember(Member member)")
	fmt.Fprintln(os.Stderr, "  Member getMemberByLineID(string line_id)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := member.NewMemberServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "assignDay":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AssignDay requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.AssignDay(value0))
		fmt.Print("\n")
		break
	case "getResults":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetResults requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetResults())
		fmt.Print("\n")
		break
	case "getMember":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMember requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetMember(value0))
		fmt.Print("\n")
		break
	case "getResultByDay":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetResultByDay requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetResultByDay(value0))
		fmt.Print("\n")
		break
	case "getNotAssign":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetNotAssign requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetNotAssign())
		fmt.Print("\n")
		break
	case "addMember":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddMember requires 1 args")
			flag.Usage()
		}
		arg23 := flag.Arg(1)
		mbTrans24 := thrift.NewTMemoryBufferLen(len(arg23))
		defer mbTrans24.Close()
		_, err25 := mbTrans24.WriteString(arg23)
		if err25 != nil {
			Usage()
			return
		}
		factory26 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt27 := factory26.GetProtocol(mbTrans24)
		argvalue0 := member.NewMember()
		err28 := argvalue0.Read(jsProt27)
		if err28 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddMember(value0))
		fmt.Print("\n")
		break
	case "getMemberByLineID":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMemberByLineID requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetMemberByLineID(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
