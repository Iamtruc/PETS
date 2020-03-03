package main

import (
	"fmt"
	"sync"
)

type PartyID uint64

type Party struct {
	ID   PartyID
	Addr string
	// Each party is composed of its ID and an address.
}


type LocalParty struct {
	// A LocalParty has a Party and some RemoteParty attached to it
	Party
	*sync.WaitGroup
	Peers map[PartyID]*RemoteParty
}

func check(err error) {
	// Just checks if there is an error.
	if err != nil {
		panic(err.Error())
	}
}

func NewLocalParty(id PartyID, peers map[PartyID]string) (*LocalParty, error) {
	// Create a new local party from the peers map 
	p := &LocalParty{}
	p.ID = id

	p.Peers = make(map[PartyID]*RemoteParty, len(peers))
	// Assigns avery peer to a RemoteParty associated with the LocalParty
	p.Addr = peers[id]

	var err error
	for pId, pAddr := range peers {
		p.Peers[pId], err = NewRemoteParty(pId, pAddr)
		// Check that we can attach every peer to our RemoteParty
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (lp *LocalParty) String() string {
	// Print the party number
	return fmt.Sprintf("party-%d", lp.ID)
}


type RemoteParty struct {
	Party
}

func (rp *RemoteParty) String() string {
	return fmt.Sprintf("party-%d", rp.ID)
}

func NewRemoteParty(id PartyID, addr string) (*RemoteParty, error) {
	p := &RemoteParty{}
	p.ID = id
	p.Addr = addr
	return p, nil
}