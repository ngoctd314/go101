package dbconn

import (
	"strconv"
	"strings"
	"sync"

	"github.com/aerospike/aerospike-client-go"
)

// AerospikeCrud ...
type AerospikeCrud interface {
	IsConnected() bool
}

var (
	// _aerospikeConnMgmt manages aerospike connection by host
	_aerospikeConnMgmt = make(map[string]AerospikeCrud)
	_aerospikeConnMu   sync.RWMutex
)

// GetAerospikeConn return connection to aerospike
func GetAerospikeConn(hosts string) AerospikeCrud {
	_aerospikeConnMu.RLock()
	conn, ok := _aerospikeConnMgmt[hosts]
	// already connected and ready to take to the database => return immediatedly
	if ok && conn != nil && conn.IsConnected() {
		return conn
	}
	_aerospikeConnMu.RUnlock()

	client, err := aerospike.NewClientWithPolicyAndHost(nil, newAerHost(hosts)...)
	if err != nil {
		panic(err)
	}
	if !client.IsConnected() {
		panic("aerospike client isn't ready to talk to the database server")
	}

	// store client instance to storage manager
	_aerospikeConnMu.Lock()
	_aerospikeConnMgmt[hosts] = client
	_aerospikeConnMu.Unlock()

	return client
}

// conv name:port,name:port => []*aerospike.Host
func newAerHost(hosts string) []*aerospike.Host {
	listHostRaw := strings.Split(hosts, ",")
	listHost := make([]*aerospike.Host, 0)

	for i := 0; i < len(listHostRaw); i++ {
		hostRaw := strings.Split(listHostRaw[i], ":")
		name, portRaw := hostRaw[0], hostRaw[1]
		port, err := strconv.Atoi(portRaw)
		if err != nil {
			panic(err)
		}

		listHost = append(listHost, aerospike.NewHost(name, port))
	}

	return listHost
}
