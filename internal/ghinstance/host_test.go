package ghinstance

import (
	"testing"
)

func TestIsEnterprise(t *testing.T) {
	tests := []struct {
		host string
		want bool
	}{
		{
			host: "github.com",
			want: false,
		},
		{
			host: "api.github.com",
			want: false,
		},
		{
			host: "ghe.io",
			want: true,
		},
		{
			host: "example.com",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			if got := IsEnterprise(tt.host); got != tt.want {
				t.Errorf("IsEnterprise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeHostname(t *testing.T) {
	tests := []struct {
		host string
		want string
	}{
		{
			host: "GitHub.com",
			want: "github.com",
		},
		{
			host: "api.github.com",
			want: "github.com",
		},
		{
			host: "ssh.github.com",
			want: "github.com",
		},
		{
			host: "upload.github.com",
			want: "github.com",
		},
		{
			host: "GHE.IO",
			want: "ghe.io",
		},
		{
			host: "git.my.org",
			want: "git.my.org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			if got := NormalizeHostname(tt.host); got != tt.want {
				t.Errorf("NormalizeHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphQLEndpoint(t *testing.T) {
	tests := []struct {
		host string
		want string
	}{
		{
			host: "github.com",
			want: "https://api.github.com/graphql",
		},
		{
			host: "ghe.io",
			want: "https://ghe.io/api/graphql",
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			if got := GraphQLEndpoint(tt.host); got != tt.want {
				t.Errorf("GraphQLEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRESTPrefix(t *testing.T) {
	tests := []struct {
		host string
		want string
	}{
		{
			host: "github.com",
			want: "https://api.github.com/",
		},
		{
			host: "ghe.io",
			want: "https://ghe.io/api/v3/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			if got := RESTPrefix(tt.host); got != tt.want {
				t.Errorf("RESTPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
