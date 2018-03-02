package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

//import "bytes"

type ActionAPIResponse struct {
	Resource      string   `json:"resource,omitempty"`
	Size          int      `json:"size,omitempty"`
	PageNumber    int      `json:"pageNumber,omitempty"`
	TotalElements int      `json:"totalElements,omitempty"`
	TotalPages    int      `json:"totalPages"`
	ActionJaxbs   []Action `json:"actionJaxbs,omitempty"`
}

type Action struct {
	Id                     string                  `json:"id,omitempty"`
	Resource               string                  `json:"resource,omitempty"`
	Perms                  []string                `json:"perms,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	Description            string                  `json:"description,omitempty"`
	ActionType             string                  `json:"actionType,omitempty"`
	LastUpdatedTime        string                  `json:"lastUpdatedTime,omitempty"`
	TimeOut                float32                 `json:"timeOut,omitempty"`
	Enabled                bool                    `json:"enabled,omitempty"`
	Encrypted              bool                    `json:"encrypted,omitempty"`
	Deleted                bool                    `json:"deleted,omitempty"`
	SystemDefined          bool                    `json:"systemDefined,omitempty"`
	BulkOperationSupported bool                    `json:"bulkOperationSupported,omitempty"`
	IsAvailableToUser      bool                    `json:"isAvailableToUser,omitempty"`
	Owner                  int                     `json:"owner,omitempty"`
	ActionParameters       []ActionParameter       `json:"actionParameters,omitempty"`
	ActionResourceMappings []ActionResourceMapping `json:"actionResourceMappings,omitempty"`
	ActionCustomParamSpecs []ActionCustomParamSpec `json:"actionCustomParamSpecs,omitempty"`
}

type ActionParameter struct {
	ParamName   string `json:"paramName,omitempty"`
	ParamValue  string `json:"paramValue,omitempty"`
	CustomParam bool   `json:"customParam,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Preference  string `json:"preference,omitempty"`
}

type ActionResourceMapping struct {
	Type                  string                 `json:"type,omitempty"`
	ActionResourceFilters []ActionResourceFilter `json:"actionResourceFilters,omitempty"`
}

type ActionResourceFilter struct {
	DeploymentResource string     `json:"deploymentResource,omitempty"`
	VmResource         VmResource `json:"vmResource,omitempty"`
	IsEditable         bool       `json:"isEditable,omitempty"`
}

type VmResource struct {
	Type                  string   `json:"type,omitempty"`
	AppProfiles           []string `json:"appProfiles,omitempty"`
	CloudRegions          []string `json:"cloudRegions,omitempty"`
	CloudAccounts         []string `json:"cloudAccounts,omitempty"`
	Services              []string `json:"services,omitempty"`
	OsTypes               []string `json:"osTypes,omitempty"`
	CloudFamilyNames      []string `json:"cloudFamilyNames,omitempty"`
	NodeStates            []string `json:"nodesStates,omitempty"`
	CloudResourceMappings []string `json:"cloudResourceMappings,omitempty"`
}

type ActionCustomParamSpec struct {
	ParamName            string              `json:"paramName,omitempty"`
	DisplayName          string              `json:"displayName,omitempty"`
	HelpText             string              `json:"helpText,omitempty"`
	Type                 string              `json:"type,omitempty"`
	ValueList            string              `json:"valueList,omitempty"`
	DefaultValue         string              `json:"defaultValue,omitempty"`
	ConfirmValue         string              `json:"confirmValue,omitempty"`
	PathSuffixValue      string              `json:"pathSuffixValue,omitempty"`
	UserVisible          bool                `json:"userVisible,omitempty"`
	UserEditable         bool                `json:"userEditable,omitempty"`
	SystemParam          bool                `json:"systemParam,omitempty"`
	ExampleValue         string              `json:"exampleValue,omitempty"`
	DataUnit             string              `json:"dataUnit,omitempty"`
	Optional             bool                `json:"optional,omitempty"`
	MultiselectSupported bool                `json:"multiselectSupported,omitempty"`
	ValueConstraint      ValueConstraint     `json:"valueConstraint,omitempty"`
	Scope                string              `json:"scope,omitempty"`
	WebserviceListParams WebserviceListParam `json:"webserviceListParams,omitempty"`
	Preference           string              `json:"preference,omitempty"`
}

type WebserviceListParam struct {
	URL           string `json:"url,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	RequestType   string `json:"requestType,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	CommandParams string `json:"commandParams,omitempty"`
	RequestBody   string `json:"requestBody,omitempty"`
	ResultString  string `json:"resultString,omitempty"`
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

	url := fmt.Sprintf(s.BaseURL + "/v1/actions/" + action.Id)

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
