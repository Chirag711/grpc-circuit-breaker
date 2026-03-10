package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "grpc-circuit-breaker/grpc-circuit-breaker/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {

	fmt.Println("Processing payment for order:", req.OrderId)

	// simulate random failure
	if rand.Intn(2) == 0 {
		return nil, fmt.Errorf("bank service unavailable")
	}

	return &pb.PaymentResponse{
		Status: "Payment Successful",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, &server{})

	fmt.Println("Payment Server running...")

	grpcServer.Serve(lis)
}
