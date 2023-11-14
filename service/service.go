package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	train "ticketing-svc/proto"
)

// This example assumes each section (A and B) has 10 seats each.
const seatsPerSection = 10

// Seat represents a seat on the train.
type Seat struct {
	Section string
	Number  int
}

// server is used to implement train.TicketServiceServer.
type server struct {
	train.UnimplementedTicketServiceServer
	mu        sync.Mutex // protects the following fields
	tickets   map[string]*train.Receipt
	seats     map[string]Seat
	nextSeatA int // Next available seat number in Section A
	nextSeatB int // Next available seat number in Section B
}

// NewServer creates a TicketService server with an initialized in-memory store.
func NewServer() *server {
	return &server{
		mu:      sync.Mutex{},
		tickets: make(map[string]*train.Receipt),
		seats:   make(map[string]Seat),
	}
}

// PurchaseTicket creates a ticket purchase entry.
func (s *server) PurchaseTicket(ctx context.Context, in *train.PurchaseRequest) (*train.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// assign a seat
	seat := s.assignSeat(in.User.Email)

	receipt := &train.Receipt{
		From:      in.From,
		To:        in.To,
		User:      in.User,
		PricePaid: in.PricePaid,
		Seat:      fmt.Sprintf("%s-%d", seat.Section, seat.Number),
	}
	s.tickets[in.User.Email] = receipt

	return receipt, nil
}

// GetReceipt retrieves the receipt details for a user.
func (s *server) GetReceipt(ctx context.Context, in *train.UserRequest) (*train.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, ok := s.tickets[in.Email]
	if !ok {
		return nil, fmt.Errorf("no ticket found for email: %s", in.Email)
	}
	return receipt, nil
}

// assignSeat assigns a seat to a user, alternating between sections A and B.
func (s *server) assignSeat(email string) Seat {
	var section string
	var number int

	// Assign to Section A or B depending on availability.
	if s.nextSeatA <= s.nextSeatB && s.nextSeatA <= seatsPerSection {
		section = "A"
		number = s.nextSeatA
		s.nextSeatA++
	} else if s.nextSeatB <= seatsPerSection {
		section = "B"
		number = s.nextSeatB
		s.nextSeatB++
	} else {
		log.Fatalf("No more seats available.")
	}

	// store the seat assignment
	s.seats[email] = Seat{Section: section, Number: number}

	return Seat{Section: section, Number: number}
}

// ViewSeats lists all the users in a requested section.
func (s *server) ViewSeats(ctx context.Context, in *train.SectionRequest) (*train.SeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var usersInRequestedSection []*train.User
	for email, receipt := range s.tickets {
		seat, ok := s.seats[email]
		if ok && seat.Section == in.Section {
			usersInRequestedSection = append(usersInRequestedSection, receipt.User)
		}
	}

	return &train.SeatResponse{Users: usersInRequestedSection}, nil
}

// RemoveUser removes a user from the train.
func (s *server) RemoveUser(ctx context.Context, in *train.UserRequest) (*train.StatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tickets, in.Email)

	return &train.StatusResponse{Message: "User removed successfully"}, nil
}

// ModifySeat changes a user's seat.
func (s *server) ModifySeat(ctx context.Context, in *train.ModifySeatRequest) (*train.StatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, ok := s.tickets[in.Email]
	if !ok {
		return nil, fmt.Errorf("no ticket found for email: %s", in.Email)
	}

	receipt.Seat = in.NewSeat
	return &train.StatusResponse{Message: "Seat modified successfully"}, nil
}
