package shelly_test

import (
	"testing"

	"resty.dev/v3"
)

func TestRestyRpc(t *testing.T) {
	// Test POST request to Shelly RPC API using resty v3
	client := resty.New()
	defer client.Close()

	url := "http://192.168.1.169/rpc"
	body := map[string]interface{}{
		"id":     0,
		"method": "Shelly.GetStatus",
		"params": map[string]interface{}{},
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	t.Logf("Status: %s", res.Status())
	t.Logf("Body: %s", res.String())
}
