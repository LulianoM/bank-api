package models

import (
	"fmt"
	"testing"
)

func TestAccount_IsValidBody(t *testing.T) {
	type fields struct {
		DocumentNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "should return no error with valid DocumentNumber",
			fields: fields{
				DocumentNumber: "12345678900",
			},
			wantErr: nil,
		},
		{
			name: "should return error with empty DocumentNumber",
			fields: fields{
				DocumentNumber: "",
			},
			wantErr: fmt.Errorf("invalid document number"),
		},
		{
			name: "should return error with invalid DocumentNumber",
			fields: fields{
				DocumentNumber: "1234567890",
			},
			wantErr: fmt.Errorf("invalid document number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				DocumentNumber: tt.fields.DocumentNumber,
			}

			err := a.IsValidBody()
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("error: %v, want: %v", err, tt.wantErr)
				}
			}
		})
	}
}
