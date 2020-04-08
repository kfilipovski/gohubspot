package gohubspot

import (
	"fmt"
	"testing"
	"time"
)

func TestContactService_Create(t *testing.T) {
	setup()
	defer teardown()

	var now UnixTime
	now.Time = time.Now().UTC()

	contact := Properties{}

	contact.AddProperty("Email", "ion.suruceanu@gapsquare.com")
	contact.AddProperty("firstname", "Ion")
	contact.AddProperty("lastname", "Suruceanu")
	contact.AddProperty("lastlogindatetime", &now)

	fmt.Println(contact)

	_, err := testclient.Contacts.Create(contact)
	if err != nil {
		t.Error(err)
	}
}

func TestContactService_GetById(t *testing.T) {
	setup()
	defer teardown()

	// demo id
	id := 9332174

	_, err := testclient.Contacts.GetById(id)
	if err != nil {
		t.Error(err)
	}
}

func TestContactService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	_, err := testclient.Contacts.GetAll()
	if err != nil {
		t.Error(err)
	}
}
