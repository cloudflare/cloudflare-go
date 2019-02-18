package cloudflare

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// LogpushJob describes a Logpush job
type LogpushJob struct {
	ID              int        `json:"id,omitempty"`
	Enabled         bool       `json:"enabled"`
	Name            string     `json:"name"`
	LogpullOptions  string     `json:"logpull_options"`
	DestinationConf string     `json:"destination_conf"`
	LastComplete    *time.Time `json:"last_complete,omitempty"`
	LastError       *time.Time `json:"last_error,omitempty"`
	ErrorMessage    string     `json:"error_message,omitempty"`
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

// CreateLogpushJob creates a new LogpushJob for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-create-logpush-job
func (api *API) CreateLogpushJob(zoneID string, job LogpushJob) (*LogpushJob, error) {
	uri := "/zones/" + zoneID + "/logpush/jobs"
	res, err := api.makeRequest("POST", uri, job)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return &r.Result, nil
}

// LogpushJobs returns all Logpush Jobs for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs
func (api *API) LogpushJobs(zoneID string) ([]LogpushJob, error) {
	uri := "/zones/" + zoneID + "/logpush/jobs"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errMakeRequestError)
	}
	var r LogpushJobsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LogpushJob fetches detail about one Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-logpush-job-details
func (api *API) LogpushJob(zoneID string, jobID int) (LogpushJob, error) {
	uri := "/zones/" + zoneID + "/logpush/jobs/" + strconv.Itoa(jobID)
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LogpushJob{}, errors.Wrap(err, errMakeRequestError)
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateLogpushJob lets you update a Logpush Job.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-update-logpush-job
func (api *API) UpdateLogpushJob(zoneID string, jobID int, job LogpushJob) error {
	uri := "/zones/" + zoneID + "/logpush/jobs/" + strconv.Itoa(jobID)
	res, err := api.makeRequest("PUT", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}

// DeleteLogpushJob deletes a Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-delete-logpush-job
func (api *API) DeleteLogpushJob(zoneID string, jobID int) error {
	uri := "/zones/" + zoneID + "/logpush/jobs/" + strconv.Itoa(jobID)
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
