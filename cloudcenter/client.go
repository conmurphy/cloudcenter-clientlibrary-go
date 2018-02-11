//
// This client library provides Create, Read, Update, and Delete operations for Cisco Cloud Center.
//
// A Basic Example
//
//  package main
//
//  import "github.com/cloudcenter-clientlibrary-go/cloudcenter”
//
//  // Define new cloudcenter client
//
//  client := cloudcenter.NewClient("cliqradmin", ”myAPIKey", "https://ccm.dcloud.cisco.com")
//
//  // Create user
//
//  newUser := cloudcenter.User{
//	  TenantId:    "1",
//	  FirstName:   "client",
//	  LastName:    "library",
//	  Password:    "myPassword",
//	  EmailAddr:   "clientlibrary@cloudcenter.com",
//	  CompanyName: "Company",
//	  PhoneNumber: "12345",
//	  ExternalId:  "23456",
//  }
//
//  user, err := client.AddUser(&newUser)
//
//  if err != nil {
//	  fmt.Println(err)
//  } else {
//	  fmt.Println(”New user created. \n UserId: " + user.Id + ", Username: " + user.LastName)
//  }
package cloudcenter

import "fmt"
import "net/http"
import "io/ioutil"
import "crypto/tls"

//import "encoding/json"

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

	req.Header.Add("Content-Type", "application/json")
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

	if 200 != resp.StatusCode && 201 != resp.StatusCode && 202 != resp.StatusCode && 204 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
