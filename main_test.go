package testmainsample

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before")
	_ =m.Run()
	fmt.Println("after")

}