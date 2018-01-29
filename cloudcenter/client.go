package cloudcenter

import "fmt"
import "net/http"
import "io/ioutil"
import "crypto/tls"

type Client struct {
	Username string
	Password string
	BaseURL  string
}

func NewClient(username, password, baseURL string) *Client {
	return &Client{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
