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
	sv := gosf.NewCallbackServer("8443", homeDir+"/certs/certificate.crt", homeDir+"/certs/private.key", nil)
	sv.Start()
	for {
		select {
		case acct := <-sv.GetAcctChannel():
			t.Logf("%+v", acct)
		case tx := <-sv.GetTxChannel():
			t.Logf("%+v", tx)
		}
	}
}
