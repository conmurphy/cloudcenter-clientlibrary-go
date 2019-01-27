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

type ActivationProfileAPIResponse struct {
	Resource           *string             `json:"resource"`
	Size               *int64              `json:"size"`
	PageNumber         *int64              `json:"pageNumber"`
	TotalElements      *int64              `json:"totalElements"`
	TotalPages         *int64              `json:"totalPages"`
	ActivationProfiles []ActivationProfile `json:"activationProfiles"`
}

type ActivationProfile struct {
	Id                  *string           `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty" validate:"nonzero"`
	Description         *string           `json:"description,omitempty"`
	Resource            *string           `json:"resource"`
	TenantId            *int64            `json:"tenantId,omitempty" validate:"nonzero"`
	PlanId              *string           `json:"planId,omitempty"`
	BundleId            *string           `json:"bundleId,omitempty"`
	ContractId          *string           `json:"contractId,omitempty"`
	DepEnvId            *string           `json:"depEnvId,omitempty"`
	ActivateRegions     *[]ActivateRegion `json:"activateRegions,omitempty"`
	ImportApps          *[]string         `json:"importApps,omitempty"`
	AgreeToContract     *bool             `json:"agreeToContract,omitempty"`
	SendActivationEmail *bool             `json:"sendActivationEmail,omitempty"`
}

type ActivateRegion struct {
	RegionId string `json:"regionId"`
}

func (s *Client) GetActivationProfiles(tenantId int) ([]ActivationProfile, error) {

	var data ActivationProfileAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles")

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

	activationProfile := data.ActivationProfiles
	return activationProfile, nil
}

func (s *Client) GetActivationProfile(tenantId int, activationProfileId int) (*ActivationProfile, error) {

	var data ActivationProfile

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles/" + strconv.Itoa(activationProfileId))
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

	activationProfile := &data
	return activationProfile, nil
}

func (s *Client) AddActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error) {

	var data ActivationProfile

	if errs := validator.Validate(activationProfile); errs != nil {
		return nil, errs
	}

	activationProfileTenantId := int(*activationProfile.TenantId)

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(activationProfileTenantId) + "/activationProfiles")

	j, err := json.Marshal(activationProfile)

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

	activationProfile = &data

	return activationProfile, nil
}

func (s *Client) UpdateActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error) {

	var data ActivationProfile

	if errs := validator.Validate(activationProfile); errs != nil {
		return nil, errs
	}

	if nonzero(activationProfile.Id) {
		return nil, errors.New("ActivationProfile.Id is missing")
	}

	activationProfileTenantId := int(*activationProfile.TenantId)
	activationProfileId := *activationProfile.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(activationProfileTenantId) + "/activationProfiles/" + activationProfileId)

	j, err := json.Marshal(activationProfile)

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

	activationProfile = &data

	return activationProfile, nil
}

func (s *Client) DeleteActivationProfile(tenantId int, activationProfileId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles/" + strconv.Itoa(activationProfileId))

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
