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

//import "errors"

//RoleAPIResponse
type RoleAPIResponse struct {
	Resource      *string `json:"resource,omitempty"`
	Size          *int64  `json:"size,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	TotalElements *int64  `json:"totalElements,omitempty"`
	TotalPages    *int64  `json:"totalPages,omitempty"`
	Roles         []Role  `json:"roles,omitempty"`
}

type Role struct {
	Id          *string       `json:"id,omitempty"`
	Resource    *string       `json:"resource,omitempty"`
	Perms       *[]string     `json:"perms,omitempty"`
	Name        *string       `json:"name,omitempty" validate:"nonzero"`
	Description *string       `json:"description,omitempty"`
	TenantId    *string       `json:"tenantId,omitempty" validate:"nonzero"`
	ObjectPerms *[]ObjectPerm `json:"objectPerms,omitempty"`
	Users       *[]User       `json:"users,omitempty"`
	Groups      *[]Group      `json:"groups,omitempty"`
	OobRole     *bool         `json:"oobRole,omitempty"`
	LastUpdated *int64        `json:"lastUpdated,omitempty"`
	Created     *int64        `json:"created,omitempty"`
}

type ObjectPerm struct {
	ObjectType *string   `json:"objectType,omitempty"`
	Perms      *[]string `json:"perms,omitempty"`
}

func (s *Client) GetRoles(tenantId int) ([]Role, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/roles/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data RoleAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	roles := data.Roles
	return roles, nil
}

func (s *Client) GetRole(tenantId int, roleId int) (*Role, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/roles/" + strconv.Itoa(roleId))

	var data Role

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

	role := &data

	return role, nil
}

func (s *Client) AddRole(role *Role) (*Role, error) {

	var data Role

	if errs := validator.Validate(role); errs != nil {
		return nil, errs
	}

	roleTenantId := *role.TenantId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + roleTenantId + "/roles")

	j, err := json.Marshal(role)

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

	role = &data

	return role, nil
}

func (s *Client) UpdateRole(role *Role) (*Role, error) {

	var data Role

	if errs := validator.Validate(role); errs != nil {
		return nil, errs
	}

	roleTenantId := *role.TenantId

	if nonzero(role.Id) {
		return nil, errors.New("Role.Id is missing")
	}

	roleId := *role.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + roleTenantId + "/roles/" + roleId)

	j, err := json.Marshal(role)

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

	role = &data

	return role, nil
}

func (s *Client) DeleteRole(tenantId int, roleId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/roles/" + strconv.Itoa(roleId))

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
