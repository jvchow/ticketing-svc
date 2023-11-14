package service

import (
	"context"
	"reflect"
	"testing"
	train "ticketing-svc/proto"
)

func Test_server_PurchaseTicket(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *train.PurchaseRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *train.Receipt
		wantErr bool
	}{
		{
			name: "success - purchase ticket in section A",
			args: args{
				ctx: context.Background(),
				in: &train.PurchaseRequest{
					From: "London",
					To:   "France",
					User: &train.User{
						FirstName: "John",
						LastName:  "Doe",
						Email:     "john.doe@example.com",
					},
					PricePaid: 20.0,
				},
			},
			want: &train.Receipt{
				From:      "London",
				To:        "France",
				User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
				PricePaid: 20.0,
				Seat:      "A-0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer()

			got, err := s.PurchaseTicket(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.PurchaseTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.PurchaseTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetReceipt(t *testing.T) {
	s := NewServer()

	s.PurchaseTicket(context.TODO(), &train.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
	})

	type args struct {
		ctx context.Context
		in  *train.UserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *train.Receipt
		wantErr bool
	}{
		{
			name: "success - get receipt for user",
			args: args{
				ctx: context.Background(),
				in:  &train.UserRequest{Email: "john.doe@example.com"},
			},
			want: &train.Receipt{
				From:      "London",
				To:        "France",
				User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
				PricePaid: 20.0,
				Seat:      "A-0",
			},
			wantErr: false,
		},
		{
			name: "fail - get receipt for user - not found",
			args: args{
				ctx: context.Background(),
				in:  &train.UserRequest{Email: "test@example.com"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetReceipt(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetReceipt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetReceipt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ViewSeats(t *testing.T) {
	s := NewServer()

	s.PurchaseTicket(context.TODO(), &train.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
	})

	type args struct {
		ctx context.Context
		in  *train.SectionRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *train.SeatResponse
		wantErr bool
	}{
		{
			name: "success - view seats in section A",
			args: args{
				ctx: context.Background(),
				in:  &train.SectionRequest{Section: "A"},
			},
			want: &train.SeatResponse{
				Users: []*train.User{
					{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ViewSeats(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.ViewSeats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.ViewSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_RemoveUser(t *testing.T) {
	s := NewServer()

	s.PurchaseTicket(context.TODO(), &train.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
	})

	type args struct {
		ctx context.Context
		in  *train.UserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *train.StatusResponse
		wantErr bool
	}{
		{
			name: "success - remove user",
			args: args{
				ctx: context.Background(),
				in:  &train.UserRequest{Email: "john.doe@example.com"},
			},
			want:    &train.StatusResponse{Message: "User removed successfully"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RemoveUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.RemoveUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ModifySeat(t *testing.T) {
	s := NewServer()

	s.PurchaseTicket(context.TODO(), &train.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &train.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
	})

	type args struct {
		ctx context.Context
		in  *train.ModifySeatRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *train.StatusResponse
		wantErr bool
	}{
		{
			name: "success - modify seat",
			args: args{
				ctx: context.Background(),
				in:  &train.ModifySeatRequest{Email: "john.doe@example.com"},
			},
			want:    &train.StatusResponse{Message: "Seat modified successfully"},
			wantErr: false,
		},
		{
			name: "fail - modify seat - not found",
			args: args{
				ctx: context.Background(),
				in:  &train.ModifySeatRequest{Email: "invalid@example.com"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ModifySeat(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.ModifySeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.ModifySeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
