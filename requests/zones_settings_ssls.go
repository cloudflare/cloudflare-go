package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type SslListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// SSL encrypts your visitor's connection and safeguards credit card numbers and
	// other personal data to and from your website. SSL can take up to 5 minutes to
	// fully activate. Requires Cloudflare active on your root domain or www domain.
	// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
	// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
	// your web server. You don't need to have an SSL cert on your web server, but your
	// vistors will still see the site as being HTTPS enabled. Full: SSL between the
	// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
	// Cloudflare and your web server. You'll need to have your own SSL cert or
	// self-signed cert at the very least. Full (Strict): SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
	// your web server. You'll need to have a valid SSL certificate installed on your
	// web server. This certificate must be signed by a certificate authority, have an
	// expiration date in the future, and respond for the request domain name
	// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
	Result fields.Field[Ssl] `json:"result"`
}

// MarshalJSON serializes SslListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SslListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslListResponse) String() (result string) {
	return fmt.Sprintf("&SslListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SslUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// SSL encrypts your visitor's connection and safeguards credit card numbers and
	// other personal data to and from your website. SSL can take up to 5 minutes to
	// fully activate. Requires Cloudflare active on your root domain or www domain.
	// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
	// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
	// your web server. You don't need to have an SSL cert on your web server, but your
	// vistors will still see the site as being HTTPS enabled. Full: SSL between the
	// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
	// Cloudflare and your web server. You'll need to have your own SSL cert or
	// self-signed cert at the very least. Full (Strict): SSL between the visitor and
	// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
	// your web server. You'll need to have a valid SSL certificate installed on your
	// web server. This certificate must be signed by a certificate authority, have an
	// expiration date in the future, and respond for the request domain name
	// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
	Result fields.Field[Ssl] `json:"result"`
}

// MarshalJSON serializes SslUpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SslUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslUpdateResponse) String() (result string) {
	return fmt.Sprintf("&SslUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SslUpdateParams struct {
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value fields.Field[SslValue] `json:"value,required"`
}

// MarshalJSON serializes SslUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SslUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslUpdateParams) String() (result string) {
	return fmt.Sprintf("&SslUpdateParams{Value:%s}", r.Value)
}
