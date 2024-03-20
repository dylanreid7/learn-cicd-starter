package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

// Define custom errors
var ErrMalformedAuthHeader = errors.New("malformed authorization header")

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "Valid Authorization header",
			headers: http.Header{"Authorization": {"ApiKey -api-key"}},
			want:    "my-api-key",
			wantErr: nil,
		},
		{
			name:    "No Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "Malformed Authorization header",
			headers: http.Header{"Authorization": {"Bearer token"}},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
