package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

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
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	EntityType  *string    `json:"entityType,omitempty"`
	EventName   *string    `json:"eventName,omitempty"`
	Actions     *[]Actions `json:"actions,omitempty"`
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
