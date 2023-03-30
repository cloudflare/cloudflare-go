package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type Tls_1_3ListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result fields.Field[Tls_1_3] `json:"result"`
}

// MarshalJSON serializes Tls_1_3ListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *Tls_1_3ListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Tls_1_3ListResponse) String() (result string) {
	return fmt.Sprintf("&Tls_1_3ListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Tls_1_3UpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result fields.Field[Tls_1_3] `json:"result"`
}

// MarshalJSON serializes Tls_1_3UpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *Tls_1_3UpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Tls_1_3UpdateResponse) String() (result string) {
	return fmt.Sprintf("&Tls_1_3UpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Tls_1_3UpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[Tls_1_3Value] `json:"value,required"`
}

// MarshalJSON serializes Tls_1_3UpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *Tls_1_3UpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Tls_1_3UpdateParams) String() (result string) {
	return fmt.Sprintf("&Tls_1_3UpdateParams{Value:%s}", r.Value)
}
