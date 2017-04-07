package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/heltonmarx/goami/ami"
)

var (
	username = flag.String("username", "admin", "AMI username")
	secret   = flag.String("secret", "admin", "AMI secret")
	host     = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {
	flag.Parse()

	socket, err := ami.NewSocket(*host)
	if err != nil {
		log.Fatalf("socket error: %v\n", err)
	}
	if _, err := ami.Connect(socket); err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	//Login
	uuid, _ := ami.GetUUID()
	if err := ami.Login(socket, *username, *secret, "Off", uuid); err != nil {
		log.Fatalf("login error: %v\n", err)
	}
	defer ami.Logoff(socket, uuid)
	fmt.Printf("login ok!\n")

	data := ami.KhompSMSData{
		Device:       "b0",
		Destination:  "4899893791",
		Confirmation: true,
		Message:      "hey ho, let's go",
	}
	s, err := ami.KSendSMS(socket, uuid, data)
	if err != nil {
		fmt.Printf("sms sms error\n", err)
	}
	fmt.Printf("response: [%v]\n", s)
}
