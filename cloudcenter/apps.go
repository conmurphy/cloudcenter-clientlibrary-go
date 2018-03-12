package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

//import "errors"

type AppAPIResponse struct {
	Apps []App `json:"apps"`
}

type App struct {
	Id              *string           `json:"id,omitempty"`
	Resource        *string           `json:"resource,omitempty"`
	Perms           *[]string         `json:"perms,omitempty"`
	Name            *string           `json:"name,omitempty"`
	Description     *string           `json:"description,omitempty"`
	ServiceTierId   *string           `json:"serviceTierId,omitempty"`
	Versions        *[]string         `json:"versions,omitempty"`
	Version         *string           `json:"version,omitempty"`
	Executor        *string           `json:"executor,omitempty"`
	Category        *string           `json:"category,omitempty"`
	ServiceTiers    *[]App            `json:"serviceTiers,omitempty"`
	ProfileCategory *string           `json:"profileCategory,omitempty"`
	Service         *Service          `json:"service,omitempty"`
	Clusterable     *bool             `json:"clusterable,omitempty"`
	HWProfile       *HWProfile        `json:"hwprofile,omitempty"`
	ParameterSpecs  *ParameterSpecs   `json:"parameterSpecs,omitempty"`
	Parameters      *Parameters       `json:"parameters,omitempty"`
	RevisionId      *int64            `json:"revisionId,omitempty"`
	Metadatas       *[]Metadata       `json:"metadatas,omitempty"`
	AppCategories   *[]AppCategory    `json:"appCategories,omitempty"`
	LogoPath        *string           `json:"logoPath,omitempty"`
	SupportedClouds *[]SupportedCloud `json:"supportedClouds,omitempty"`
}

type Parameters struct {
	AppParams *[]AppParam `json:"appParams,omitempty"`
	EnvParams *[]EnvParam `json:"envParams,omitempty"`
}

type AppCategory struct {
	Id       *string   `json:"id,omitempty"`
	Resource *string   `json:"resource,omitempty"`
	Perms    *[]string `json:"perms,omitempty"`
	Name     *string   `json:"name,omitempty"`
	Type     *string   `json:"type,omitempty"`
}

type SupportedCloud struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type HWProfile struct {
	MemorySize               *int64 `json:"memorySize,omitempty"`
	NumOfCPUs                *int64 `json:"numOfCPUs,omitempty"`
	NetworkSpeed             *int64 `json:"networkSpeed,omitempty"`
	NumOfNICs                *int64 `json:"numOfNICs,omitempty"`
	LocalStorageCount        *int64 `json:"localStorageCount,omitempty"`
	LocalStorageSize         *int64 `json:"localStorageSize,omitempty"`
	CudaSupport              *bool  `json:"cudaSupport,omitempty"`
	SSDSupport               *bool  `json:"ssdSupport,omitempty"`
	SupportHardwareProvision *bool  `json:"supportHardwareProvision,omitempty"`
}

type ParameterSpecs struct {
	SystemParams *SystemParams `json:"SystemParams,omitempty"`
	CustomParams *CustomParams `json:"customParams,omitempty"`
	EnvVars      *EnvVar       `json:"envVars,omitempty"`
}

type EnvVar struct {
	EnvVars *[]EnvVar `json:"envVars,omitempty"`
	Size    *int64    `json:"size,omitempty"`
}

type SystemParams struct {
	Params *[]Param `json:"params,omitempty"`
	Size   *int64   `json:"size,omitempty"`
}

type CustomParams struct {
	Params *[]Param `json:"params,omitempty"`
	Size   *int64   `json:"size,omitempty"`
}

type Param struct {
	ParamName            *string           `json:"paramName,omitempty"`
	DisplayName          *string           `json:"displayName,omitempty"`
	HelpText             *string           `json:"helpText,omitempty"`
	Type                 *string           `json:"type,omitempty"`
	ValueList            *string           `json:"valueList,omitempty"`
	DefaultValue         *string           `json:"defaultValue,omitempty"`
	ConfirmValue         *string           `json:"confirmValue,omitempty"`
	PathSuffixValue      *string           `json:"pathSuffixValue,omitempty"`
	UserVisible          *bool             `json:"userVisible,omitempty"`
	UserEditable         *bool             `json:"userEditable,omitempty"`
	SystemParam          *bool             `json:"systemParam,omitempty"`
	ExampleValue         *string           `json:"exampleValue,omitempty"`
	DataUnit             *string           `json:"dataUnit,omitempty"`
	Optional             *bool             `json:"optional,omitempty"`
	MultiselectSupported *bool             `json:"multiselectSupported,omitempty"`
	ValueConstraint      *ValueConstraint  `json:"valueConstraint,omitempty"`
	Scope                *string           `json:"scope,omitempty"`
	WebserviceListParams *string           `json:"webserviceListParams,omitempty"`
	CollectionList       *[]CollectionList `json:"collectionList,omitempty"`
}

type ValueConstraint struct {
	MinValue            *int64  `json:"minValue,omitempty"`
	MaxValue            *int64  `json:"maxValue,omitempty"`
	MaxLength           *int64  `json:"maxLength,omitempty"`
	Regex               *string `json:"regex,omitempty"`
	AllowSpaces         *bool   `json:"allowSpaces,omitempty"`
	SizeValue           *int64  `json:"sizeValue,omitempty"`
	Step                *int64  `json:"step,omitempty"`
	CalloutWorkflowName *string `json:"calloutWorkflowName,omitempty"`
}

type CollectionList struct {
	ParamCollectionItem []ParamCollectionItem `json:"paramCollectionItem,omitempty"`
}

type ParamCollectionItem struct {
	CollectionType         *string `json:"collectionType,omitempty"`
	CollectionName         *string `json:"collectionName,omitempty"`
	CollectionDisplayName  *string `json:"collectionDisplayName,omitempty"`
	CollectionValue        *string `json:"collectionValue,omitempty"`
	CollectionDefaultValue *string `json:"collectionDefaultValue,omitempty"`
	CollectionHelpText     *string `json:"collectionHelpText,omitempty"`
	CollectionSampleText   *string `json:"collectionSampleText,omitempty"`
	Optional               *bool   `json:"optional,omitempty"`
}

func (s *Client) GetApps() ([]App, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/apps")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data AppAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	apps := data.Apps
	return apps, nil
}

func (s *Client) GetApp(appId int) (*App, error) {

	var data App

	url := fmt.Sprintf(s.BaseURL + "/v1/apps/" + strconv.Itoa(appId))
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

	app := &data
	return app, nil
}

func (s *Client) ImportApp(filename string) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/apps/portation")

	body, err := s.sendFile(filename, url)

	fmt.Println(string(body))
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) UpdateApp(app *App) error {

	appId := *app.Id
	appVersion := *app.Version

	url := fmt.Sprintf(s.BaseURL + "/v1/apps/" + appId + "?version=" + appVersion)

	j, err := json.Marshal(app)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	_, err = s.doRequest(req)

	if err != nil {

		return err
	}

	return nil
}

func (s *Client) DeleteApp(appId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/apps/" + strconv.Itoa(appId))

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
