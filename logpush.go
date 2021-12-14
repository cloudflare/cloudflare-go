package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// LogpushJob describes a Logpush job.
type LogpushJob struct {
	ID                 int        `json:"id,omitempty"`
	Dataset            string     `json:"dataset"`
	Enabled            bool       `json:"enabled"`
	Name               string     `json:"name"`
	LogpullOptions     string     `json:"logpull_options"`
	DestinationConf    string     `json:"destination_conf"`
	OwnershipChallenge string     `json:"ownership_challenge,omitempty"`
	LastComplete       *time.Time `json:"last_complete,omitempty"`
	LastError          *time.Time `json:"last_error,omitempty"`
	ErrorMessage       string     `json:"error_message,omitempty"`
}

// LogpushJobsResponse is the API response, containing an array of Logpush Jobs.
type LogpushJobsResponse struct {
	Response
	Result []LogpushJob `json:"result"`
}

// LogpushJobDetailsResponse is the API response, containing a single Logpush Job.
type LogpushJobDetailsResponse struct {
	Response
	Result LogpushJob `json:"result"`
}

// LogpushFieldsResponse is the API response for a datasets fields
type LogpushFieldsResponse struct {
	Response
	Result LogpushFields `json:"result"`
}

// LogpushFields is a map of available Logpush field names & descriptions
type LogpushFields map[string]string

// LogpushGetOwnershipChallenge describes a ownership validation.
type LogpushGetOwnershipChallenge struct {
	Filename string `json:"filename"`
	Valid    bool   `json:"valid"`
	Message  string `json:"message"`
}

// LogpushGetOwnershipChallengeResponse is the API response, containing a ownership challenge.
type LogpushGetOwnershipChallengeResponse struct {
	Response
	Result LogpushGetOwnershipChallenge `json:"result"`
}

// LogpushGetOwnershipChallengeRequest is the API request for get ownership challenge.
type LogpushGetOwnershipChallengeRequest struct {
	DestinationConf string `json:"destination_conf"`
}

// LogpushOwnershipChallengeValidationResponse is the API response,
// containing a ownership challenge validation result.
type LogpushOwnershipChallengeValidationResponse struct {
	Response
	Result struct {
		Valid bool `json:"valid"`
	}
}

// LogpushValidateOwnershipChallengeRequest is the API request for validate ownership challenge.
type LogpushValidateOwnershipChallengeRequest struct {
	DestinationConf    string `json:"destination_conf"`
	OwnershipChallenge string `json:"ownership_challenge"`
}

// LogpushDestinationExistsResponse is the API response,
// containing a destination exists check result.
type LogpushDestinationExistsResponse struct {
	Response
	Result struct {
		Exists bool `json:"exists"`
	}
}

// LogpushDestinationExistsRequest is the API request for check destination exists.
type LogpushDestinationExistsRequest struct {
	DestinationConf string `json:"destination_conf"`
}

// CreateLogpushJob creates a new LogpushJob for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-create-logpush-job
func (api *API) createLogpushJob(ctx context.Context, identifierType RouteRoot, identifier string, job LogpushJob) (*LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, job)
	if err != nil {
		return nil, err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return &r.Result, nil
}
func (api *API) CreateAccountLogpushJob(ctx context.Context, accountID string, job LogpushJob) (*LogpushJob, error) {
	return api.createLogpushJob(ctx, AccountRouteRoot, accountID, job)
}
func (api *API) CreateZoneLogpushJob(ctx context.Context, zoneID string, job LogpushJob) (*LogpushJob, error) {
	return api.createLogpushJob(ctx, ZoneRouteRoot, zoneID, job)
}
// Eventually deprecate this
func (api *API) CreateLogpushJob(ctx context.Context, zoneID string, job LogpushJob) (*LogpushJob, error) {
	return api.createLogpushJob(ctx, ZoneRouteRoot, zoneID, job)
}

// LogpushJobs returns all Logpush Jobs for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs
func (api *API) logpushJobs(ctx context.Context, identifierType RouteRoot, identifier string) ([]LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []LogpushJob{}, err
	}
	var r LogpushJobsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) ListAccountLogpushJobs(ctx context.Context, accountID string) ([]LogpushJob, error) {
	return api.logpushJobs(ctx, AccountRouteRoot, accountID, job)
}
func (api *API) ListZoneLogpushJobs(ctx context.Context, zoneID string) ([]LogpushJob, error) {
	return api.logpushJobs(ctx, ZoneRouteRoot, zoneID, job)
}
// Eventually deprecate this
func (api *API) LogpushJobs(ctx context.Context, zoneID string) ([]LogpushJob, error) {
	return api.logpushJobs(ctx, ZoneRouteRoot, zoneID, job)
}

