package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type PseudoIpv4ListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// The value set for the Pseudo IPv4 setting.
	Result fields.Field[PseudoIpv4] `json:"result"`
}

// MarshalJSON serializes PseudoIpv4ListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PseudoIpv4ListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PseudoIpv4ListResponse) String() (result string) {
	return fmt.Sprintf("&PseudoIpv4ListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PseudoIpv4UpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// The value set for the Pseudo IPv4 setting.
	Result fields.Field[PseudoIpv4] `json:"result"`
}

// MarshalJSON serializes PseudoIpv4UpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PseudoIpv4UpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PseudoIpv4UpdateResponse) String() (result string) {
	return fmt.Sprintf("&PseudoIpv4UpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PseudoIpv4UpdateParams struct {
	// Value of the Pseudo IPv4 setting.
	Value fields.Field[PseudoIpv4Value] `json:"value,required"`
}

// MarshalJSON serializes PseudoIpv4UpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PseudoIpv4UpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PseudoIpv4UpdateParams) String() (result string) {
	return fmt.Sprintf("&PseudoIpv4UpdateParams{Value:%s}", r.Value)
}
