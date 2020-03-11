// Copyright 2020 The go-spdx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spdx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DefaultListURL    = "https://spdx.org/licenses/licenses.json"
	DefaultDetailsURL = "https://spdx.org/licenses/%[1]s.json"
)

// SPDX is an API client for accessing SPDX data.
//
// Configure any fields on the struct prior to calling any functions. After
// calling functions, do not access the fields again.
//
// The functions on SPDX are safe to call concurrently.
type SPDX struct {
	// hc is the hc client to use for requests. If this is nil, then
	// a default new hc client will be used.
	hc *http.Client

	// listURL and DetailsURL are the URLs for listing licenses and accessing
	// a single license, respectively. If these are not set, they will default
	// to the default values specified in constants.
	listURL string

	// detailsURL use the placeholder "%[1]s" to interpolate the SPDX ID.
	detailsURL string
}

// Options represents a set options to SPDX.
type Options func(s *SPDX)

// WithClient inject custom http.Client to SPDX.
func WithClient(hc *http.Client) Options {
	return func(s *SPDX) { s.hc = hc }
}

// WithListURL inject custom listURL to SPDX.
func WithListURL(listURL string) Options {
	return func(s *SPDX) { s.listURL = listURL }
}

// WithDetailsURL inject custom detailsURL to SPDX.
func WithDetailsURL(detailsURL string) Options {
	return func(s *SPDX) { s.detailsURL = detailsURL }
}

// New returns the new SPDX client.
func New(opts ...Options) *SPDX {
	s := &SPDX{
		hc:         http.DefaultClient,
		listURL:    DefaultListURL,
		detailsURL: DefaultDetailsURL,
	}

	for _, o := range opts {
		o(s)
	}

	return s
}

// List returns the list of licenses.
func (s *SPDX) List() (*LicenseList, error) {
	resp, err := s.hc.Get(s.listURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result LicenseList
	return &result, json.NewDecoder(resp.Body).Decode(&result)
}

// License returns the SPDX license data by ID.
//
// This often includes more detailed information than List such as the full license text.
//
// The ID is usually case sensitive. Please ensure the ID is set exactly
// to the SPDX ID, including casing.
func (s *SPDX) License(id string) (*License, error) {
	resp, err := s.hc.Get(fmt.Sprintf(s.detailsURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result License
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	respDetails, err := s.hc.Get(result.DetailsURL)
	if err != nil {
		return nil, err
	}
	defer respDetails.Body.Close()

	var details Details
	if err := json.NewDecoder(respDetails.Body).Decode(&details); err != nil {
		return nil, err
	}
	result.Details = details

	return &result, nil
}
