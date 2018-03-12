package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"

import "bytes"

//import "errors"

//GroupAPIResponse
type GroupAPIResponse struct {
	Resource      *string `json:"resource,omitempty"`
	Size          *int    `json:"size,omitempty"`
	PageNumber    *int    `json:"pageNumber,omitempty"`
	TotalElements *int    `json:"totalElements,omitempty"`
	TotalPages    *int    `json:"totalPages,omitempty"`
	Groups        []Group `json:"groups,omitempty"`
}

type Group struct {
	Id           *string   `json:"id,omitempty"`
	Resource     *string   `json:"resource,omitempty"`
	Perms        *[]string `json:"perms,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Description  *string   `json:"description,omitempty"`
	TenantId     *string   `json:"tenantId,omitempty"`
	Users        *[]User   `json:"users,omitempty"`
	Roles        *[]Role   `json:"roles,omitempty"`
	Created      *int      `json:"created,omitempty"`
	LastUpdated  *int      `json:"lastUpdated,omitempty"`
	CreatedBySso *bool     `json:"createdBySso,omitempty"`
}

func (s *Client) GetGroups(tenantId int) ([]Group, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/groups/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data GroupAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	groups := data.Groups
	return groups, nil
}

func (s *Client) GetGroup(tenantId int, groupId int) (*Group, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/groups/" + strconv.Itoa(groupId))

	var data Group

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

	group := &data

	return group, nil
}

func (s *Client) AddGroup(group *Group) (*Group, error) {

	var data Group

	groupTenantId := *group.TenantId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + groupTenantId + "/groups")

	j, err := json.Marshal(group)

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

	group = &data

	return group, nil
}

func (s *Client) UpdateGroup(group *Group) (*Group, error) {

	var data Group

	groupTenantId := *group.TenantId
	groupId := *group.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + groupTenantId + "/groups/" + groupId)

	j, err := json.Marshal(group)

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

	group = &data

	return group, nil
}

func (s *Client) DeleteGroup(tenantId int, groupId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/groups/" + strconv.Itoa(groupId))

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
