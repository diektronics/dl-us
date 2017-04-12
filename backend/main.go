package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/diektronics/dl-us/backend/dl"
	"github.com/diektronics/dl-us/cfg"
	dlpb "github.com/diektronics/dl-us/protos/dl"
	"google.golang.org/grpc"
)

var cfgFile = flag.String(
	"cfg",
	filepath.Join(os.Getenv("HOME"), ".config", "dl", "config.cfg"),
	"Configuration file in text protobuf format indicating DB credentials and mailing details.",
)

func main() {
	flag.Parse()
	c, err := cfg.GetConfig(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	d := dl.New(c, 5)
	s := grpc.NewServer()
	dlpb.RegisterDlServer(s, d)
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", c.Backend.Port))
	if err != nil {
		log.Fatal("listening:", err)
	}
	go s.Serve(l)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("entering lame duck mode for 2 seconds")
	time.Sleep(time.Duration(2) * time.Second)
}
