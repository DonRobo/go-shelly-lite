package shelly

import (
	"resty.dev/v3"
)

type ScriptConfig struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Name of the script instance.
	Name *string `json:"name,omitempty"`

	// Enable should true if the script runs by default on boot, false otherwise
	Enable *bool `json:"enable,omitempty"`
}

type ScriptStatus struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Running is true if the script is currently running, false otherwise
	Running bool `json:"running"`

	// Errors present only when the script execution resulted in an error. The array
	// contains description of the type of error. Possible errors are: crashed, syntax_error,
	// reference_error, type_error, out_of_memory, out_of_codespace, internal_error,
	// not_implemented, file_read_error, bad_arguments.
	Errors []string `json:"errors,omitempty"`
}

type ScriptGetConfigRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`
}

func (r *ScriptGetConfigRequest) Method() string {
	return "Script.GetConfig"
}

func (r *ScriptGetConfigRequest) DoResty(
	client *resty.Client,
) (
	*ScriptConfig,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptGetConfigRequest) NewTypedResponse() *ScriptConfig {
	return &ScriptConfig{}
}

func (r *ScriptGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptSetConfigRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config ScriptConfig `json:"config"`
}

func (r *ScriptSetConfigRequest) Method() string {
	return "Script.SetConfig"
}

func (r *ScriptSetConfigRequest) DoResty(
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

func (r *ScriptSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *ScriptSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptGetStatusRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`
}

func (r *ScriptGetStatusRequest) Method() string {
	return "Script.GetStatus"
}

func (r *ScriptGetStatusRequest) DoResty(
	client *resty.Client,
) (
	*ScriptStatus,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptGetStatusRequest) NewTypedResponse() *ScriptStatus {
	return &ScriptStatus{}
}

func (r *ScriptGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptCreateResponse struct {
	// ID of the created script component instance.
	ID int `json:"id"`
}

type ScriptCreateRequest struct {
	Name *string `json:"name,omitempty"`
}

func (r *ScriptCreateRequest) Method() string {
	return "Script.Create"
}

func (r *ScriptCreateRequest) DoResty(
	client *resty.Client,
) (
	*ScriptCreateResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptCreateRequest) NewTypedResponse() *ScriptCreateResponse {
	return &ScriptCreateResponse{}
}

func (r *ScriptCreateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptPutCodeResponse struct {
	Len int `json:"len"`
}

type ScriptPutCodeRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Code which will be included in the script (the length must be greater than 0).
	Code string `json:"code"`

	// Append is true if more data will be appended afterwards, default false.
	Append bool `json:"append,omitempty"`
}

func (r *ScriptPutCodeRequest) Method() string {
	return "Script.PutCode"
}

func (r *ScriptPutCodeRequest) DoResty(
	client *resty.Client,
) (
	*ScriptPutCodeResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptPutCodeRequest) NewTypedResponse() *ScriptPutCodeResponse {
	return &ScriptPutCodeResponse{}
}

func (r *ScriptPutCodeRequest) NewResponse() any {
	return r.NewTypedResponse()
}

//TODO Not supported yet
// ScriptPutCode is a helper method which uploads the provided code to the
// Script.PutCode method, line-by-line to accomodate limits on payload size.
// func ScriptPutCode(
// 	ctx context.Context,
// 	c mgrpc.MgRPC,
// 	credsCallback mgrpc.GetCredsCallback,
// 	data io.Reader,
// ) error {
// 	s := bufio.NewScanner(data)
// 	req := &ScriptPutCodeRequest{}
// 	for s.Scan() {
// 		req.Code = s.Text() + "\n"
// 		if _, _, err := req.Do(ctx, c, credsCallback); err != nil {
// 			return err
// 		}
// 		req.Append = true
// 	}
// 	if err := s.Err(); err != nil {
// 		return fmt.Errorf("reading input data for Script.PutCode: %w", err)
// 	}
// 	return nil
// }

type ScriptEvalResponse struct {
	Result string `json:"result"`
}

type ScriptEvalRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Argument to evaluate (the length must be greater than 0). Required
	// Note: the examples don't show this field?
	Code string `json:"code"`
}

func (r *ScriptEvalRequest) Method() string {
	return "Script.Eval"
}

func (r *ScriptEvalRequest) DoResty(
	client *resty.Client,
) (
	*ScriptEvalResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptEvalRequest) NewTypedResponse() *ScriptEvalResponse {
	return &ScriptEvalResponse{}
}

func (r *ScriptEvalRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptStartResponse struct {
	WasRunning bool `json:"was_running"`
}

type ScriptStartRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`
}

func (r *ScriptStartRequest) Method() string {
	return "Script.Start"
}

func (r *ScriptStartRequest) DoResty(
	client *resty.Client,
) (
	*ScriptStartResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptStartRequest) NewTypedResponse() *ScriptStartResponse {
	return &ScriptStartResponse{}
}

func (r *ScriptStartRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptStopResponse struct {
	WasRunning bool `json:"was_running"`
}

type ScriptStopRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`
}

func (r *ScriptStopRequest) Method() string {
	return "Script.Stop"
}

func (r *ScriptStopRequest) DoResty(
	client *resty.Client,
) (
	*ScriptStopResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptStopRequest) NewTypedResponse() *ScriptStopResponse {
	return &ScriptStopResponse{}
}

func (r *ScriptStopRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptListScript struct {
	// ID of the script component instance.
	ID int `json:"id"`

	// Name of the script
	Name string `json:"name"`

	// Enable is true if the script runs by default on boot, false otherwise.
	Enable bool `json:"enable"`

	// Running is true if currently running, false otherwise
	Running bool `json:"running"`
}

type ScriptListResponse struct {
	// Scripts is a list of all created scripts
	Scripts []ScriptListScript
}

type ScriptListRequest struct{}

func (r *ScriptListRequest) Method() string {
	return "Script.List"
}

func (r *ScriptListRequest) DoResty(
	client *resty.Client,
) (
	*ScriptListResponse,
	*Frame,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := DoResty(client, r, resp)
	return resp, raw, err
}

func (r *ScriptListRequest) NewTypedResponse() *ScriptListResponse {
	return &ScriptListResponse{}
}

func (r *ScriptListRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ScriptDeleteRequest struct {
	// ID of the script component instance.
	ID int `json:"id"`
}

func (r *ScriptDeleteRequest) Method() string {
	return "Script.Delete"
}

func (r *ScriptDeleteRequest) DoResty(
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

func (r *ScriptDeleteRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ScriptDeleteRequest) NewResponse() any {
	return r.NewTypedResponse()
}
