package models

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestTransaction_SetTransactionAmount(t *testing.T) {
	type fields struct {
		OperationID int
		Amount      float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "should return positive with operationID equal 4",
			fields: fields{
				OperationID: 4,
				Amount:      100.0,
			},
			want: 100.0,
		},
		{
			name: "should return negative with operationID equal 3",
			fields: fields{
				OperationID: 3,
				Amount:      100.0,
			},
			want: -100.0,
		},
		{
			name: "should return negative with operationID equal 2",
			fields: fields{
				OperationID: 2,
				Amount:      100.0,
			},
			want: -100.0,
		},
		{
			name: "should return negative with operationID equal 1",
			fields: fields{
				OperationID: 1,
				Amount:      100.0,
			},
			want: -100.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				OperationID: tt.fields.OperationID,
				Amount:      tt.fields.Amount,
			}

			tr.SetTransactionAmount()

			if tr.Amount != tt.want {
				t.Errorf("amount: %v, want: %v", tr.Amount, tt.want)
			}
		})
	}
}

func TestTransaction_ValidateBody(t *testing.T) {
	type fields struct {
		AccountID   uuid.UUID
		OperationID int
		Amount      float64
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "should return no error with all fields its ok",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: 1,
				Amount:      100,
			},
			want: nil,
		},
		{
			name: "should return error with no AccountID ID",
			fields: fields{
				AccountID:   uuid.Nil,
				OperationID: 1,
				Amount:      100,
			},
			want: fmt.Errorf("'account_id' cannot be empty"),
		},
		{
			name: "should return error with OperationID bigger then 4",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: 5,
				Amount:      100,
			},
			want: fmt.Errorf("'operation_id' must be a number between 1 and 4"),
		},
		{
			name: "should return error with OperationID lower then 0",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: -1,
				Amount:      100,
			},
			want: fmt.Errorf("'operation_id' must be a number between 1 and 4"),
		},
		{
			name: "should return error with OperationID equal 0",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: 0,
				Amount:      100,
			},
			want: fmt.Errorf("'operation_id' must be a number between 1 and 4"),
		},
		{
			name: "should return error with Amount equal 0",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: 1,
				Amount:      0,
			},
			want: fmt.Errorf("'amount' must be a number greater than zero"),
		},
		{
			name: "should return error with Amount negative",
			fields: fields{
				AccountID:   uuid.New(),
				OperationID: 1,
				Amount:      -100,
			},
			want: fmt.Errorf("'amount' must be a number greater than zero"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				AccountID:   tt.fields.AccountID,
				OperationID: tt.fields.OperationID,
				Amount:      tt.fields.Amount,
			}

			err := tr.ValidateBody()
			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("error: %v, want: %v", err, tt.want)
				}
			}
		})
	}
}
