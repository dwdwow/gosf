package gosf_test

import (
	"testing"

	"github.com/dwdwow/gosf"
)

func TestCallbackServer(t *testing.T) {
	sv := gosf.NewCallbackServer("8443", "8442", "8441", nil)
	sv.Serve()
}
