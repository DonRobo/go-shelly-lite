package shelly

import "resty.dev/v3"

type EthStatus struct {
	// IP of the device in the network.
	IP *string `json:"ip"`
}

type EthGetStatusRequest struct{}

func (r *EthGetStatusRequest) Method() string {
	return "Eth.GetStatus"
}

func (r *EthGetStatusRequest) NewTypedResponse() *EthStatus {
	return &EthStatus{}
}

func (r *EthGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *EthGetStatusRequest) Do(
	client *resty.Client,
) (
	*EthStatus,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(client, r, resp)
	return resp, raw, err
}

type EthConfig struct {
	// Enable is true if the configuration is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`

	// IPv4Mode Range of values: dhcp, static
	IPv4Mode *string `json:"ipv4mode,omitempty"`

	// IP to use when ipv4mode is static.
	IP *string `json:"ip,omitempty"`

	// Netmask to use when ipv4mode is static
	Netmask *string `json:"netmask,omitempty"`

	// GW is the gateway to use when ipv4mode is static
	GW *string `json:"gw,omitempty"`

	// Nameserver to use when ipv4mode is static
	Nameserver *string `json:"nameserver,omitempty"`
}

type EthGetConfigRequest struct{}

func (r *EthGetConfigRequest) Method() string {
	return "Eth.GetConfig"
}

func (r *EthGetConfigRequest) NewTypedResponse() *EthConfig {
	return &EthConfig{}
}

func (r *EthGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *EthGetConfigRequest) Do(
	client *resty.Client,
) (
	*EthConfig,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(client, r, resp)
	return resp, raw, err
}

type EthSetConfigRequest struct {
	Config EthConfig `json:"config"`
}

func (r *EthSetConfigRequest) Method() string {
	return "Eth.SetConfig"
}

func (r *EthSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *EthSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *EthSetConfigRequest) Do(
	client *resty.Client,
) (
	*SetConfigResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(client, r, resp)
	return resp, raw, err
}
