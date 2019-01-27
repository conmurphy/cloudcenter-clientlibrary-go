/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

               https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.
*/

package cloudcenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	validator "gopkg.in/validator.v2"
)

type BundleAPIResponse struct {
	Resource      *string  `json:"resource"`
	Size          *int64   `json:"size"`
	PageNumber    *int64   `json:"pageNumber"`
	TotalElements *int64   `json:"totalElements"`
	TotalPages    *int64   `json:"totalPages"`
	Bundles       []Bundle `json:"bundles"`
}

type Bundle struct {
	Id               *string   `json:"id,omitempty"`
	Resource         *string   `json:"resource,omitempty"`
	Perms            *[]string `json:"perms,omitempty"`
	Type             *string   `json:"type,omitempty" validate:"nonzero"`
	Name             *string   `json:"name,omitempty" validate:"nonzero"`
	Description      *string   `json:"description,omitempty"`
	Limit            *float64  `json:"limit,omitempty"`
	Price            *float64  `json:"price,omitempty"`
	ExpirationDate   *float64  `json:"expirationDate,omitempty" validate:"nonzero"`
	ExpirationMonths *int64    `json:"expirationMonths,omitempty"`
	Disabled         *bool     `json:"disabled,omitempty"`
	ShowOnlyToAdmin  *bool     `json:"showOnlyToAdmin,omitempty"`
	NumberOfUsers    *float64  `json:"numberOfUsers,omitempty"`
	TenantId         *string   `json:"tenantId,omitempty" validate:"nonzero"`
	PublishedAppIds  *[]string `json:"publishedAppIds,omitempty"`
}

func (s *Client) GetBundles(TenantId int) ([]Bundle, error) {

	var data BundleAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	bundles := data.Bundles
	return bundles, nil
}

func (s *Client) GetBundle(TenantId int, BundleId int) (*Bundle, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles/" + strconv.Itoa(BundleId))

	var data Bundle

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	bundle := &data

	return bundle, nil
}

func (s *Client) GetBundleFromName(TenantId int, BundleNameSearchString string) (*Bundle, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data BundleAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	bundles := data.Bundles

	for _, bundle := range bundles {
		bundleName := *bundle.Name
		if bundleName == BundleNameSearchString {

			return &bundle, nil
		}
	}

	return nil, errors.New("BUNDLE NOT FOUND")

}

func (s *Client) AddBundle(bundle *Bundle) (*Bundle, error) {

	var data Bundle

	if errs := validator.Validate(bundle); errs != nil {
		return nil, errs
	}

	bundleTenantId := *bundle.TenantId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundleTenantId + "/bundles")

	j, err := json.Marshal(bundle)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	bundle = &data

	return bundle, nil
}

func (s *Client) UpdateBundle(bundle *Bundle) (*Bundle, error) {

	var data Bundle

	if errs := validator.Validate(bundle); errs != nil {
		return nil, errs
	}

	if nonzero(bundle.Id) {
		return nil, errors.New("Bundle.Id is missing")
	}

	bundleTenantId := *bundle.TenantId
	bundleId := *bundle.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundleTenantId + "/bundles/" + bundleId)

	j, err := json.Marshal(bundle)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	bundle = &data

	return bundle, nil
}

func (s *Client) DeleteBundle(tenantId int, bundleId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/bundles/" + strconv.Itoa(bundleId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
