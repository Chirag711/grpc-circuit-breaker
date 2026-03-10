package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-circuit-breaker/grpc-circuit-breaker/pb"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "PaymentService",
		MaxRequests: 3,
		Timeout:     5 * time.Second,
	})

	for i := 1; i <= 10; i++ {

		result, err := cb.Execute(func() (interface{}, error) {

			return client.ProcessPayment(context.Background(), &pb.PaymentRequest{
				OrderId: fmt.Sprintf("ORD-%d", i),
				Amount:  100,
			})

		})

		if err != nil {

			fmt.Println("Circuit Breaker Triggered:", err)

		} else {

			res := result.(*pb.PaymentResponse)
			fmt.Println("Payment Status:", res.Status)

		}

		time.Sleep(1 * time.Second)
	}
}
