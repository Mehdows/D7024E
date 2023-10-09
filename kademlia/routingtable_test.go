package d7024e

import (
	"testing"
)

func TestNewRoutingTable(t *testing.T) {
	me := NewContact(NewRandomKademliaID(), "localhost:8000")
	rt := NewRoutingTable(me)

	if rt == nil {
		t.Errorf("NewRoutingTable returned nil")
	}

	if rt.me != me {
		t.Errorf("NewRoutingTable did not set the correct 'me' field")
	}

	for i := 0; i < IDLength*8; i++ {
		if rt.buckets[i] == nil {
			t.Errorf("NewRoutingTable did not initialize bucket %d", i)
		}
	}
}
func TestRoutingTable_FindClosestContacts(t *testing.T) {
	me := NewContact(NewRandomKademliaID(), "localhost:8000")
	rt := NewRoutingTable(me)

	// Add some contacts to the routing table
	for i := 0; i < 10; i++ {
		id := NewRandomKademliaID()
		rt.AddContact(NewContact(id, "localhost:8000"))
	}

	// Find closest contacts to a target ID
	target := NewRandomKademliaID()
	contacts := rt.FindClosestContacts(target, 5)

	// Check that the correct number of contacts were returned
	if len(contacts) != 5 {
		t.Errorf("FindClosestContacts returned %d contacts, expected %d", len(contacts), 5)
	}
}

func TestRoutingTable_FindClosestContacts2(t *testing.T) {
}

func TestRoutingTable_AddContact(t *testing.T) {
	me := NewContact(NewRandomKademliaID(), "localhost:8000")
	rt := NewRoutingTable(me)

	// Add a contact to the routing table
	id := NewRandomKademliaID()
	rt.AddContact(NewContact(id, "localhost:8000"))

	// Check that the contact was added to the correct bucket
	bucketIndex := rt.getBucketIndex(id)
	bucket := rt.buckets[bucketIndex]
	if bucket.Len() != 1 {
		t.Errorf("Contact was not added to the correct bucket")
	}
}
