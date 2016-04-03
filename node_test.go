package p2p


import (
	"testing"

	// "time"
	"net"
)




func TestListen(t *testing.T) {
	conns := make(chan *net.Conn)
	// q := make(chan struct{})

	err := Listen(conns, 5)

	if err, ok := err.(*net.OpError); ok && err.Timeout() {
		println("timed out")
	} else if err != nil{ t.Fatal(err) }


	// time.Sleep(1 * time.Second)
	// q <- struct{}{}
}