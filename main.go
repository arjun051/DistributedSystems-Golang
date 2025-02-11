package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"
	"github.com/arjun051/DistributedSystems-Golang/p2p"
)
func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	// Proactively Dial Bootstrap Nodes

	// This is additional functionality, if a server initially refuse to connect we try upto 5 times to reconnect.
for _, node := range nodes {
    go func(node string) {
        for i := 0; i < 5; i++ { // Try 5 times
            log.Println("[DEBUG]", listenAddr, "trying to connect to", node)
            err := tcpTransport.Dial(node)
            if err != nil {
                log.Println("[ERROR] Could not connect to", node, ":", err)
                time.Sleep(2 * time.Second) // Wait and retry
            } else {
                log.Println("[SUCCESS]", listenAddr, "connected to", node)
                break // Exit loop once connected
            }
        }
    }(node)
}

	return s
}
func main() {
	s1 := makeServer(":3000")
	s2 := makeServer(":4000")
	s3 := makeServer(":5000", ":3000", ":4000")

	go func() { log.Fatal(s1.Start()) }()
	time.Sleep(500 * time.Millisecond)
	go func() { log.Fatal(s2.Start()) }()
	time.Sleep(2 * time.Second)

	go s3.Start()
	time.Sleep(2 * time.Second)
	for i := 1; i < 2; i++ {
		key := fmt.Sprintf("picture_%d.png", i)
		data := bytes.NewReader([]byte("my big data file here!"))
		
		err := s3.Store(key, data)
		if err != nil {
			log.Fatal("Error storing file: ", err)
		}
		// time.Sleep(2 * time.Second) // Add delay before deleting

		if err := s3.store.Delete(s3.ID, key); err != nil {
		    log.Fatal(err)
		}
	
		r, err := s3.Get(key)
		if err != nil {
			log.Fatal("Error retrieving file: ", err)
		}
	
		b, err := io.ReadAll(r)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
	
		fmt.Println("Retrieved file content:", string(b))
	}
}