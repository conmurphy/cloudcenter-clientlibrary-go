package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"

type JobAPIResponse struct {
	Resource      string `json:"resource"`
	Size          int    `json:"size"`
	PageNumber    int    `json:"pageNumber"`
	TotalElements int    `json:"totalElements"`
	TotalPages    int    `json:"totalPages"`
	Jobs          []Job  `json:"jobs"`
}

type Job struct {
	Id                     string                `json:"id"`
	Resource               string                `json:"resource"`
	Name                   string                `json:"name"`
	Description            string                `json:"description"`
	Status                 string                `json:"status"`
	JobStatusMessage       string                `json:"jobStatusMessage"`
	StartTime              string                `json:"startTime"`
	EndTime                string                `json:"endTime"`
	FavoriteCreationTime   string                `json:"favouriteCreationTime"`
	CloudFamily            string                `json:"cloudFamily"`
	AgentUpgradeInProgress bool                  `json:"agentUpgradeInProgress"`
	DeploymentEnvironment  DeploymentEnvironment `json:"deploymentEnvironment"`
	Application            Application           `json:"application"`
	DeploymentEntity       DeploymentEntity      `json:"deploymentEntity"`
	TerminateProtection    bool                  `json:"terminateProtection"`
	Hidden                 bool                  `json:"hidden"`
	Favorite               bool                  `json:"favorite"`
	Benchmark              bool                  `json:"benchmark"`
	Owner                  bool                  `json:"owner"`
	OwnerEmailAddress      string                `json:"ownerEmailAddress"`
	TotalCost              float32               `json:"totalCost"`
	NodeHours              float32               `json:"nodeHours"`
	AppId                  string                `json:"AppId,omitempty"`
	AppName                string                `json:"AppName,omitempty"`
}

type DeploymentEnvironment struct {
	Id       string `json:"id"`
	Resource string `json:"resource"`
}

type Application struct {
	Id       string `json:"id"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}

type DeploymentEntity struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type JobsRequest struct {
	Name                   string                 `json:"name,omitempty"`
	PolicyId               string                 `json:"policyId,omitempty"`
	AppId                  string                 `json:"appId,omitempty"`
	Metadatas              []MetadataRequest      `json:"metadatas,omitempty"`
	EnvironmentId          string                 `json:"environmentId,omitempty"`
	AppVersion             string                 `json:"appVersion,omitempty"`
	KeepExistingDeployment bool                   `json:"keepExistingDeployment,omitempty"`
	TagIds                 []float32              `json:"tagIds,omitempty"`
	Parameters             ParameterRequest       `json:"parameters,omitempty"`
	Jobs                   []JobRequest           `json:"jobs,omitempty"`
	CloudProperties        []CloudPropertyRequest `json:"cloudProperties,omitempty"`
}

type MetadataRequest struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Editable  bool   `json:"editable,omitempty"`
	Required  bool   `json:"required,omitempty"`
}

type JobRequest struct {
	TierId     string           `json:"tierId,omitempty"`
	NodeIds    string           `json:"nodeIds,omitempty"`
	Parameters ParameterRequest `json:"parameters,omitempty"`
}

type ParameterRequest struct {
	CloudParams []CloudParamRequest `json:"cloudParams,omitempty"`
	AppParams   []AppParamRequest   `json:"appParams,omitempty"`
	EnvParams   []EnvParamRequest   `json:"envParams,omitempty"`
}

type CloudParamRequest struct {
	Cloud           string                 `json:"cloud,omitempty"`
	Instance        string                 `json:"instance,omitempty"`
	Storage         []StorageRequest       `json:"storage,omitempty"`
	RootVolumeSize  string                 `json:"rootVolumeSize,omitempty"`
	CloudProperties []CloudPropertyRequest `json:"cloudProperties,omitempty"`
}

type StorageRequest struct {
	RegionId              string                        `json:"regionId,omitempty"`
	CloudAccountId        string                        `json:"cloudAccountId,omitempty"`
	Size                  int                           `json:"size,omitempty"`
	NumNodes              int                           `json:"numNodes,omitempty"`
	CloudSpecificSettings []CloudSpecificSettingRequest `json:"cloudSpecificSettings,omitempty"`
	Address               string                        `json:"address,omitempty"`
}

type CloudSpecificSettingRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CloudPropertyRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type AppParamRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type EnvParamRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (s *Client) GetJobs() ([]Job, error) {

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data JobAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	jobs := data.Jobs
	return jobs, nil
}

func (s *Client) GetJob(id int) (*Job, error) {

	var data Job

	url := fmt.Sprintf(s.BaseURL + "/v1/jobs/" + strconv.Itoa(id))
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

	job := &data
	return job, nil
}
