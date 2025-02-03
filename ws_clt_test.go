package gosf_test

import (
	"testing"

	"github.com/dwdwow/gosf"
)

func TestWsClient(t *testing.T) {
	handler := gosf.NewSimpleWsMsgHandler(nil)
	clt := gosf.NewWsClient("ws://localhost:8442", handler, nil)
	err := clt.Start()
	if err != nil {
		t.Fatal(err)
	}
	select {}
}
