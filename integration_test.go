package shelly_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"

	shelly "github.com/DonRobo/go-shelly-lite"
)

func GetCallWithVerify(t *testing.T, req shelly.RPCRequestBody, respBody interface{}) {
	client := resty.New()
	client.SetBaseURL("http://192.168.1.169")
	defer client.Close()

	respFrame, err := shelly.DoResty(client, req, respBody)
	require.NoError(t, err)
	fmt.Println(string(respFrame.Result))

	// The reencoded JSON *SHOULD* match.
	// NOTE: in practice there seem to be some undocumented fields and inconsistency in what
	// is NULL and what is omited when NULL.
	jsonOut, err := json.Marshal(respBody)
	require.NoError(t, err)
	assert.JSONEq(t, string(respFrame.Result), string(jsonOut))
}
