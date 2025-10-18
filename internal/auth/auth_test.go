package auth

import (
	"testing"
	"net/http"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name string
		input http.Header
		want string
		wantErr error
		setup func(*http.Header)
	}{
		{name: "empty auth", input: http.Header{}, want: "", wantErr: ErrNoAuthHeaderIncluded},
		{name: "malformed", input: http.Header{}, want: "", wantErr: ErrMalformedAuthHeader, setup: func(h *http.Header) { h.Add("Authorization", "apple") }},
		{name: "happy path", input: http.Header{}, want: "abcdef", wantErr: nil, setup: func(h *http.Header) { h.Add("Authorization", "ApiKey abcdef") }},
	}

	for _, tc := range tests {

		if tc.setup != nil {
			tc.setup(&tc.input)
		}

		got, err := GetAPIKey(tc.input)

		if err != tc.wantErr {
			t.Fatalf("Test:%s\nInput:%+v\nExpected error: %v; Got: %v", tc.name, tc.input, tc.wantErr, err)
		}

		if got != tc.want {
			t.Fatalf("Expected output: %s; Got: %s", tc.want, got)
		}
	}
}
