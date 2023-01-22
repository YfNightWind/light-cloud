package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	// Creating UUID Version 4
	u1 := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", u1)

}
