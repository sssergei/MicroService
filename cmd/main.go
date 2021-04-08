package main

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	pb "github.com/sssergei/userservice/proto/userservice/v1/myservice.proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverCert, err := credentials.NewServerTLSFromFile("../cert/baseserver.crt", "../cert/baseserver.key")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	pb.RegisterReminderServiceServer(grpcServer, new(pb.MyServer))

	clientCert, err := credentials.NewClientTLSFromFile("../cert/baseserver.crt", "")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	router := runtime.NewServeMux()
	if err = pb.RegisterReminderServiceHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	http.ListenAndServeTLS(":8080", "../cert/baseserver.crt", "../cert/baseserver.key", httpGrpcRouter(grpcServer, router))
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
