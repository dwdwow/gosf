package gosf_test

import (
	"os"
	"testing"

	"github.com/dwdwow/gosf"
)

func TestClientCallbackList(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	callbacks, err := client.CallbackList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callbacks)
}

func TestClientCreateCallback(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	callback, err := client.CreateCallback(gosf.CallbackCreateBody{
		Network: gosf.NetworkMainnetBeta,
		Addresses: []string{
			"JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4",
			// "675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8",
		},
		CallbackUrl: "http://52.199.6.58:8443" + gosf.CB_SERVER_TX_CALLBACK_PATH,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callback)
}

func TestClientRemoveCallback(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	err := client.RemoveCallback("679f6e515c51ef30ab7d63d0")
	if err != nil {
		t.Fatal(err)
	}
}
