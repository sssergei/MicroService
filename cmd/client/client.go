package main

import (
	"context"
	"time"

	pb "github.com/sssergei/userservice/proto/userservice/v1/myservice.proto"

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

	clientCert, err := credentials.NewClientTLSFromFile("../../cert/baseserver.crt", "")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	reminderConn, err := grpc.DialContext(ctx, "localhost:8080",
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	reminderClient := pb.NewReminderServiceClient(reminderConn)
	fiveSeconds, _ := ptypes.TimestampProto(time.Now().Add(5 * time.Second))
	resp, err := reminderClient.ScheduleReminder(ctx,
		&pb.ScheduleReminderRequest{
			When: fiveSeconds,
		})
	if err != nil {
		log.Fatalln("Failed to schedule a reminder: ", err)
	}
	log.Infof("Reminder have been successfully scheduled. New  reminder id is %s", resp.GetId())
}
