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

type ServiceAPIResponse struct {
	Resource      *string   `json:"resource,omitempty"`
	Size          *int64    `json:"size,omitempty"`
	PageNumber    *int64    `json:"pageNumber,omitempty"`
	TotalElements *int64    `json:"totalElements,omitempty"`
	TotalPages    *int64    `json:"totalPages,omitempty"`
	Services      []Service `json:"services,omitempty"`
}

type Service struct {
	Id                     *string              `json:"id,omitempty"`
	OwnerUserId            *string              `json:"ownerUserId,omitempty"`
	TenantId               *string              `json:"tenantId,omitempty" validate:"nonzero"`
	ParentService          *bool                `json:"parentService,omitempty"`
	ParentServiceId        *string              `json:"parentServiceId,omitempty"`
	Resource               *string              `json:"resource,omitempty"`
	Perms                  *[]string            `json:"perms,omitempty"`
	Name                   *string              `json:"name,omitempty" validate:"nonzero"`
	DisplayName            *string              `json:"displayName,omitempty" validate:"nonzero"`
	LogoPath               *string              `json:"logoPath,omitempty"`
	Description            *string              `json:"description,omitempty"`
	DefaultImageId         *int64               `json:"defaultImageId,omitempty"`
	ServiceType            *string              `json:"serviceType,omitempty"`
	SystemService          *bool                `json:"systemService,omitempty"`
	ExternalService        *bool                `json:"externalService,omitempty"`
	Visible                *bool                `json:"visible,omitempty"`
	ExternalBundleLocation *string              `json:"externalBundleLocation,omitempty"`
	BundleLocation         *string              `json:"bundleLocation,omitempty"`
	CostPerHour            *float64             `json:"costPerHour,omitempty"`
	OwnerId                *string              `json:"ownerId,omitempty"`
	ServiceActions         *[]ServiceAction     `json:"serviceActions,omitempty"`
	ServicePorts           *[]ServicePort       `json:"servicePorts,omitempty"`
	ServiceParamSpecs      *[]ServiceParamSpec  `json:"serviceParamSpecs,omitempty"`
	EgressRestrictions     *[]EgressRestriction `json:"egressRestrictions,omitempty"`
	Images                 *[]Image             `json:"images,omitempty" validate:"nonzero"`
	Repositories           *[]Repository        `json:"repositories,omitempty"`
	ChildServices          *[]Service           `json:"childServices,omitempty"`
	ExternalActions        *[]ExternalAction    `json:"externalActions,omitempty"`
}

type ServiceAction struct {
	ActionName  *string `json:"actionName,omitempty"`
	ActionType  *string `json:"actionType,omitempty"`
	ActionValue *string `json:"actionValue,omitempty"`
}

type ServicePort struct {
	Protocol *string `json:"protocol,omitempty"`
	FromPort *string `json:"fromPort,omitempty"`
	ToPort   *string `json:"toPort,omitempty"`
	CloudId  *string `json:"cloudId,omitempty"`
}

type ServiceParamSpec struct {
	ParamName            *string                `json:"paramName,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	HelpText             *string                `json:"helpText,omitempty"`
	Type                 *string                `json:"type,omitempty"`
	ValueList            *string                `json:"valueList,omitempty"`
	WebserviceListParams *[]WebserviceListParam `json:"webserviceListParams,omitempty"`
	DefaultValue         *string                `json:"defaultValue,omitempty"`
	UserVisible          *bool                  `json:"userVisible,omitempty"`
	UserEditable         *bool                  `json:"userEditable,omitempty"`
	SystemParam          *bool                  `json:"systemParam,omitempty"`
	ExampleValue         *string                `json:"exampleValue,omitempty"`
	Optional             *bool                  `json:"optional,omitempty"`
	ValueConstraint      *ValueConstraint       `json:"valueConstraint,omitempty"`
}

type EgressRestriction struct {
	EgressServiceName *string `json:"egressServiceName,omitempty"`
}

type Repository struct {
	Id          *string   `json:"id,omitempty"`
	Resource    *string   `json:"resource,omitempty"`
	Perms       *[]string `json:"perms,omitempty"`
	Hostname    *string   `json:"hostname,omitempty"`
	DisplayName *string   `json:"displayName,omitempty"`
	Protocol    *string   `json:"protocol,omitempty"`
	Description *string   `json:"description,omitempty"`
	Port        *int64    `json:"port,omitempty"`
}

func (s *Client) GetServices(tenantId int) ([]Service, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ServiceAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	service := data.Services
	return service, nil
}

func (s *Client) GetService(tenantId int, serviceId int) (*Service, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services/" + strconv.Itoa(serviceId))

	var data Service

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

	service := &data

	return service, nil
}

func (s *Client) AddService(service *Service) (*Service, error) {

	var data Service

	if errs := validator.Validate(service); errs != nil {
		return nil, errs
	}

	serviceTenantId := *service.TenantId

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + serviceTenantId + "/services")

	j, err := json.Marshal(service)

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

	service = &data

	return service, nil

}

func (s *Client) UpdateService(service *Service) (*Service, error) {

	var data Service

	if errs := validator.Validate(service); errs != nil {
		return nil, errs
	}

	serviceTenantId := *service.TenantId

	if nonzero(service.Id) {
		return nil, errors.New("Service.Id is missing")
	}

	serviceId := *service.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + serviceTenantId + "/services/" + serviceId)

	j, err := json.Marshal(service)

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

	service = &data

	return service, nil
}

func (s *Client) DeleteService(tenantId int, serviceId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services/" + strconv.Itoa(serviceId))

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
