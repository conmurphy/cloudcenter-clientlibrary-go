package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"

import "bytes"

//import "errors"

//RoleAPIResponse
type RoleAPIResponse struct {
	Resource      string `json:"resource,omitempty"`
	Size          int    `json:"size,omitempty"`
	PageNumber    int    `json:"pageNumber,omitempty"`
	TotalElements int    `json:"totalElements,omitempty"`
	TotalPages    int    `json:"totalPages,omitempty"`
	Roles         []Role `json:"roles,omitempty"`
}

type Role struct {
	Id          string       `json:"id,omitempty"`
	Resource    string       `json:"resource,omitempty"`
	Perms       []string     `json:"perms,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	TenantId    string       `json:"tenantId,omitempty"` //required
	ObjectPerms []ObjectPerm `json:"objectPerms,omitempty"`
	Users       []User       `json:"users,omitempty"`
	Groups      []Group      `json:"groups,omitempty"`
	OobRole     bool         `json:"oobRole,omitempty"`
	LastUpdated int64        `json:"lastUpdated,omitempty"`
	CreatedB    int64        `json:"created,omitempty"`
}

type ObjectPerm struct {
	ObjectType string   `json:"objectType,omitempty"`
	Perms      []string `json:"perms,omitempty"`
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

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + role.TenantId + "/roles")

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

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + role.TenantId + "/roles/" + role.Id)

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
