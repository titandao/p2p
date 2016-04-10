package p2p


import (
	"os"
	"fmt"
	// "net"
	"time"


	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
)




func Listen() {
	nodekey, _ := crypto.GenerateKey()


	println("Starting p2p server...")

	srv := p2p.Server{
	    MaxPeers:   10,
	    PrivateKey: nodekey,
	    Name:       "my node name",
	    ListenAddr: ":30300",
	    Protocols:  []p2p.Protocol{},
	}

	if err := srv.Start(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    setInterval(func() {
    	println("ticking...")	
	})

    select {}
}



func setInterval(f func()) chan struct{} {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
	    for {
	       select {
	        case <- ticker.C:
	            // do stuff
	            f()
	        case <- quit:
	            ticker.Stop()
	            return
	        }
	    }
 	}()

 	return quit
}



const messageId = 0

type Message string

func MyProtocol() p2p.Protocol {
    return p2p.Protocol{
        Name:    "MyProtocol",
        Version: 1,
        Length:  1,
        Run:     msgHandler,
    }
}

func msgHandler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
    for {

        msg, err := ws.ReadMsg()
        if err != nil {
            return err
        }

        var myMessage Message
        err = msg.Decode(&myMessage)
        if err != nil {
            // handle decode error
            continue
        }

        switch myMessage {
        case "foo":
            err := p2p.SendItems(ws, messageId, "bar")
            if err != nil {
                return err
            }
        default:
            fmt.Println("recv:", myMessage)
        }
    }

    return nil
}