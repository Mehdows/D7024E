package kademlia

import "testing"


func TestNewKademliaNode(t *testing.T) {
	address := "localhost:8000"
	kademlia := NewKademliaNode(address)

	// Check that the Kademlia node was created with the correct address
	if kademlia.me.Address != address {
		t.Errorf("Expected address %s, but got %s", address, kademlia.me.Address)
	}

	// Check that the Kademlia node was created with an empty dictionary
	if len(kademlia.dictionary) != 0 {
		t.Error("Expected an empty dictionary, but got a non-empty one")
	}

	// Check that the Kademlia node was created with the correct replication factor
	if kademlia.replicationFactor != 1 {
		t.Errorf("Expected replication factor 1, but got %d", kademlia.replicationFactor)
	}

	// Check that the Kademlia node was created with the correct k value
	if kademlia.k != 1 {
		t.Errorf("Expected k value 1, but got %d", kademlia.k)
	}
}

