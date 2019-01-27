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

type EnvironmentAPIResponse struct {
	Resource      *string       `json:"resource,omitempty"`
	Size          *int64        `json:"size,omitempty"`
	PageNumber    *int64        `json:"pageNumber,omitempty"`
	TotalElements *int64        `json:"totalElements,omitempty"`
	TotalPages    *int64        `json:"totalPages,omitempty"`
	Environments  []Environment `json:"deploymentEnvironments,omitempty"`
}

type Environment struct {
	Id                 *string            `json:"id,omitempty"`
	Resource           *string            `json:"resource,omitempty"`
	Name               *string            `json:"name,omitempty" validate:"nonzero"`
	Perms              *[]string          `json:"perms,omitempty"`
	Description        *string            `json:"description,omitempty"`
	AllowedClouds      *string            `json:"allowedClouds,omitempty"`
	DefaultSettings    *string            `json:"defaultSettings,omitempty"`
	RequiresApproval   *bool              `json:"requiresApproval,omitempty"`
	AssociatedClouds   *[]AssociatedCloud `json:"associatedClouds,omitempty"`
	TotalDeployments   *int64             `json:"totalDeployments,omitempty"`
	ExtensionId        *string            `json:"extensionId,omitempty"`
	CostDetails        *CostDetail        `json:"costDetails,omitempty"`
	NetworkTypes       *[]NetworkType     `json:"networkTypes,omitempty"`
	NetworkTypeEnabled *bool              `json:"networkTypeEnabled,omitempty"`
	RestrictedUser     *bool              `json:"restrictedUser,omitempty"`
	DefaultRegionId    *string            `json:"defaultRegionId,omitempty"`
	Owner              *int64             `json:"owner,omitempty"`
}

type CostDetail struct {
	TotalCloudCost *float64 `json:"totalCloudCost,omitempty"`
	TotalAppCost   *float64 `json:"totalAppCost,omitempty"`
	TotalJobsCost  *float64 `json:"totalJobsCost,omitempty"`
}

type NetworkType struct {
	Name                  *string  `json:"name,omitempty"`
	Description           *string  `json:"description,omitempty"`
	NumberOfNetworkMapped *float64 `json:"numberOfNetworkMapped,omitempty"`
}

type AssociatedCloud struct {
	RegionId                 *string                    `json:"regionId,omitempty"`
	RegionName               *string                    `json:"regionName,omitempty"`
	RegionDisplayName        *string                    `json:"regionDisplayName,omitempty"`
	CloudFamily              *string                    `json:"cloudFamily,omitempty"`
	CloudId                  *string                    `json:"cloudId,omitempty"`
	CloudAccountId           *string                    `json:"cloudAccountId,omitempty"`
	CloudName                *string                    `json:"cloudName,omitempty"`
	CloudAccountName         *string                    `json:"cloudAccountName,omitempty"`
	CloudAssociationDefaults *[]CloudAssociationDefault `json:"cloudAssociationDefaults,omitempty"`
	Default                  *bool                      `json:"default,omitempty"`
}

type CloudAssociationDefault struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

func (s *Client) GetEnvironments() ([]Environment, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/environments")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data EnvironmentAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	environment := data.Environments
	return environment, nil
}

func (s *Client) GetEnvironment(id int) (*Environment, error) {

	var data Environment

	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + strconv.Itoa(id))
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

	environment := &data
	return environment, nil
}

func (s *Client) AddEnvironment(environment *Environment) (*Environment, error) {

	var data Environment

	if errs := validator.Validate(environment); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/environments")

	j, err := json.Marshal(environment)

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

	environment = &data

	return environment, nil
}

func (s *Client) UpdateEnvironment(environment *Environment) (*Environment, error) {

	var data Environment

	if errs := validator.Validate(environment); errs != nil {
		return nil, errs
	}

	if nonzero(environment.Id) {
		return nil, errors.New("Environment.Id is missing")
	}

	environmentId := *environment.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + environmentId)

	j, err := json.Marshal(environment)

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

	environment = &data

	return environment, nil
}

func (s *Client) DeleteEnvironment(environmentId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + strconv.Itoa(environmentId))

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
