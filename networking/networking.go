package networking

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/tunnelhop/tunnelhop-types/v2/messages"
)

// DialWithRetries is a helper util to enable retry logic when connecting to remote relay
func DialWithRetries(addr string, retries int, timeout time.Duration, keepAlive bool) (net.Conn, error) {
	var conn net.Conn
	var err error

	for i := 0; i < retries; i++ {
		conn, err = net.DialTimeout("tcp", addr, timeout)
		if err == nil {
			if keepAlive {
				conn.(*net.TCPConn).SetKeepAlive(true)
			}
			return conn, nil
		}

		fmt.Printf("Attempt %d failed: %s\n", i+1, err)
		time.Sleep(time.Second)
	}

	return nil, err
}

// SendMessage is a utility method that should be used to send messages over a
// TCP control tunnel. The payload should be serialized as valid JSON bytes or the
// messages.Message parsing will fail.
func SendMessage(conn net.Conn, message messages.Message) error {
	msgTypeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(msgTypeBytes, uint64(message.MessageType))

	msgLenBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(msgLenBytes, uint64(len(message.Payload)))
	_, err := conn.Write(append(append(msgTypeBytes, msgLenBytes...), []byte(message.Payload)...))
	return err
}

// ReadMessage is a utility method that should be used to read messages off a
// TCP control tunnel. The payload of size N is assumed to be in valid JSON format.
func ReadMessage(conn net.Conn) (messages.Message, error) {
	msgTypeBytes := make([]byte, 8)
	_, err := conn.Read(msgTypeBytes)

	msg := messages.Message{}
	if err != nil {
		return messages.Message{}, err
	}

	msg.MessageType = messages.TunnelMessage(binary.BigEndian.Uint64(msgTypeBytes))
	msgLenBytes := make([]byte, 8)
	_, err = conn.Read(msgLenBytes)
	if err != nil {
		return messages.Message{}, err
	}

	msgLen := int(binary.BigEndian.Uint64(msgLenBytes))
	msgBytes := make([]byte, msgLen)
	_, err = conn.Read(msgBytes)
	if err != nil {
		return messages.Message{}, err
	}

	msg.Payload = msgBytes
	return msg, nil
}
