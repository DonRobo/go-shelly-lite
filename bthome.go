package shelly

import (
	"resty.dev/v3"
)

// BTHomeAddDeviceRequest contains parameters for the BTHome.GetConfig RPC request.
type BTHomeAddDeviceRequest struct {
	// ID for the new component. Accepted range: [200..299]. Optional. If omitted, the first free
	// ID will be used. If the desired ID is not available, an error will be returned.
	ID int `json:"id"`

	// Config to be used for the new component.
	Config BTHomeDeviceConfig `json:"config"`
}

func (r *BTHomeAddDeviceRequest) Method() string {
	return "BTHome.AddDevice"
}

func (r *BTHomeAddDeviceRequest) NewTypedResponse() *BTHomeAddDeviceResponse {
	return &BTHomeAddDeviceResponse{}
}

func (r *BTHomeAddDeviceRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeAddDeviceRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeAddDeviceResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type BTHomeAddDeviceResponse struct {
	// Key of the newly created component. (in format <type>:<cid>, for example bthomedevice:200)
	Key string `json:"key"`
}

// BTHomeDeleteDeviceRequest contains parameters for the BTHome.DeleteDevice RPC request.
type BTHomeDeleteDeviceRequest struct {
	// ID of existing BTHomeDevice component (Required)
	ID int `json:"id"`
}

func (r *BTHomeDeleteDeviceRequest) Method() string {
	return "BTHome.DeleteDevice"
}

func (r *BTHomeDeleteDeviceRequest) NewTypedResponse() *BTHomeDeleteDeviceResponse {
	return &BTHomeDeleteDeviceResponse{}
}

func (r *BTHomeDeleteDeviceRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeleteDeviceRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeDeleteDeviceResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type BTHomeDeleteDeviceResponse struct{}

// BTHomeAddSensorRequest contains parameters for the BTHome.AddSensor RPC request.
type BTHomeAddSensorRequest struct {
	// ID for the new component. Accepted range: [200..299]. Optional. If omitted, the first free
	// ID will be used. If the desired ID is not available, an error will be returned.
	ID *int `json:"id"`

	// Config to be used for the new component.
	Config BTHomeSensorConfig `json:"config"`
}

func (r *BTHomeAddSensorRequest) Method() string {
	return "BTHome.AddSensor"
}

func (r *BTHomeAddSensorRequest) NewTypedResponse() *BTHomeAddSensorResponse {
	return &BTHomeAddSensorResponse{}
}

func (r *BTHomeAddSensorRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeAddSensorRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeAddSensorResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type BTHomeAddSensorResponse struct {
	// Key of the newly created component. (in format <type>:<cid>, for example bthomesensor:200)
	Key string `json:"key"`
}

// BTHomeDeleteSensorRequest contains parameters for the BTHome.DeleteSensor RPC request.
type BTHomeDeleteSensorRequest struct {
	// ID of existing BTHomeSensor component (Required)
	ID int `json:"id"`
}

func (r *BTHomeDeleteSensorRequest) Method() string {
	return "BTHome.DeleteSensor"
}

func (r *BTHomeDeleteSensorRequest) NewTypedResponse() *BTHomeDeleteSensorResponse {
	return &BTHomeDeleteSensorResponse{}
}

func (r *BTHomeDeleteSensorRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeleteSensorRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeDeleteSensorResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type BTHomeDeleteSensorResponse struct{}

// BTHomeStartDeviceDiscoveryRequest contains parameters for the BTHome.GetConfig RPC request.
type BTHomeStartDeviceDiscoveryRequest struct {
	// Duration of the discovery process in seconds. Default is 30 seconds.
	Duration int `json:"duration,omitempty"`
}

func (r *BTHomeStartDeviceDiscoveryRequest) Method() string {
	return "BTHome.StartDeviceDiscovery"
}

func (r *BTHomeStartDeviceDiscoveryRequest) NewTypedResponse() *BTHomeStartDeviceDiscoveryResponse {
	return &BTHomeStartDeviceDiscoveryResponse{}
}

func (r *BTHomeStartDeviceDiscoveryRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeStartDeviceDiscoveryRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeStartDeviceDiscoveryResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

type BTHomeStartDeviceDiscoveryResponse struct{}

// BTHomeGetObjectInfosRequest contains parameters for the BTHome.GetConfig RPC request.
type BTHomeGetObjectInfosRequest struct {
	// Offset index of the component from which to start generating the result (Optional)
	Offset *int `json:"offset,omitempty"`
}

func (r *BTHomeGetObjectInfosRequest) Method() string {
	return "BTHome.GetObjectInfos"
}

func (r *BTHomeGetObjectInfosRequest) NewTypedResponse() *BTHomeGetObjectInfosResponse {
	return &BTHomeGetObjectInfosResponse{}
}

func (r *BTHomeGetObjectInfosRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeGetObjectInfosRequest) DoResty(
	client *resty.Client,
) (
	*BTHomeGetObjectInfosResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

// BTHomeGetObjectInfosResponse ...
// As of 2024-12-17 this isn't officially documented.
type BTHomeGetObjectInfosResponse struct {
	Objects []BTHomeGetObjectInfo `json:"objects"`

	Offset int `json:"offset,omitempty"`

	Count int `json:"count,omitempty"`

	Total int `json:"total,omitempty"`
}

type BTHomeGetObjectInfo struct {
	ObjID int `json:"obj_id,omitempty"`

	ObjName string `json:"obj_name,omitempty"`

	Type string `json:"type,omitempty"`

	Unit string `json:"unit,omitempty"`
}
