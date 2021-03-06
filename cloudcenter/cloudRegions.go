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

type CloudRegionAPIResponse struct {
	Resource      *string       `json:"resource,omitempty"`
	Size          *int64        `json:"size,omitempty"`
	PageNumber    *int64        `json:"pageNumber,omitempty"`
	TotalElements *int64        `json:"totalElements,omitempty"`
	TotalPages    *int64        `json:"totalPages,omitempty"`
	CloudRegions  []CloudRegion `json:"cloudRegions,omitempty"`
}

type CloudRegion struct {
	Id                     *string           `json:"id,omitempty"`
	TenantId               *string           `json:"tenantId,omitempty" validate:"nonzero"`
	CloudId                *string           `json:"cloudId,omitempty" validate:"nonzero"`
	CloudRegionId          *string           `json:"cloudRegionId,omitempty"`
	ImportRegion           *ImportRegion     `json:"importRegion,omitempty"`
	Resource               *string           `json:"resource,omitempty"`
	Perms                  *[]string         `json:"perms,omitempty"`
	DisplayName            *string           `json:"displayName,omitempty" validate:"nonzero`
	RegionName             *string           `json:"regionName,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Gateway                *Gateway          `json:"gateway,omitempty"`
	Storage                *Storage          `json:"storage,omitempty"`
	Enabled                *bool             `json:"enabled,omitempty"`
	Activated              *bool             `json:"activated,omitempty"`
	PublicCloud            *bool             `json:"publicCloud,omitempty"`
	NumUsers               *int32            `json:"numUsers,omitempty"`
	Status                 *string           `json:"status,omitempty"`
	StatusDetail           *string           `json:"statusDetail,omitempty"`
	RegionProperties       *[]RegionProperty `json:"regionProperties,omitempty"`
	ExternalBundleLocation *string           `json:"externalBundleLocation,omitempty"`
	ExternalActions        *[]ExternalAction `json:"externalActions,omitempty"`
}

type ImportRegion struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

type Gateway struct {
	Address        *string `json:"address,omitempty"`
	DNSName        *string `json:"dnsName,omitempty"`
	Status         *string `json:"status,omitempty"`
	CloudId        *string `json:"cloudId,omitempty"`
	CloudAccountId *string `json:"cloudAccountId,omitempty"`
}

type Storage struct {
	RegionId              *string                 `json:"regionId,omitempty"`
	CloudAccountId        *string                 `json:"cloudAccountId,omitempty"`
	Size                  *int64                  `json:"size,omitempty"`
	NumNodes              *int64                  `json:"numNodes,omitempty"`
	CloudSpecificSettings *[]CloudSpecificSetting `json:"cloudSpecificSettings,omitempty"`
	Address               *string                 `json:"address,omitempty"`
}

type CloudSpecificSetting struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type RegionProperty struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ExternalAction struct {
	ActionName  *string `json:"actionName,omitempty"`
	ActionType  *string `json:"actionType,omitempty"`
	ActionValue *string `json:"actionValue,omitempty"`
}

func (s *Client) GetCloudRegions(tenantId int, cloudId int) ([]CloudRegion, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudRegionAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudRegion := data.CloudRegions
	return cloudRegion, nil
}

func (s *Client) GetCloudRegion(tenantId int, cloudId int, regionId int) (*CloudRegion, error) {

	var data CloudRegion

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId))
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

	cloudRegion := &data
	return cloudRegion, nil
}

func (s *Client) AddCloudRegion(cloudRegion *CloudRegion) (*CloudRegion, error) {

	var data CloudRegion

	if errs := validator.Validate(cloudRegion); errs != nil {
		return nil, errs
	}

	if nonzero(cloudRegion.RegionName) {
		return nil, errors.New("CloudRegion.RegionName is missing")
	}

	cloudRegionTenantId := *cloudRegion.TenantId
	cloudRegionCloudId := *cloudRegion.CloudId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudRegionTenantId + "/clouds/" + cloudRegionCloudId + "/regions")

	j, err := json.Marshal(cloudRegion)

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

	cloudRegion = &data

	return cloudRegion, nil
}

func (s *Client) UpdateCloudRegion(cloudRegion *CloudRegion) (*CloudRegion, error) {

	var data CloudRegion

	if errs := validator.Validate(cloudRegion); errs != nil {
		return nil, errs
	}

	if nonzero(cloudRegion.Id) {
		return nil, errors.New("CloudRegionRole.Id is missing")
	}

	cloudRegionTenantId := *cloudRegion.TenantId
	cloudRegionCloudId := *cloudRegion.CloudId
	cloudRegionId := *cloudRegion.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudRegionTenantId + "/clouds/" + cloudRegionCloudId + "/regions/" + cloudRegionId)

	j, err := json.Marshal(cloudRegion)

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

	cloudRegion = &data

	return cloudRegion, nil
}

func (s *Client) DeleteCloudRegion(tenantId int, cloudId int, cloudRegionId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(cloudRegionId))

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
