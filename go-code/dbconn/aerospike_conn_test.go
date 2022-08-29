package dbconn

import (
	"reflect"
	"testing"

	"github.com/aerospike/aerospike-client-go"
)

func Test_newAerHost(t *testing.T) {
	type args struct {
		hosts string
	}
	tests := []struct {
		name string
		args args
		want []*aerospike.Host
	}{
		{
			name: "single host",
			args: args{
				hosts: "127.0.0.1:8080",
			},
			want: []*aerospike.Host{
				aerospike.NewHost("127.0.0.1", 8080),
			},
		},
		{
			name: "cluster",
			args: args{
				hosts: "127.0.0.1:8080,127.0.0.1:8081",
			},
			want: []*aerospike.Host{
				aerospike.NewHost("127.0.0.1", 8080),
				aerospike.NewHost("127.0.0.1", 8081),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAerHost(tt.args.hosts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAerHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
