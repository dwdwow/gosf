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
