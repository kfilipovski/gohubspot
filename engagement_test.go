package gohubspot

import (
	"testing"
)

func TestEngagementsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	_, err := testclient.Engagements.GetAll()
	if err != nil {
		t.Error(err)
	}
}
