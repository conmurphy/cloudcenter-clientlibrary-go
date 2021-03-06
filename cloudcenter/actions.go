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

//import "bytes"

type ActionAPIResponse struct {
	Resource      *string  `json:"resource,omitempty"`
	Size          *int64   `json:"size,omitempty"`
	PageNumber    *int64   `json:"pageNumber,omitempty"`
	TotalElements *int64   `json:"totalElements,omitempty"`
	TotalPages    *int64   `json:"totalPages"`
	ActionJaxbs   []Action `json:"actionJaxbs,omitempty"`
}

type Action struct {
	Id                     *string                  `json:"id,omitempty"`
	Resource               *string                  `json:"resource,omitempty"`
	Perms                  *[]string                `json:"perms,omitempty"`
	Name                   *string                  `json:"name,omitempty" validate:"nonzero"`
	Description            *string                  `json:"description,omitempty"`
	ActionType             *string                  `json:"actionType,omitempty" validate:"nonzero"`
	LastUpdatedTime        *string                  `json:"lastUpdatedTime,omitempty"`
	TimeOut                *float64                 `json:"timeOut,omitempty"`
	Enabled                *bool                    `json:"enabled,omitempty"`
	Encrypted              *bool                    `json:"encrypted,omitempty"`
	Deleted                *bool                    `json:"deleted,omitempty"`
	SystemDefined          *bool                    `json:"systemDefined,omitempty"`
	BulkOperationSupported *bool                    `json:"bulkOperationSupported,omitempty"`
	IsAvailableToUser      *bool                    `json:"isAvailableToUser,omitempty"`
	Owner                  *int64                   `json:"owner,omitempty"`
	ActionParameters       *[]ActionParameter       `json:"actionParameters,omitempty" validate:"nonzero"`
	ActionResourceMappings *[]ActionResourceMapping `json:"actionResourceMappings,omitempty" validate:"nonzero"`
	ActionCustomParamSpecs *[]ActionCustomParamSpec `json:"actionCustomParamSpecs,omitempty"`
}

type ActionParameter struct {
	ParamName   *string `json:"paramName,omitempty"`
	ParamValue  *string `json:"paramValue,omitempty"`
	CustomParam *bool   `json:"customParam,omitempty"`
	Required    *bool   `json:"required,omitempty"`
	Preference  *string `json:"preference,omitempty"`
}

type ActionResourceMapping struct {
	Type                  *string                 `json:"type,omitempty" `
	ActionResourceFilters *[]ActionResourceFilter `json:"actionResourceFilters,omitempty" `
}

type ActionResourceFilter struct {
	DeploymentResource *string     `json:"deploymentResource,omitempty"`
	VmResource         *VmResource `json:"vmResource,omitempty"`
	IsEditable         *bool       `json:"isEditable,omitempty"`
}

type VmResource struct {
	Type                  *string                 `json:"type,omitempty"`
	AppProfiles           *[]string               `json:"appProfiles,omitempty"`
	CloudRegions          *[]string               `json:"cloudRegions,omitempty"`
	CloudAccounts         *[]string               `json:"cloudAccounts,omitempty"`
	Services              *[]string               `json:"services,omitempty"`
	OsTypes               *[]string               `json:"osTypes,omitempty"`
	CloudFamilyNames      *[]string               `json:"cloudFamilyNames,omitempty"`
	NodeStates            *[]string               `json:"nodesStates,omitempty"`
	CloudResourceMappings *[]CloudResourceMapping `json:"cloudResourceMappings,omitempty"`
}

type CloudResourceMapping struct {
	CloudFamily *string   `json:"cloudFamily,omitempty"`
	NodeStates  *[]string `json:"nodeStates,omitempty"`
}

type ActionCustomParamSpec struct {
	ParamName            *string              `json:"paramName,omitempty"`
	DisplayName          *string              `json:"displayName,omitempty"`
	HelpText             *string              `json:"helpText,omitempty"`
	Type                 *string              `json:"type,omitempty"`
	ValueList            *string              `json:"valueList,omitempty"`
	DefaultValue         *string              `json:"defaultValue,omitempty"`
	ConfirmValue         *string              `json:"confirmValue,omitempty"`
	PathSuffixValue      *string              `json:"pathSuffixValue,omitempty"`
	UserVisible          *bool                `json:"userVisible,omitempty"`
	UserEditable         *bool                `json:"userEditable,omitempty"`
	SystemParam          *bool                `json:"systemParam,omitempty"`
	ExampleValue         *string              `json:"exampleValue,omitempty"`
	DataUnit             *string              `json:"dataUnit,omitempty"`
	Optional             *bool                `json:"optional,omitempty"`
	MultiselectSupported *bool                `json:"multiselectSupported,omitempty"`
	ValueConstraint      *ValueConstraint     `json:"valueConstraint,omitempty"`
	Scope                *string              `json:"scope,omitempty"`
	WebserviceListParams *WebserviceListParam `json:"webserviceListParams,omitempty"`
	Preference           *string              `json:"preference,omitempty"`
}

type WebserviceListParam struct {
	URL           *string `json:"url,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	Username      *string `json:"username,omitempty"`
	Password      *string `json:"password,omitempty"`
	RequestType   *string `json:"requestType,omitempty"`
	ContentType   *string `json:"contentType,omitempty"`
	CommandParams *string `json:"commandParams,omitempty"`
	RequestBody   *string `json:"requestBody,omitempty"`
	ResultString  *string `json:"resultString,omitempty"`
}

func (s *Client) GetActions() ([]Action, error) {

	var data ActionAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/actions")

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

	actions := data.ActionJaxbs
	return actions, nil
}

func (s *Client) GetAction(id int) (*Action, error) {

	var data Action

	url := fmt.Sprintf(s.BaseURL + "/v1/actions/" + strconv.Itoa(id))
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

	action := &data
	return action, nil
}

func (s *Client) AddAction(action *Action) (*Action, error) {

	var data Action

	if errs := validator.Validate(action); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/actions")

	j, err := json.Marshal(action)

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

	action = &data

	return action, nil
}

func (s *Client) UpdateAction(action *Action) (*Action, error) {

	var data Action

	if errs := validator.Validate(action); errs != nil {
		return nil, errs
	}

	if nonzero(action.Id) {
		return nil, errors.New("Action.Id is missing")
	}

	actionId := *action.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/actions/" + actionId)

	j, err := json.Marshal(action)

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

	action = &data

	return action, nil
}

func (s *Client) DeleteAction(actionId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/actions/" + strconv.Itoa(actionId))

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
