package gohubspot

import (
	"testing"
)

func TestCompaniesService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	_, err := testclient.Companies.GetAll()
	if err != nil {
		t.Error(err)
	}
}
