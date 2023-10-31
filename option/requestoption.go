// File generated from our OpenAPI spec by Stainless.

package option

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/tidwall/sjson"
)

// RequestOption is an option for the requests made by the cloudflare API Client
// which can be supplied to clients, services, and methods. You can read more about this functional
// options pattern in our [README].
//
// [README]: https://pkg.go.dev/SDK_PackagePath#readme-requestoptions
type RequestOption = func(*requestconfig.RequestConfig) error

// WithBaseURL returns a RequestOption that sets the BaseURL for the client.
func WithBaseURL(base string) RequestOption {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatalf("failed to parse BaseURL: %s\n", err)
	}
	return func(r *requestconfig.RequestConfig) error {
		r.BaseURL = u
		return nil
	}
}

// WithHTTPClient returns a RequestOption that changes the underlying [http.Client] used to make this
// request, which by default is [http.DefaultClient].
func WithHTTPClient(client *http.Client) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.HTTPClient = client
		return nil
	}
}

// MiddlewareNext is a function which is called by a middleware to pass an HTTP request
// to the next stage in the middleware chain.
type MiddlewareNext = func(*http.Request) (*http.Response, error)

// Middleware is a function which intercepts HTTP requests, processing or modifying
// them, and then passing the request to the next middleware or handler
// in the chain by calling the provided MiddlewareNext function.
type Middleware = func(*http.Request, MiddlewareNext) (*http.Response, error)

// WithMiddleware returns a RequestOption that applies the given middleware
// to the requests made. Each middleware will execute in the order they were given.
func WithMiddleware(middlewares ...Middleware) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Middlewares = append(r.Middlewares, middlewares...)
		return nil
	}
}

// WithMaxRetries returns a RequestOption that sets the maximum number of retries that the client
// attempts to make. When given 0, the client only makes one request. By
// default, the client retries two times.
//
// WithMaxRetries panics when retries is negative.
func WithMaxRetries(retries int) RequestOption {
	if retries < 0 {
		panic("option: cannot have fewer than 0 retries")
	}
	return func(r *requestconfig.RequestConfig) error {
		r.MaxRetries = retries
		return nil
	}
}

// WithHeader returns a RequestOption that sets the header value to the associated key. It overwrites
// any value if there was one already present.
func WithHeader(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = []string{value}
		return nil
	}
}

// WithHeaderAdd returns a RequestOption that adds the header value to the associated key. It appends
// onto any existing values.
func WithHeaderAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = append(r.Request.Header[key], value)
		return nil
	}
}

// WithHeaderDel returns a RequestOption that deletes the header value(s) associated with the given key.
func WithHeaderDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		delete(r.Request.Header, key)
		return nil
	}
}

// WithQuery returns a RequestOption that sets the query value to the associated key. It overwrites
// any value if there was one already present.
func WithQuery(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Set(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithQueryAdd returns a RequestOption that adds the query value to the associated key. It appends
// onto any existing values.
func WithQueryAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Add(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithQueryDel returns a RequestOption that deletes the query value(s) associated with the key.
func WithQueryDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Del(key)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithJSONSet returns a RequestOption that sets the body's JSON value associated with the key.
// The key accepts a string as defined by the [sjson format].
//
// [sjson format]: https://github.com/tidwall/sjson
func WithJSONSet(key string, value interface{}) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.SetBytes(r.Buffer, key, value)
		return err
	}
}

// WithJSONDel returns a RequestOption that deletes the body's JSON value associated with the key.
// The key accepts a string as defined by the [sjson format].
//
// [sjson format]: https://github.com/tidwall/sjson
func WithJSONDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.DeleteBytes(r.Buffer, key)
		return err
	}
}

// WithResponseBodyInto returns a RequestOption that overwrites the deserialization target with
// the given destination. If provided, we don't deserialize into the default struct.
func WithResponseBodyInto(dst any) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseBodyInto = dst
		return nil
	}
}

// WithResponseInto returns a RequestOption that copies the [*http.Response] into the given address.
func WithResponseInto(dst **http.Response) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseInto = dst
		return nil
	}
}

// WithRequestTimeout returns a RequestOption that sets the timeout for
// each request attempt. This should be smaller than the timeout defined in
// the context, which spans all retries.
func WithRequestTimeout(dur time.Duration) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.RequestTimeout = dur
		return nil
	}
}

// WithEnvironmentProduction returns a RequestOption that sets the current
// environment to be the "production" environment. An environment specifies which base URL
// to use by default.
func WithEnvironmentProduction() RequestOption {
	return WithBaseURL("https://api.cloudflare.com/client/v4/")
}

// WithAPIKey returns a RequestOption that sets the client setting "api_key".
func WithAPIKey(value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.APIKey = value
		return r.Apply(WithHeader("authorization", fmt.Sprintf("Bearer %s", r.APIKey)))
	}
}
