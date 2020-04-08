package gohubspot

import (
	"testing"
)

func TestDealService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	_, err := testclient.Deals.GetAll()
	if err != nil {
		t.Error(err)
	}
}
