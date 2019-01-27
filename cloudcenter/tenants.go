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

type TenantAPIResponse struct {
	Resource      *string  `json:"resource"`
	Size          *int     `json:"size"`
	PageNumber    *int     `json:"pageNumber"`
	TotalElements *int     `json:"totalElements"`
	TotalPages    *int     `json:"totalPages"`
	Tenants       []Tenant `json:"tenants"`
}

type Tenant struct {
	Id                              *string       `json:"id,omitempty"`
	Resource                        *string       `json:"resource,omitempty"`
	Name                            *string       `json:"name,omitempty" validate:"nonzero"`
	Url                             *string       `json:"url,omitempty"`
	About                           *string       `json:"about,omitempty"`
	ContactEmail                    *string       `json:"contactEmail,omitempty"`
	Phone                           *string       `json:"phone,omitempty"`
	UserId                          *string       `json:"userId,omitempty" validate:"nonzero"`
	TermsOfService                  *string       `json:"termsOfService,omitempty"`
	PrivacyPolicy                   *string       `json:"privacyPolicy,omitempty"`
	RevShareRate                    *float64      `json:"revShareRate,omitempty"`
	CcTransactionFeeRate            *float64      `json:"ccTransactionFeeRate,omitempty"`
	MinAppFeeRate                   *float64      `json:"minAppFeeRate,omitempty"`
	EnableConsolidatedBilling       *bool         `json:"enableConsolidatedBilling,omitempty"`
	ShortName                       *string       `json:"shortName,omitempty" validate:"nonzero"`
	EnablePurchaseOrder             *bool         `json:"enablePurchaseOrder,omitempty"`
	EnableEmailNotificationsToUsers *bool         `json:"enableEmailNotificationsToUsers,omitempty"`
	ParentTenantId                  *int64        `json:"parentTenantId,omitempty"`
	ExternalId                      *string       `json:"externalId,omitempty"`
	DefaultActivationProfileId      *string       `json:"defaultActivationProfileId,omitempty"`
	EnableMonthlyBilling            *bool         `json:"enableMonthlyBilling,omitempty"`
	DefaultChargeType               *string       `json:"defaultChargeType,omitempty"`
	LoginLogo                       *string       `json:"loginLogo,omitempty"`
	HomePageLogo                    *string       `json:"homePageLogo,omitempty"`
	DomainName                      *string       `json:"domainName,omitempty"`
	SkipDefaultUserSecurityGroup    *bool         `json:"skipDefaultUserSecurityGroup,omitempty"`
	DisableAllEmailNotification     *bool         `json:"disableAllEmailNotification,omitempty"`
	TrademarkURL                    *string       `json:"trademarkURL,omitempty"`
	Deleted                         *bool         `json:"deleted,omitempty"`
	Preferences                     *[]Preference `json:"preferences,omitempty"`
}

type Preference struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

func (s *Client) GetTenants() ([]Tenant, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data TenantAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	tenants := data.Tenants
	return tenants, nil
}

func (s *Client) GetTenant(id int) (*Tenant, error) {

	var data Tenant

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(id))
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

	tenant := &data
	return tenant, nil
}

func (s *Client) AddTenant(tenant *Tenant) error {

	if errs := validator.Validate(tenant); errs != nil {
		return errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants")

	j, err := json.Marshal(tenant)

	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)

	return err
}

func (s *Client) UpdateTenant(tenant *Tenant) (*Tenant, error) {

	var data Tenant

	if errs := validator.Validate(tenant); errs != nil {
		return nil, errs
	}

	if nonzero(tenant.Id) {
		return nil, errors.New("Tenant.Id is missing")
	}

	tenantId := *tenant.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + tenantId)

	j, err := json.Marshal(tenant)

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

	tenant = &data

	return tenant, nil
}

func (s *Client) DeleteTenantSync(tenantId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)

	if err != nil {

		byt := []byte(err.Error())

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			return err
		}

		msg := dat["msg"].(string)

		if msg == "Delete tenant request accepted" {
			return errors.New("Delete tenant request accepted. The tenant deletion is successful only when the following conditions are completed: \n\n - All the running jobs must be terminated for all users – users cannot be deleted before the jobs are terminated.\n\n - All users in the tenant are deleted \n\n - All the sub tenants under the tenant must be deleted prior to issuing this API call. If any sub-tenant is not deleted, then a validation message states that you do this first.\n ")
		}

		return err
	}

	return nil
}

func (s *Client) DeleteTenantAsync(tenantId int) (*OperationStatus, error) {

	var data OperationStatus

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
