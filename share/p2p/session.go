package p2p

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"

	"github.com/celestiaorg/go-libp2p-messenger/serde"
)

type Session struct {
	writeTimeout, readTimeout time.Duration
	Stream                    network.Stream

	// optionally session level middleware.
}

func NewSession(
	stream network.Stream,
	writeTimeout, readTimeout time.Duration) *Session {
	return &Session{
		writeTimeout: writeTimeout,
		readTimeout:  readTimeout,
		Stream:       stream,
	}
}

func (s *Session) Read(msg serde.Message) error {
	err := s.Stream.SetReadDeadline(time.Now().Add(s.readTimeout))
	if err != nil {
		log.Debugf("error setting deadline: %s", err)
	}

	_, err = serde.Read(s.Stream, msg)
	return err
}

func (s *Session) Write(msg serde.Message) error {
	err := s.Stream.SetWriteDeadline(time.Now().Add(s.writeTimeout))
	if err != nil {
		log.Debugf("error setting deadline: %s", err)
	}

	_, err = serde.Write(s.Stream, msg)
	return err
}
