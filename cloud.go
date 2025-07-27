package shelly

import "resty.dev/v3"

type CloudSetConfigRequest struct {
	Config CloudConfig `json:"config"`
}

func (r *CloudSetConfigRequest) Method() string {
	return "Cloud.SetConfig"
}

func (r *CloudSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *CloudSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudSetConfigRequest) DoResty(
	client *resty.Client,
) (
	*SetConfigResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type CloudConfig struct {
	// Enable is true if cloud connection is enabled, false otherwise
	Enable bool `json:"enable"`

	// Server is the name of the server to which the device is connected (optional).
	Server *string `json:"server"`
}

type CloudGetConfigRequest struct{}

func (r *CloudGetConfigRequest) Method() string {
	return "Cloud.GetConfig"
}

func (r *CloudGetConfigRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *CloudGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudGetConfigRequest) DoResty(
	client *resty.Client,
) (
	*RPCEmptyResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type CloudStatus struct {
	Connected bool `json:"connected"`
}

type CloudGetStatusRequest struct{}

func (r *CloudGetStatusRequest) Method() string {
	return "Cloud.GetStatus"
}

func (r *CloudGetStatusRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *CloudGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudGetStatusRequest) DoResty(
	client *resty.Client,
) (
	*RPCEmptyResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}
