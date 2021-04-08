package main

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sssergei/microservice/internal/server"
	"github.com/sssergei/microservice/proto/reminder/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverCert, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	reminder.RegisterReminderServiceServer(grpcServer, new(server.MyServer))

	clientCert, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
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
	if err = sservice.RegisterReminderServiceHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	http.ListenAndServeTLS(":8080", "../cert/server.crt", "../cert/server.key", httpGrpcRouter(grpcServer, router))
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
