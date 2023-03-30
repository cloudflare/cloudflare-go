package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type CipherListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result fields.Field[Ciphers] `json:"result"`
}

// MarshalJSON serializes CipherListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CipherListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CipherListResponse) String() (result string) {
	return fmt.Sprintf("&CipherListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type CipherUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result fields.Field[Ciphers] `json:"result"`
}

// MarshalJSON serializes CipherUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CipherUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CipherUpdateResponse) String() (result string) {
	return fmt.Sprintf("&CipherUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type CipherUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[[]string] `json:"value,required"`
}

// MarshalJSON serializes CipherUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CipherUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CipherUpdateParams) String() (result string) {
	return fmt.Sprintf("&CipherUpdateParams{Value:%s}", core.Fmt(r.Value))
}
