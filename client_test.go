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
	err := client.RemoveCallback("67a090fa63ed2a505db27a8a")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientPauseCallback(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	err := client.PauseCallback("67a090fa63ed2a505db27a8a")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientResumeCallback(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	err := client.ResumeCallback("67a090fa63ed2a505db27a8a")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientUpdateCallback(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	callback, err := client.UpdateCallback(gosf.CallbackUpdateBody{
		ID: "67a090fa63ed2a505db27a8a",
		Addresses: []string{
			"JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4",
			"675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8",
		},
		CallbackUrl: "http://52.199.6.58:8443" + gosf.CB_SERVER_TX_CALLBACK_PATH,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callback)
}

func TestClientAddCallbackAddresses(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	callback, err := client.AddCallbackAddresses("67a090fa63ed2a505db27a8a", []string{
		"675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callback)
}

func TestClientRemoveCallbackAddresses(t *testing.T) {
	client := gosf.NewClient(os.Getenv("SHYFT_API_KEY"), nil, nil)
	callback, err := client.RemoveCallbackAddresses("67a090fa63ed2a505db27a8a", []string{
		"JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callback)
}
