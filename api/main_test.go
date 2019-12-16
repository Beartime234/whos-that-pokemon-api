package main

import (
	"context"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	type args struct {
		request Request
		ctx     context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		{
			name:    "HandlerTest1",
			args:    args{
				request: Request{ID:3},
				ctx:     nil,
			},
			want:    Response{
				StatusCode:        200,
				Headers:           map[string]string{
					"Content-Type":           "application/json",
					"X-WTP-Func-Reply": "api-Handler",
				},
				MultiValueHeaders: nil,
				Body:              `{"message":"Your ID is 3"}`,
				IsBase64Encoded:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}