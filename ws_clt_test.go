package gosf_test

import (
	"testing"

	"github.com/dwdwow/gosf"
)

func TestWsClient(t *testing.T) {
	clt := gosf.NewWsClient("ws://localhost:8442/ws", nil)
	err := clt.Start()
	if err != nil {
		t.Fatal(err)
	}
	select {}
}
