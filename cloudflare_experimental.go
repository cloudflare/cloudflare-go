package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"errors"

	"golang.org/x/time/rate"
)

type service struct {
	client *Client
}

type ClientParams struct {
	Key            string
	Email          string
	UserServiceKey string
	Token          string
	BaseURL        *url.URL
	UserAgent      string
	Headers        http.Header
	HTTPClient     *http.Client
	RateLimiter    *rate.Limiter
	RetryPolicy    RetryPolicy
	Logger         Logger
	Debug          bool
}

// A Client manages communication with the Cloudflare API.
type Client struct {
	clientMu sync.Mutex

	*ClientParams

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Zones *ZonesService
}

// Client returns the http.Client used by this Cloudflare client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.HTTPClient
	return &clientCopy
}

// Call is the entrypoint to making API calls with the correct request setup.
func (c *Client) Call(ctx context.Context, method, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, method, path, payload, nil)
}

// CallWithHeaders is the entrypoint to making API calls with the correct
// request setup and allows passing in additional HTTP headers with the request.
func (c *Client) CallWithHeaders(ctx context.Context, method, path string, payload interface{}, headers http.Header) ([]byte, error) {
	return c.makeRequest(ctx, method, path, payload, headers)
}

// New creates a new instance of the API client by merging ClientParams with the
// default values.
func NewExperimental(config *ClientParams) (*Client, error) {
	c := &Client{ClientParams: &ClientParams{}}
	c.common.client = c

	defaultURL, _ := url.Parse(defaultScheme + "://" + defaultHostname + defaultBasePath)
	if config.BaseURL != nil {
		c.ClientParams.BaseURL = config.BaseURL
	} else {
		c.ClientParams.BaseURL = defaultURL
	}

	if config.UserAgent != "" {
		c.ClientParams.UserAgent = config.UserAgent
	} else {
		c.ClientParams.UserAgent = userAgent + "/" + Version
	}

	if config.HTTPClient != nil {
		c.ClientParams.HTTPClient = config.HTTPClient
	} else {
		c.ClientParams.HTTPClient = http.DefaultClient
	}

	if config.RateLimiter != nil {
		c.ClientParams.RateLimiter = config.RateLimiter
	} else {
		c.ClientParams.RateLimiter = rate.NewLimiter(rate.Limit(4), 1) // 4rps equates to default api limit (1200 req/5 min)
	}

	retryPolicy := RetryPolicy{
		MaxRetries:    3,
		MinRetryDelay: time.Duration(1) * time.Second,
		MaxRetryDelay: time.Duration(30) * time.Second,
	}
	c.ClientParams.RetryPolicy = retryPolicy

	if config.Headers != nil {
		c.ClientParams.Headers = config.Headers
	} else {
		c.ClientParams.Headers = make(http.Header)
	}

	if config.Logger != nil {
		c.ClientParams.Logger = config.Logger
	} else {
		c.ClientParams.Logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	if config.Key != "" && config.Token != "" {
		return nil, errors.New("API key and tokens are mutually exclusive")
	}

	if config.Key != "" {
		c.ClientParams.Key = config.Key
		c.ClientParams.Email = config.Email
	}

	if config.Token != "" {
		c.ClientParams.Token = config.Token
	}

	if config.UserServiceKey != "" {
		c.ClientParams.UserServiceKey = config.UserServiceKey
	}

	c.ClientParams.Debug = config.Debug

	c.Zones = (*ZonesService)(&c.common)

	return c, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (c *Client) request(ctx context.Context, method, uri string, reqBody io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL.String()+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, c.Headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if c.Key == "" && c.Email == "" && c.Token == "" && c.UserServiceKey == "" {
		return nil, errors.New("no user credentials provided")
	}

	if c.Key != "" {
		req.Header.Set("X-Auth-Key", c.ClientParams.Key)
		req.Header.Set("X-Auth-Email", c.ClientParams.Email)
	}

	if c.UserServiceKey != "" {
		req.Header.Set("X-Auth-User-Service-Key", c.ClientParams.UserServiceKey)
	}

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.ClientParams.Token)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.ClientParams.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
}

func (c *Client) makeRequest(ctx context.Context, method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	var reqBody io.Reader
	var err error

	if params != nil {
		if r, ok := params.(io.Reader); ok {
			reqBody = r
		} else if paramBytes, ok := params.([]byte); ok {
			reqBody = bytes.NewReader(paramBytes)
		} else {
			var jsonBody []byte
			jsonBody, err = json.Marshal(params)
			if err != nil {
				return nil, fmt.Errorf("error marshalling params to JSON: %w", err)
			}
			reqBody = bytes.NewReader(jsonBody)
		}
	}

	var resp *http.Response
	var respErr error
	var respBody []byte
	for i := 0; i <= c.RetryPolicy.MaxRetries; i++ {
		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(c.RetryPolicy.MinRetryDelay))

			if sleepDuration > c.RetryPolicy.MaxRetryDelay {
				sleepDuration = c.RetryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			c.Logger.Printf("sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)

			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
			}
		}

		err = c.RateLimiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("error caused by request rate limiting: %w", err)
		}

		if c.Debug {
			if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
				buf := &bytes.Buffer{}
				tee := io.TeeReader(reqBody, buf)
				debugBody, _ := ioutil.ReadAll(tee)
				payloadBody, _ := ioutil.ReadAll(buf)
				fmt.Printf("cloudflare-go [DEBUG] REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, c.BaseURL.String()+uri, headers, string(debugBody))
				// ensure we recreate the io.Reader for use
				reqBody = bytes.NewReader(payloadBody)
			} else {
				fmt.Printf("cloudflare-go [DEBUG] REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, c.BaseURL.String()+uri, headers, nil)
			}
		}

		resp, respErr = c.request(ctx, method, uri, reqBody, headers)

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			// if we got a valid http response, try to read body so we can reuse the connection
			// see https://golang.org/pkg/net/http/#Client.Do
			if respErr == nil {
				respBody, err = ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = fmt.Errorf("could not read response body: %w", err)
			} else {
				c.Logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("could not read response body: %w", err)
			}
			break
		}
	}

	if c.Debug {
		fmt.Printf("cloudflare-go [DEBUG] RESPONSE URI:%s StatusCode:%d Body:%#v RayID:%s\n", c.BaseURL.String()+uri, resp.StatusCode, string(respBody), resp.Header.Get("cf-ray"))
	}

	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr") {
			return nil, fmt.Errorf("%s", respBody)
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, &ServiceError{cloudflareError: &Error{
				StatusCode: resp.StatusCode,
				RayID:      resp.Header.Get("cf-ray"),
				Errors: []ResponseInfo{{
					Message: errInternalServiceError,
				}},
			}}
		}

		errBody := &Response{}
		err = json.Unmarshal(respBody, &errBody)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
		}

		errCodes := make([]int, 0, len(errBody.Errors))
		errMsgs := make([]string, 0, len(errBody.Errors))
		for _, e := range errBody.Errors {
			errCodes = append(errCodes, e.Code)
			errMsgs = append(errMsgs, e.Message)
		}

		err := &Error{
			StatusCode:    resp.StatusCode,
			RayID:         resp.Header.Get("cf-ray"),
			Errors:        errBody.Errors,
			ErrorCodes:    errCodes,
			ErrorMessages: errMsgs,
		}

		switch resp.StatusCode {
		case http.StatusUnauthorized:
			err.Type = ErrorTypeAuthorization
			return nil, &AuthorizationError{cloudflareError: err}
		case http.StatusForbidden:
			err.Type = ErrorTypeAuthentication
			return nil, &AuthenticationError{cloudflareError: err}
		case http.StatusNotFound:
			err.Type = ErrorTypeNotFound
			return nil, &NotFoundError{cloudflareError: err}
		case http.StatusTooManyRequests:
			err.Type = ErrorTypeRateLimit
			return nil, &RatelimitError{cloudflareError: err}
		default:
			err.Type = ErrorTypeRequest
			return nil, &RequestError{cloudflareError: err}
		}
	}

	return respBody, nil
}
