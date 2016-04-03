package p2p


import (
	"net"
	"time"
)


type File struct {
	Path string
	Hash string
}


type Node struct {
	// list of file hashes it has fully downloaded
	Files []File
}





func Listen(conns chan *net.Conn, timeout int) error {
	// , quit chan struct{}
	l, err := net.ListenTCP("tcp", nil)
	if err != nil { return err }
	defer l.Close()

	if timeout > 0 {
		l.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	}


	for {
		conn, err := l.Accept()
		if err != nil { return err }
		conns <- &conn
	}


	// for {
	// 	select {
	// 	case <- quit:
	// 		return nil // TODO cleanup
	// 	default:
	// 		conn, err := l.Accept()
	// 		if err != nil { return err }
	// 		conns <- &conn
	// 	}	
	// }



	// go func() {
	// 	select {
	// 	case <- quit:
	// 		l.Close()
	// 	}
	// }()

	// // defer l.Close()
	// // println(l.Addr().String())

	// return l.Addr().String(), nil
}





