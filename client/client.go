package main

import (
	"context"
	"log"
	"time"

	"github.com/betaapskaita/beta-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	accountService := proto.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := accountService.AuthenticateByEmailAndPassword(ctx, &proto.User{
		Email:    "dainiauskas@gmail.com",
		Password: "password",
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal(err)
		}
		return
	}

	log.Printf("Response: %+v", res)
}
