package gosf_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dwdwow/gosf"
)

func TestCallbackServer(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	sv := gosf.NewCallbackServer("8443", homeDir+"/certs/certificate.crt", homeDir+"/certs/private.key", nil)
	err = sv.StartHTTPServer()
	if err != nil {
		t.Fatal(err)
	}
	for {
		select {
		case acct := <-sv.GetAcctChannel():
			fmt.Printf("%s %+v\n", time.Now().Format(time.RFC3339), acct)
		case tx := <-sv.GetTxChannel():
			fmt.Printf("%s %+v\n", time.Now().Format(time.RFC3339), tx)
		}
	}
}
