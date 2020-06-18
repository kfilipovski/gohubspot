package gohubspot

import (
	"fmt"
	"testing"
)

func TestCompaniesService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	companies, err := testclient.Companies.GetAll()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(companies[1].Properties)
}
