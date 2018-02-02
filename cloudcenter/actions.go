package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"

type ActionAPIResponse struct {
	Resource      string   `json:"resource"`
	Size          int      `json:"size"`
	PageNumber    int      `json:"pageNumber"`
	TotalElements int      `json:"totalElements"`
	TotalPages    int      `json:"totalPages"`
	ActionJaxbs   []Action `json:"actionJaxbs"`
}

type Action struct {
	Id                     string                  `json:"id"`
	Resource               string                  `json:"resource"`
	Perms                  []string                `json:"perms"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	ActionType             string                  `json:"actionType"`
	LastUpdatedTime        string                  `json:"lastUpdatedTime"`
	TimeOut                float32                 `json:"timeOut"`
	Enabled                bool                    `json:"enabled"`
	Encrypted              bool                    `json:"encrypted	"`
	Deleted                bool                    `json:"deleted"`
	SystemDefined          bool                    `json:"systemDefined"`
	BulkOperationSupported bool                    `json:"bulkOperationSupported"`
	IsAvailableToUser      bool                    `json:"isAvailableToUser"`
	Owner                  int                     `json:"owner"`
	ActionParameters       []ActionParameter       `json:"actionParameters"`
	ActionResourceMappings []ActionResourceMapping `json:"actionResourceMappings"`
	ActionCustomParamSpecs []ActionCustomParamSpec `json:"actionCustomParamSpecs"`
}

type ActionParameter struct {
	ParamName   string `json:"paramName"`
	ParamValue  string `json:"paramValue"`
	CustomParam bool   `json:"customParam"`
	Required    bool   `json:"required"`
	Preference  string `json:"preference"`
}

type ActionResourceMapping struct {
	Type                  string                 `json:"type"`
	ActionResourceFilters []ActionResourceFilter `json:"actionResourceFilters"`
}

type ActionResourceFilter struct {
	DeploymentResource string     `json:"deploymentResource"`
	VmResource         VmResource `json:"vmResource"`
	IsEditable         bool       `json:"isEditable"`
}

type VmResource struct {
	Type                  string   `json:"type"`
	AppProfiles           []string `json:"appProfile"`
	CloudRegions          []string `json:"cloudRegions"`
	CloudAccounts         []string `json:"cloudAccounts"`
	Services              []string `json:"services"`
	OsTypes               []string `json:"osTypes"`
	CloudFamilyNames      []string `json:"cloudFamilyNames"`
	NodeStates            []string `json:"nodesStates"`
	CloudResourceMappings []string `json:"cloudResourceMappings"`
}

type ActionCustomParamSpec struct {
	ParamName            string              `json:"paramName"`
	DisplayName          string              `json:"displayName"`
	HelpText             string              `json:"helpText"`
	Type                 string              `json:"type"`
	ValueList            string              `json:"valueList"`
	DefaultValue         string              `json:"defaultValue"`
	ConfirmValue         string              `json:"confirmValue"`
	PathSuffixValue      string              `json:"pathSuffixValue"`
	UserVisible          bool                `json:"userVisible"`
	UserEditable         bool                `json:"userEditable"`
	SystemParam          bool                `json:"systemParam"`
	ExampleValue         string              `json:"exampleValue"`
	DataUnit             string              `json:"dataUnit"`
	Optional             bool                `json:"optional"`
	MultiselectSupported bool                `json:"multiselectSupported"`
	ValueConstraint      ValueConstraint     `json:"valueConstraint"`
	Scope                string              `json:"scope"`
	WebserviceListParams WebserviceListParam `json:"webserviceListParams"`
	Preference           string              `json:"preference"`
}

type ValueConstraint struct {
	MinValue            int    `json:"minValue"`
	MaxValue            int    `json:"maxValue"`
	MaxLength           int    `json:"maxLength"`
	Regex               string `json:"regex"`
	AllowSpaces         bool   `json:"allowSpaces"`
	SizeValue           int    `json:"sizeValue"`
	Step                int    `json:"step"`
	CalloutWorkflowName string `json:"calloutWorkflowName"`
}

type WebserviceListParam struct {
	Url           string `json:"url"`
	Protocol      string `json:"protocol"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	RequestType   string `json:"requestType"`
	ContentType   string `json:"contentType"`
	CommandParams string `json:"commandParams"`
	RequestBody   string `json:"requestBody"`
	ResultString  string `json:"resultString"`
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
