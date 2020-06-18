package gohubspot

import (
	"testing"
)

func TestEngagementsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	engagements, err := testclient.Engagements.GetAll()
	if err != nil {
		t.Error(err)
	}

	t.Log(*engagements[1])
}
