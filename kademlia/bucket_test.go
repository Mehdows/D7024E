package kademlia

import (
	"testing"
)

func Test_AddContact(t *testing.T) {
	bucket := newBucket()
	contact := NewContact(NewKademliaID("0000000000000000000000000000000000000001"), "localhost:8000")
	contact2 := NewContact(NewKademliaID("0000000000000000000000000000000000000002"), "localhost:8000")

	bucket.AddContact(contact)
	bucket.AddContact(contact2)
	bucket.AddContact(contact)

	if bucket.list.Front().Value != contact {
		t.Errorf("Contact was not added to the front of the bucket")
	}

}
