package server

import (
	"log"
	"net"
	"os"
)

const (
	// BDDP protocol version used by the server.
	// Until the server supports multiple versions,
	// client version should match server version.
	Version = "1"

	// Error logger params
	LogPrefix = "BDDP: "
	LogFlags  = 0
)

// Method handler type
type MHandler func(ctx MContext)

// Use an interface so if it's required to mock a bddp server
// for tests it can be done easily.
type Server interface {
	// Returns a channel from where the user can get errors.
	// Make sure the error handler doesn't block the goroutine.
	Errors() (errCh chan error)

	// Registers a bddp method handler with given name.
	// If a handler already exists, it will be replaced.
	Method(name string, handler MHandler)

	// Stops listening for new connections
	Close() (err error)

	// Listen method starts listening on pre set tcp address.
	// The address must be provided when creating the struct.
	Listen() (err error)
}

type server struct {
	address string

	returnErr bool
	errChan   chan error

	methods  map[string]MHandler
	listener net.Listener
	closed   bool

	logger *log.Logger
}

// Create a new Server which will connect to provided
// address. example: `server.New(":3000")`
func New(address string) (s Server) {
	return &server{
		address: address,
		errChan: make(chan error),
		methods: make(map[string]MHandler),
		logger:  log.New(os.Stderr, LogPrefix, LogFlags),
	}
}

func (s *server) Errors() (errCh chan error) {
	s.returnErr = true
	return s.errChan
}

func (s *server) Method(name string, handler MHandler) {
	s.methods[name] = handler
}

func (s *server) Close() (err error) {
	if s.closed {
		return nil
	}

	s.closed = true

	err = s.listener.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *server) Listen() (err error) {
	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	for !s.closed {
		conn, err := s.listener.Accept()
		if err != nil {
			s.handleErr(err)
			continue
		}

		ses := newSession(conn, s)
		go s.handleSession(ses)
	}

	return nil
}

// `ses.process()` will block until it disconnects
// TODO: retain session data across reconnects
func (s *server) handleSession(ses *session) {
	err := ses.process()
	if err != nil {
		s.handleErr(err)
	}
}

// If user receives errors with the channel, send errors
// to the user. Otherwise log them to console STDERR.
func (s *server) handleErr(err error) {
	if s.returnErr {
		s.errChan <- err
	} else {
		s.logger.Println(err)
	}
}
