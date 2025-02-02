package gosf_test

import (
	"os"
	"testing"

	"github.com/dwdwow/gosf"
)

func TestCallbackServer(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	sv := gosf.NewCallbackServer("8443", homeDir+"/certs/certificate.pem", homeDir+"/certs/private.key", nil)
	sv.Start()
}
