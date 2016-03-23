package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"github.com/majest/go-microservice/consul"
	"github.com/majest/user-service/dao"
	"github.com/majest/user-service/pb"
	"github.com/majest/user-service/server"
	"github.com/majest/user-service/service"
)

var serviceIP, consulHost, consulIP, dbHost, dbRegion string
var servicePort, consulPort, dbPort int

func init() {
	flag.StringVar(&serviceIP, "ip", "127.0.0.1", "Service ip. Should be local ip if run locally")
	flag.IntVar(&servicePort, "port", 9090, "Service port")
	flag.StringVar(&consulHost, "consulhost", "192.168.99.101", "Consul node ip")
	flag.StringVar(&dbHost, "dbhost", "192.168.99.101", "Dynamodb hostname")
	flag.IntVar(&dbPort, "dbport", 8000, "Dynamodb port")
	flag.StringVar(&dbRegion, "dbregion", "us-east-1", "Dynamodb region")
	flag.IntVar(&consulPort, "consulport", 8500, "Consul node port")
	flag.Parse()
}

func main() {

	dao.ConfigureDb(dbHost, dbRegion, dbPort)
	root := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servicePort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()

	consul.RegisterService("UserService", &consul.Config{
		ServiceIp:   serviceIP,
		ServicePort: servicePort,
		NodeIp:      consulHost,
		NodePort:    consulPort})

	var svc server.UserService
	{
		svc = &service.UserService{}
	}
	pb.RegisterUserServiceServer(g, newGRPCBinding(root, svc))
	g.Serve(lis)
}
