package networking

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type Netowrk struct {
	ctx  context.Context
	port int
	host host.Host
}

func NewNetwork(ctx context.Context, port int) *Netowrk {
	host, err := libp2p.New(
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port),
			fmt.Sprintf("/ip6/::/tcp/%d", port),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Host created. Listening on: %s\n", host.Addrs())
	return &Netowrk{
		ctx:  ctx,
		port: port,
		host: host,
	}
}

func (n *Netowrk) GetHost() host.Host {
	return n.host
}

func (n *Netowrk) ConnectToPeer(addr string) (err error) {
	mulAddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(mulAddr)
	if err != nil {
		return
	}

	n.host.Peerstore().AddAddr(peerInfo.ID, mulAddr, peerstore.PermanentAddrTTL)
	if err := n.host.Connect(n.ctx, *peerInfo); err != nil {
		return err
	}

	log.Printf("Connected to peer: %s\n", addr)
	return nil
}

func (n *Netowrk) StartSimpleProtocol(protocolID protocol.ID) {
	n.host.SetStreamHandler(protocolID, streamHandler)
}

func (n *Netowrk) SendMessage(peerID peer.ID, protocolID protocol.ID, msg string) (err error) {
	stream, err := n.host.NewStream(n.ctx, peerID, protocolID)
	if err != nil {
		return
	}
	defer stream.Close()

	_, err = stream.Write([]byte(msg))
	if err != nil {
		return err
	}

	log.Printf("Message sent: %s\n", msg)
	return nil
}

func streamHandler(stream network.Stream) {
	log.Println("New stream opened")
	defer stream.Close()

	// Handle incoming data
	buf := make([]byte, 256)
	n, err := stream.Read(buf)
	if err != nil {
		log.Printf("Error reading stream: %s\n", err)
		return
	}
	log.Printf("Received message: %s\n", string(buf[:n]))
}
