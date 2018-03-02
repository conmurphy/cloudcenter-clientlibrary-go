package cloudcenter

import "fmt"
import "net/http"

import "encoding/json"
import "strconv"
import "bytes"

type ServiceAPIResponse struct {
	Resource      *string   `json:"resource,omitempty"`
	Size          *int      `json:"size,omitempty"`
	PageNumber    *int      `json:"pageNumber,omitempty"`
	TotalElements *int      `json:"totalElements,omitempty"`
	TotalPages    *int      `json:"totalPages,omitempty"`
	Services      []Service `json:"services,omitempty"`
}

type Service struct {
	Id                     *string             `json:"id,omitempty"`
	OwnerUserId            *string             `json:"ownerUserId,omitempty"`
	TenantId               *string             `json:"tenantId,omitempty"`
	ParentService          *bool               `json:"parentService,omitempty"`
	ParentServiceId        *string             `json:"parentServiceId,omitempty"`
	Resource               *string             `json:"resource,omitempty"`
	Perms                  *[]string           `json:"perms,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	DisplayName            *string             `json:"displayName,omitempty"`
	LogoPath               *string             `json:"logoPath,omitempty"`
	Description            *string             `json:"description,omitempty"`
	DefaultImageId         *int64              `json:"defaultImageId,omitempty"`
	ServiceType            *string             `json:"serviceType,omitempty"`
	SystemService          *bool               `json:"systemService,omitempty"`
	ExternalService        *bool               `json:"externalService,omitempty"`
	Visible                *bool               `json:"visible,omitempty"`
	ExternalBundleLocation *string             `json:"externalBundleLocation,omitempty"`
	BundleLocation         *string             `json:"bundleLocation,omitempty"`
	CostPerHour            *float64            `json:"costPerHour,omitempty"`
	OwnerId                *string             `json:"ownerId,omitempty"`
	ServiceActions         []ServiceAction     `json:"serviceActions,omitempty"`
	ServicePorts           []ServicePort       `json:"servicePorts,omitempty"`
	ServiceParamSpecs      []ServiceParamSpec  `json:"serviceParamSpecs,omitempty"`
	EgressRestrictions     []EgressRestriction `json:"egressRestrictions,omitempty"`
	Images                 []Image             `json:"images,omitempty"`
	Repositories           []Repository        `json:"repositories,omitempty"`
	ChildServices          []Service           `json:"childServices,omitempty"`
	ExternalActions        []ExternalAction    `json:"externalActions,omitempty"`
}

type ServiceAction struct {
	ActionName  *string `json:"actionName,omitempty"`
	ActionType  *string `json:"actionType,omitempty"`
	ActionValue *string `json:"actionValue,omitempty"`
}

type ServicePort struct {
	Protocol *string `json:"protocol,omitempty"`
	FromPort *string `json:"fromPort,omitempty"`
	ToPort   *string `json:"toPort,omitempty"`
	CloudId  *string `json:"cloudId,omitempty"`
}

type ServiceParamSpec struct {
	ParamName            *string               `json:"paramName,omitempty"`
	DisplayName          *string               `json:"displayName,omitempty"`
	HelpText             *string               `json:"helpText,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	ValueList            *string               `json:"valueList,omitempty"`
	WebserviceListParams []WebserviceListParam `json:"webserviceListParams,omitempty"`
	DefaultValue         *string               `json:"defaultValue,omitempty"`
	UserVisible          *bool                 `json:"userVisible,omitempty"`
	UserEditable         *bool                 `json:"userEditable,omitempty"`
	SystemParam          *bool                 `json:"systemParam,omitempty"`
	ExampleValue         *string               `json:"exampleValue,omitempty"`
	Optional             *bool                 `json:"optional,omitempty"`
	ValueConstraint      ValueConstraint       `json:"valueConstraint,omitempty"`
}

type EgressRestriction struct {
	EgressServiceName *string `json:"egressServiceName,omitempty"`
}

type Repository struct {
	Id          *string   `json:"id,omitempty"`
	Resource    *string   `json:"resource,omitempty"`
	Perms       *[]string `json:"perms,omitempty"`
	Hostname    *string   `json:"hostname,omitempty"`
	DisplayName *string   `json:"displayName,omitempty"`
	Protocol    *string   `json:"protocol,omitempty"`
	Description *string   `json:"description,omitempty"`
	Port        *int64    `json:"port,omitempty"`
}

func (s *Client) GetServices(tenantId int) ([]Service, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ServiceAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	service := data.Services
	return service, nil
}

func (s *Client) GetService(tenantId int, serviceId int) (*Service, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services/" + strconv.Itoa(serviceId))

	var data Service

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

	service := &data

	return service, nil
}

func (s *Client) AddService(service *Service) (*Service, error) {

	var data Service

	serviceTenantId := *service.TenantId
	serviceId := *service.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + serviceTenantId + "/services/" + serviceId)

	j, err := json.Marshal(service)

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

	service = &data

	return service, nil

}

func (s *Client) UpdateService(service *Service) (*Service, error) {

	var data Service

	serviceTenantId := *service.TenantId
	serviceId := *service.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + serviceTenantId + "/services/" + serviceId)

	j, err := json.Marshal(service)

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

	service = &data

	return service, nil
}

func (s *Client) DeleteService(tenantId int, serviceId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/services/" + strconv.Itoa(serviceId))

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
