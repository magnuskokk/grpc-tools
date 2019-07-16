package main_test

import (
	"fmt"
	"serverapp/pkg/testconn"
	"testing"
)

func TestPackageCanBeImportedFromAnotherLocalModule(t *testing.T) {
	bufnet := testconn.NewBufNet()
	if fmt.Sprintf("%T", bufnet) != "*testconn.BufNet" {
		t.Fatal("Expected *testconn.BufNet")
	}
}
