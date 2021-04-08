package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	clientCert, err := credentials.NewClientTLSFromFile("../../cert/server.crt", "")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	reminderConn, err := grpc.DialContext(ctx, "localhost:8080",
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	reminderClient := sservice.NewReminderServiceClient(reminderConn)
	fiveSeconds, _ := ptypes.TimestampProto(time.Now().Add(5 * time.Second))
	resp, err := reminderClient.ScheduleReminder(ctx,
		&sservice.ScheduleReminderRequest{
			When: fiveSeconds,
		})
	if err != nil {
		log.Fatalln("Failed to schedule a reminder: ", err)
	}
	log.Infof("Reminder have been successfully scheduled. New  reminder id is %s", resp.GetId())
}