// LogpushJobsForDataset returns all Logpush Jobs for a dataset in a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs-for-a-dataset
func (api *API) logpushJobsForDataset(ctx context.Context, identifierType RouteRoot, identifier, dataset string) ([]LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/datasets/%s/jobs", identifierType, identifier, dataset)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []LogpushJob{}, err
	}
	var r LogpushJobsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) ListAccountLogpushJobsForDataset(ctx context.Context, accountID, dataset string) ([]LogpushJob, error) {
	return api.logpushJobsForDataset(ctx, AccountRouteRoot, accountID, dataset)
}
func (api *API) ListZoneLogpushJobsForDataset(ctx context.Context, zoneID, dataset string) ([]LogpushJob, error) {
	return api.logpushJobsForDataset(ctx, ZoneRouteRoot, zoneID, dataset)
}
// Eventually deprecate this
func (api *API) LogpushJobsForDataset(ctx context.Context, zoneID, dataset string) ([]LogpushJob, error) {
	return api.logpushJobsForDataset(ctx, ZoneRouteRoot, zoneID, dataset)
}

// LogpushFields returns fields for a given dataset.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs
func (api *API) logpushFields(ctx context.Context, identifierType RouteRoot, identifier, dataset string) (LogpushFields, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/datasets/%s/fields", identifierType, identifier, dataset)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LogpushFields{}, err
	}
	var r LogpushFieldsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return LogpushFields{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) GetAccountLogpushFields(ctx context.Context, accountID, dataset string) (LogpushFields, error) {
	return api.logpushFields(ctx, AccountRouteRoot, accountID, dataset)
}
func (api *API) GetZoneLogpushFields(ctx context.Context, zoneID, dataset string) (LogpushFields, error) {
	return api.logpushFields(ctx, ZoneRouteRoot, zoneID, dataset)
}
// Eventually deprecate this
func (api *API) LogpushFields(ctx context.Context, zoneID, dataset string) (LogpushFields, error) {
	return api.logpushFields(ctx, ZoneRouteRoot, zoneID, dataset)
}

// LogpushJob fetches detail about one Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-logpush-job-details
func (api *API) logpushJob(ctx context.Context, identifierType RouteRoot, identifier string, jobID int) (LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", identifierType, identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LogpushJob{}, err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) GetAccountLogpushJob(ctx context.Context, accountID string, jobID int) (LogpushJob, error) {
	return api.logpushJob(ctx, AccountRouteRoot, accountID, jobID)
}
func (api *API) GetZoneLogpushJob(ctx context.Context, zoneID string, jobID int) (LogpushJob, error) {
	return api.logpushJob(ctx, ZoneRouteRoot, zoneID, jobID)
}
// Eventually deprecate this
func (api *API) LogpushJob(ctx context.Context, zoneID string, jobID int) (LogpushJob, error) {
	return api.logpushJob(ctx, ZoneRouteRoot, zoneID, jobID)
}

