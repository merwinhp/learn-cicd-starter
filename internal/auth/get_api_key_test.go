package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid api key",
			headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
			wantErr: false,
		},
		{
			name:    "missing authorization header",
			headers: http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "malformed header wrong scheme",
			headers: http.Header{"Authorization": []string{"Bearer my-secret-key"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "malformed header missing key",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
