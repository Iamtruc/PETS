package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Message struct {
	Party PartyID
	Value uint64
}

type Protocol struct {
	// A dummy Protocol is associated to local parties, a channel and peers.
	*LocalParty
	Chan  		 chan Message
	Peers		 map[PartyID]*Remote

	Input uint64
	Output uint64
}

type Remote struct {
	*RemoteParty
	Chan chan Message
}

func (lp *LocalParty) NewProtocol(input uint64) *Protocol {
	cep := new(Protocol)
	cep.LocalParty = lp
	cep.Chan = make(chan Message, 32)
	cep.Peers = make(map[PartyID]*Remote, len(lp.Peers))
	for i, rp := range lp.Peers {
		cep.Peers[i] = &Remote{
			RemoteParty:  rp,
			Chan:         make(chan Message, 32),
		}
	}


	cep.Input = input
	cep.Output = input
	return cep
}

func (cep *Protocol) BindNetwork(nw *TCPNetworkStruct) {
	for partyID, conn := range nw.Conns {

		if partyID == cep.ID {
			continue
		}

		rp := cep.Peers[partyID]

		// Receiving loop from remote
		// The thing below is a lambda function
		go func(conn net.Conn, rp *Remote) {
			for {
				var id, val uint64
				var err error
				err = binary.Read(conn, binary.BigEndian, &id)
				check(err)
				err = binary.Read(conn, binary.BigEndian, &val)
				check(err)
				// We check that there are no errors in the value and id of the connection. If there are no errors, then we create a new message and send it on the channel.
				msg := Message{
					Party: PartyID(id),
					Value: val,
				}
				//fmt.Println(cep, "receiving", msg, "from", rp)
				cep.Chan <- msg
			}
		}(conn, rp)

		// Sending loop of remote
		// Another lambda function
		go func(conn net.Conn, rp *Remote) {
			var m Message
			var open = true
			for open {
				m, open = <- rp.Chan
				// Here, we receive a message from the channel.
				//fmt.Println(cep, "sending", m, "to", rp)
				check(binary.Write(conn, binary.BigEndian, m.Party))
				check(binary.Write(conn, binary.BigEndian, m.Value))
			}
		}(conn, rp)
	}
}

func (cep *Protocol) Run() {

	fmt.Println(cep, "is running")

	for _, peer := range cep.Peers {
		if peer.ID != cep.ID {
			peer.Chan <- Message{cep.ID, cep.Input}
		}
	}
	// We get all the messages from the different peers
	received := 0
	for m := range cep.Chan {
		fmt.Println(cep, "received message from", m.Party, ":", m.Value)
		cep.Output += m.Value
		received++
		if received == len(cep.Peers)-1 {
			close(cep.Chan)
		}
	}
	// We enumerate all the messages that we have gottent and close the channel only once we have all of them.
	if cep.WaitGroup != nil {
		cep.WaitGroup.Done()
	}
}