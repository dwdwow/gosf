package gosf_test

import (
	"fmt"
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
	err = sv.HTTPServe()
	if err != nil {
		t.Fatal(err)
	}
	for {
		select {
		case acct := <-sv.GetAcctChannel():
			fmt.Printf("%+v\n", acct)
		case tx := <-sv.GetTxChannel():
			fmt.Printf("%+v\n", tx)
		}
	}
}
