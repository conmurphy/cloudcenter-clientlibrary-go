package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"

type AppAPIResponse struct {
	Apps []App `json:"apps"`
}

type App struct {
	Id            string   `json:"id"`
	Resource      string   `json:"resource"`
	Perms         []string `json:"perms"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	ServiceTierId string   `json:"serviceTierId"`
	Versions      []string `json:"versions"`
	Executor      string   `json:"executor"`
	Category      string   `json:"category"`
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
