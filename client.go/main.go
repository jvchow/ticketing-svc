package main

import (
	"context"
	"log"
	"time"

	train "ticketing-svc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := train.NewTicketServiceClient(conn)
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a purchase request
	purchaseReq := &train.PurchaseRequest{
		From: "London",
		To:   "France",
		User: &train.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		PricePaid: 20.0,
	}

	receipt, err := client.PurchaseTicket(ctx, purchaseReq)
	if err != nil {
		log.Fatalf("Could not purchase ticket: %v", err)
	}
	log.Printf("Purchase Receipt: %+v", receipt)

	// Get the receipt for the user
	userReq := &train.UserRequest{Email: "john.doe@example.com"}
	receipt, err = client.GetReceipt(ctx, userReq)
	if err != nil {
		log.Fatalf("Could not get receipt: %v", err)
	}
	log.Printf("Receipt Details: %+v", receipt)

	// View seats in section A
	sectionReq := &train.SectionRequest{Section: "A"}
	seatResp, err := client.ViewSeats(ctx, sectionReq)
	if err != nil {
		log.Fatalf("Could not view seats: %v", err)
	}
	log.Printf("Seats in Section A: %+v", seatResp)

	// Modify the user's seat
	modifySeatReq := &train.ModifySeatRequest{
		Email:   "john.doe@example.com",
		NewSeat: "B1", // Assuming 'B1' is a valid seat representation
	}
	statusResp, err := client.ModifySeat(ctx, modifySeatReq)
	if err != nil {
		log.Fatalf("Could not modify seat: %v", err)
	}
	log.Printf("Modify Seat Response: %+v", statusResp)

	// Finally, remove the user from the train
	removeUserReq := &train.UserRequest{Email: "john.doe@example.com"}
	statusResp, err = client.RemoveUser(ctx, removeUserReq)
	if err != nil {
		log.Fatalf("Could not remove user: %v", err)
	}
	log.Printf("Remove User Response: %+v", statusResp)
}