// UpdateLogpushJob lets you update a Logpush Job.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-update-logpush-job
func (api *API) updateLogpushJob(ctx context.Context, identifierType RouteRoot, identifier string, jobID int, job LogpushJob) error {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", identifierType, identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, job)
	if err != nil {
		return err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
func (api *API) UpdateAccountLogpushJob(ctx context.Context, accountID string, jobID int, job LogpushJob) error {
	return api.updateLogpushJob(ctx, AccountRouteRoot, accountID, jobID)
}
func (api *API) UpdateZoneLogpushJob(ctx context.Context, zoneID string, jobID int, job LogpushJob) error {
	return api.updateLogpushJob(ctx, ZoneRouteRoot, zoneID, jobID, job)
}
// Eventually deprecate this
func (api *API) UpdateLogpushJob(ctx context.Context, zoneID string, jobID int, job LogpushJob) error {
	return api.updateLogpushJob(ctx, ZoneRouteRoot, zoneID, jobID, job)
}

// DeleteLogpushJob deletes a Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-delete-logpush-job
func (api *API) deleteLogpushJob(ctx context.Context, identifierType RouteRoot, identifier string, jobID int) error {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", identifierType, identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
func (api *API) DeleteAccountLogpushJob(ctx context.Context, accountID string, jobID int) error {
	return api.deleteLogpushJob(ctx, AccountRouteRoot, accountID, jobID)
}
func (api *API) DeleteZoneLogpushJob(ctx context.Context, zoneID string, jobID int) error {
	return api.deleteLogpushJob(ctx, ZoneRouteRoot, zoneID, jobID)
}
// Eventually deprecate this
func (api *API) DeleteLogpushJob(ctx context.Context, zoneID string, jobID int) error {
	return api.deleteLogpushJob(ctx, ZoneRouteRoot, zoneID, jobID)
}

// GetLogpushOwnershipChallenge returns ownership challenge.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-get-ownership-challenge
func (api *API) getLogpushOwnershipChallenge(ctx context.Context, identifierType RouteRoot, identifier, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/ownership", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushGetOwnershipChallengeRequest{
		DestinationConf: destinationConf,
	})
	if err != nil {
		return nil, err
	}
	var r LogpushGetOwnershipChallengeResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !r.Result.Valid {
		return nil, errors.New(r.Result.Message)
	}

	return &r.Result, nil
}
func (api *API) GetAccountLogpushOwnershipChallenge(ctx context.Context, accountID, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	return api.getLogpushOwnershipChallenge(ctx, AccountRouteRoot, accountID, destinationConf)
}
func (api *API) GetZoneLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	return api.getLogpushOwnershipChallenge(ctx, ZoneRouteRoot, zoneID, destinationConf)
}
// Eventually deprecate this
func (api *API) GetLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	return api.getLogpushOwnershipChallenge(ctx, ZoneRouteRoot, zoneID, destinationConf)
}

// ValidateLogpushOwnershipChallenge returns ownership challenge validation result.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-validate-ownership-challenge
func (api *API) validateLogpushOwnershipChallenge(ctx context.Context, identifierType RouteRoot, identifier, destinationConf, ownershipChallenge string) (bool, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/ownership/validate", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushValidateOwnershipChallengeRequest{
		DestinationConf:    destinationConf,
		OwnershipChallenge: ownershipChallenge,
	})
	if err != nil {
		return false, err
	}
	var r LogpushGetOwnershipChallengeResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return false, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result.Valid, nil
}
func (api *API) ValidateAccountLogpushOwnershipChallenge(ctx context.Context, accountID, destinationConf, ownershipChallenge string) (bool, error) {
	return api.validateLogpushOwnershipChallenge(ctx, AccountRouteRoot, accountID, destinationConf)
}
func (api *API) ValidateZoneLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf, ownershipChallenge string) (bool, error) {
	return api.validateLogpushOwnershipChallenge(ctx, ZoneRouteRoot, zoneID, destinationConf, ownershipChallenge)
}
// Eventually deprecate this
func (api *API) ValidateLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf, ownershipChallenge string) (bool, error) {
	return api.validateLogpushOwnershipChallenge(ctx, ZoneRouteRoot, zoneID, destinationConf, ownershipChallenge)
}

// CheckLogpushDestinationExists returns destination exists check result.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-check-destination-exists
func (api *API) checkLogpushDestinationExists(ctx context.Context, identifierType RouteRoot, identifier, destinationConf string) (bool, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/validate/destination/exists", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushDestinationExistsRequest{
		DestinationConf: destinationConf,
	})
	if err != nil {
		return false, err
	}
	var r LogpushDestinationExistsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return false, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result.Exists, nil
}
func (api *API) CheckAccountLogpushDestinationExists(ctx context.Context, accountID, destinationConf string) (bool, error) {
	return api.checkLogpushDestinationExists(ctx, AccountRouteRoot, accountID, destinationConf)
}
func (api *API) CheckZoneLogpushDestinationExists(ctx context.Context, zoneID, destinationConf string) (bool, error) {
	return api.checkLogpushDestinationExists(ctx, ZoneRouteRoot, zoneID, destinationConf)
}
// Eventually deprecate this
func (api *API) CheckLogpushDestinationExists(ctx context.Context, zoneID, destinationConf string) (bool, error) {
	return api.checkLogpushDestinationExists(ctx, ZoneRouteRoot, zoneID, destinationConf)
}
