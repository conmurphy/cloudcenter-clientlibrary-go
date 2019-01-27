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

type ActionPolicyAPIResponse struct {
	Resource       *string        `json:"resource,omitempty"`
	Size           *int64         `json:"size,omitempty"`
	PageNumber     *int64         `json:"pageNumber,omitempty"`
	TotalElements  *int64         `json:"totalElements,omitempty"`
	TotalPages     *int64         `json:"totalPages"`
	ActionPolicies []ActionPolicy `json:"customPolicyJaxbs,omitempty"`
}

type ActionPolicy struct {
	Id          *string    `json:"id,omitempty"`
	Resource    *string    `json:"resource,omitempty"`
	Perms       *[]string  `json:"perms,omitempty"`
	Name        *string    `json:"name,omitempty" validate:"nonzero"`
	Description *string    `json:"description,omitempty"`
	EntityType  *string    `json:"entityType,omitempty" validate:"nonzero"`
	EventName   *string    `json:"eventName,omitempty" validate:"nonzero"`
	Actions     *[]Actions `json:"actions,omitempty" validate:"nonzero"`
	UserId      *string    `json:"userId,omitempty"`
	Enabled     *bool      `json:"enabled,omitempty"`
	AutoEnable  *bool      `json:"autoEnable,omitempty"`
	ForceEnable *bool      `json:"forceEnable,omitempty"`
	Global      *bool      `json:"global,omitempty"`
}

type Actions struct {
	ActionType       *string            `json:"actionType,omitempty"`
	ActionInputs     *[]ActionInput     `json:"actionInputs,omitempty"`
	AssociatedParams *[]AssociatedParam `json:"associatedParams,omitempty"`
}

type ActionInput struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type AssociatedParam struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (s *Client) GetActionPolicies() ([]ActionPolicy, error) {

	var data ActionPolicyAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/actionpolicies")

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

	actionPolicy := data.ActionPolicies
	return actionPolicy, nil
}

func (s *Client) GetActionPolicy(actionPolicyId int) (*ActionPolicy, error) {

	var data ActionPolicy

	url := fmt.Sprintf(s.BaseURL + "/v1/actionpolicies/" + strconv.Itoa(actionPolicyId))
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

	actionPolicy := &data
	return actionPolicy, nil
}

func (s *Client) AddActionPolicy(actionPolicy *ActionPolicy) (*ActionPolicy, error) {

	var data ActionPolicy

	if errs := validator.Validate(actionPolicy); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/actionpolicies")

	j, err := json.Marshal(actionPolicy)

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

	actionPolicy = &data

	return actionPolicy, nil
}

func (s *Client) UpdateActionPolicy(actionPolicy *ActionPolicy) (*ActionPolicy, error) {

	var data ActionPolicy

	if errs := validator.Validate(actionPolicy); errs != nil {
		return nil, errs
	}

	if nonzero(actionPolicy.Id) {
		return nil, errors.New("ActionPolicy.Id is missing")
	}

	actionPolicyId := *actionPolicy.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/actionpolicies/" + actionPolicyId)

	j, err := json.Marshal(actionPolicy)

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

	actionPolicy = &data

	return actionPolicy, nil
}

func (s *Client) DeleteActionPolicy(actionPolicyId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/actionpolicies/" + strconv.Itoa(actionPolicyId))

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
