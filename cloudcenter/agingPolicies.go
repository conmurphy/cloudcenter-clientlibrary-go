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

type AgingPolicyAPIResponse struct {
	Resource      *string       `json:"resource,omitempty"`
	Size          *int64        `json:"size,omitempty"`
	PageNumber    *int64        `json:"pageNumber,omitempty"`
	TotalElements *int64        `json:"totalElements,omitempty"`
	TotalPages    *int64        `json:"totalPages"`
	AgingPolicies []AgingPolicy `json:"policies,omitempty"`
}

type AgingPolicy struct {
	Id                             *string                `json:"id,omitempty"`
	Resource                       *string                `json:"resource,omitempty"`
	Perms                          *[]string              `json:"perms,omitempty"`
	Name                           *string                `json:"name,omitempty" validate:"nonzero"`
	Description                    *string                `json:"description,omitempty"`
	Enabled                        *bool                  `json:"enabled,omitempty" validate:"nonzero"`
	Type                           *string                `json:"type,omitempty" validate:"nonzero"`
	Limit                          *Limit                 `json:"limit,omitempty" validate:"nonzero"`
	TerminateWhenPolicyEnds        *bool                  `json:"terminateWhenPolicyEnds,omitempty"`
	AllowGracePeriodForTermination *bool                  `json:"allowGracePeriodForTermination,omitempty"`
	GraceLimit                     *GraceLimit            `json:"graceLimit,omitempty"`
	AllowPolicyExtension           *bool                  `json:"allowPolicyExtension,omitempty"`
	ExtensionLimit                 *ExtensionLimit        `json:"extensionLimit,omitempty"`
	AllowGracePeriodNotification   *bool                  `json:"allowGracePeriodNotification,omitempty"`
	AllowPolicyExpiryNotification  *bool                  `json:"allowPolicyExpiryNotification,omitempty"`
	Notifications                  *[]Notification        `json:"notifications,omitempty"`
	IsPolicyActiveOnResources      *bool                  `json:"isPolicyActiveOnResources,omitempty"`
	Created                        *float64               `json:"created,omitempty"`
	LastUpdated                    *float64               `json:"lastUpdated,omitempty"`
	Resources                      *[]AgingPolicyResource `json:"resources,omitempty"`
	Priority                       *float64               `json:"priority,omitempty"`
	OwnerId                        *int64                 `json:"ownerId,omitempty"`
}

type Limit struct {
	Amount *float64 `json:"amount,omitempty"`
	Unit   *string  `json:"unit,omitempty"`
}

type GraceLimit struct {
	Amount *float64 `json:"amount,omitempty"`
	Unit   *string  `json:"unit,omitempty"`
}

type ExtensionLimit struct {
	NumOfExtensions      *int64                `json:"numOfExtensions,omitempty"`
	LimitOfEachExtension *LimitOfEachExtension `json:"limitOfEachExtension,omitempty"`
}

type LimitOfEachExtension struct {
	Amount *float64 `json:"amount,omitempty"`
	Unit   *string  `json:"unit,omitempty"`
}

type Notification struct {
	Template  *string     `json:"template,omitempty"`
	Type      *string     `json:"type,omitempty"`
	Enabled   *bool       `json:"enabled,omitempty"`
	Reminders *[]Reminder `json:"reminders,omitempty"`
}

type Reminder struct {
	Amount *float64 `json:"amount,omitempty"`
	Unit   *string  `json:"unit,omitempty"`
}

type AgingPolicyResource struct {
	ResourceId                  *string  `json:"resourceId,omitempty"`
	ResourceType                *string  `json:"resourceType,omitempty"`
	AppliedDate                 *float64 `json:"appliedDate,omitempty"`
	ResourceStartTime           *float64 `json:"resourceStartTime,omitempty"`
	EstimatedPolicyEndTime      *float64 `json:"estimatedPolicyEndTime,omitempty"`
	AllowedCost                 *float64 `json:"allowedCost,omitempty"`
	AccruedCost                 *float64 `json:"accruedCost,omitempty"`
	NumberOfExtensionsUsed      *int64   `json:"numberOfExtensionsUsed,omitempty"`
	IsApprovalPending           *bool    `json:"isApprovalPending,omitempty"`
	IsPreviousExtensionDenied   *bool    `json:"isPreviousExtensionDenied,omitempty"`
	IsPolicyReachingExpiry      *bool    `json:"isPolicyReachingExpiry,omitempty"`
	IsPolicyReachingGraceExpiry *bool    `json:"isPolicyReachingGraceExpiry,omitempty"`
}

func (s *Client) GetAgingPolicies() ([]AgingPolicy, error) {

	var data AgingPolicyAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v2/agingPolicies")

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

	agingPolicy := data.AgingPolicies
	return agingPolicy, nil
}

func (s *Client) GetAgingPolicy(agingPolicyId int) (*AgingPolicy, error) {

	var data AgingPolicy

	url := fmt.Sprintf(s.BaseURL + "/v2/agingPolicies/" + strconv.Itoa(agingPolicyId))
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

	agingPolicy := &data
	return agingPolicy, nil
}

func (s *Client) AddAgingPolicy(agingPolicy *AgingPolicy) (*AgingPolicy, error) {

	var data AgingPolicy

	if errs := validator.Validate(agingPolicy); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v2/agingPolicies")

	j, err := json.Marshal(agingPolicy)

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

	agingPolicy = &data

	return agingPolicy, nil
}

func (s *Client) UpdateAgingPolicy(agingPolicy *AgingPolicy) (*AgingPolicy, error) {

	var data AgingPolicy

	if errs := validator.Validate(agingPolicy); errs != nil {
		return nil, errs
	}

	if nonzero(agingPolicy.Id) {
		return nil, errors.New("AgingPolicy.Id is missing")
	}

	agingPolicyId := *agingPolicy.Id
	url := fmt.Sprintf(s.BaseURL + "/v2/agingPolicies/" + agingPolicyId)

	j, err := json.Marshal(agingPolicy)

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

	agingPolicy = &data

	return agingPolicy, nil
}

func (s *Client) DeleteAgingPolicy(agingPolicyId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v2/agingPolicies/" + strconv.Itoa(agingPolicyId))

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
