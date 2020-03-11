// Copyright 2020 The go-spdx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spdx

// LicenseList represets a SPDX license list.
type LicenseList struct {
	LicenseListVersion string    `json:"licenseListVersion"`
	Licenses           []License `json:"licenses"`
	ReleaseDate        string    `json:"releaseDate"`
}

// License represets a SPDX license.
type License struct {
	DetailsURL            string   `json:"detailsUrl"`
	IsDeprecatedLicenseID bool     `json:"isDeprecatedLicenseId"`
	IsFsfLibre            bool     `json:"isFsfLibre"`
	IsOsiApproved         bool     `json:"isOsiApproved"`
	LicenseID             string   `json:"licenseId"`
	Name                  string   `json:"name"`
	Reference             string   `json:"reference"`
	ReferenceNumber       string   `json:"referenceNumber"`
	SeeAlso               []string `json:"seeAlso"`

	Details Details `json:"-"`
}

// Details represets a SPDX license details.
type Details struct {
	IsDeprecatedLicenseID   bool     `json:"isDeprecatedLicenseId"`
	IsFsfLibre              bool     `json:"isFsfLibre"`
	IsOsiApproved           bool     `json:"isOsiApproved"`
	LicenseID               string   `json:"licenseId"`
	LicenseText             string   `json:"licenseText"`
	Name                    string   `json:"name"`
	SeeAlso                 []string `json:"seeAlso"`
	StandardLicenseTemplate string   `json:"standardLicenseTemplate"`
}
