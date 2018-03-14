package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

//import "bytes"

type SuspensionPolicyAPIResponse struct {
	Resource           *string            `json:"resource,omitempty"`
	Size               *int64             `json:"size,omitempty"`
	PageNumber         *int64             `json:"pageNumber,omitempty"`
	TotalElements      *int64             `json:"totalElements,omitempty"`
	TotalPages         *int64             `json:"totalPages"`
	SuspensionPolicies []SuspensionPolicy `json:"suspensionpolicies,omitempty"`
}

type SuspensionPolicy struct {
	Id                        *string           `json:"id,omitempty"`
	Resource                  *string           `json:"resource,omitempty"`
	Perms                     *[]string         `json:"perms,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	Description               *string           `json:"description,omitempty"`
	Enabled                   *bool             `json:"enabled,omitempty"`
	Schedules                 *[]Schedule       `json:"schedules,omitempty"`
	BlockoutPeriods           *[]BlockoutPeriod `json:"blockoutPeriods,omitempty"`
	IsPolicyActiveOnResources *bool             `json:"isPolicyActiveOnResources,omitempty"`
	ResourcesMaps             *[]ResourcesMap   `json:"resourcesMaps,omitempty"`
	Priority                  *float64          `json:"priority,omitempty"`
	Created                   *float64          `json:"created,omitempty"`
	LastUpdated               *float64          `json:"lastUpdated,omitempty"`
}

type Schedule struct {
	Type      *string `json:"type,omitempty"`
	Days      *string `json:"days,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
	Repeats   *int64  `json:"repeats,omitempty"`
}

type BlockoutPeriod struct {
	StartDate *float64 `json:"startDate,omitempty"`
	EndDate   *float64 `json:"endDate,omitempty"`
}

type ResourcesMap struct {
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

func (s *Client) GetSuspensionPolicies() ([]SuspensionPolicy, error) {

	var data SuspensionPolicyAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v2/suspensionpolicies")

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

	suspensionPolicy := data.SuspensionPolicies
	return suspensionPolicy, nil
}

func (s *Client) GetSuspensionPolicy(suspensionPolicyId int) (*SuspensionPolicy, error) {

	var data SuspensionPolicy

	url := fmt.Sprintf(s.BaseURL + "/v2/suspensionpolicies/" + strconv.Itoa(suspensionPolicyId))
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

	suspensionPolicy := &data
	return suspensionPolicy, nil
}

func (s *Client) AddSuspensionPolicy(suspensionPolicy *SuspensionPolicy) (*SuspensionPolicy, error) {

	var data SuspensionPolicy

	url := fmt.Sprintf(s.BaseURL + "/v2/suspensionpolicies")

	j, err := json.Marshal(suspensionPolicy)

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

	suspensionPolicy = &data

	return suspensionPolicy, nil
}

func (s *Client) UpdateSuspensionPolicy(suspensionPolicy *SuspensionPolicy) (*SuspensionPolicy, error) {

	var data SuspensionPolicy

	suspensionPolicyId := *suspensionPolicy.Id
	url := fmt.Sprintf(s.BaseURL + "/v2/suspensionpolicies/" + suspensionPolicyId)

	j, err := json.Marshal(suspensionPolicy)

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

	suspensionPolicy = &data

	return suspensionPolicy, nil
}

func (s *Client) DeleteSuspensionPolicy(suspensionPolicyId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v2/suspensionpolicies/" + strconv.Itoa(suspensionPolicyId))

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
