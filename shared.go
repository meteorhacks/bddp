package bddp

import "errors"

const (
	// BDDP protocol version used by the client.
	// Until the client supports multiple versions,
	// client version should match server version.
	Version = "1"

	// Error logger params
	LogPrefix = "BDDP: "
	LogFlags  = 0
)

var (
	ErrInvalidMessage = errors.New("invalid message type")
)
