package cloudcenterclient

import "fmt"
import "net/http"
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
	TerminateProtection    bool                  `json:"terminateProtection "`
	Hidden                 bool                  `json:"hidden"`
	Favorite               bool                  `json:"favorite"`
	Benchmark              bool                  `json:"benchmark "`
	Owner                  bool                  `json:"owner "`
	OwnerEmailAddress      string                `json:"ownerEmailAddress "`
	TotalCost              int                   `json:"totalCost "`
	NodeHours              float32               `json:"nodeHours "`
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
