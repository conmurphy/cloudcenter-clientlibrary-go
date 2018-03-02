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
import "io"
import "mime/multipart"
import "os"

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

func (s *Client) sendFile(filename string, url string) ([]byte, error) {
	r, w := io.Pipe()
	writer := multipart.NewWriter(w)
	go func() {
		part, err := writer.CreateFormFile("file", filename)
		if err != nil {
			w.CloseWithError(err)
			return
		}
		_, err = io.Copy(part, os.Stdin)
		if err != nil {
			w.CloseWithError(err)
			return
		}
		err = writer.Close()
		if err != nil {
			w.CloseWithError(err)
			return
		}
	}()

	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

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

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Bool(value bool) *bool {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int(value int) *int {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int64(value int64) *int64 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func String(value string) *string {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float32(value float32) *float32 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float64(value float64) *float64 {
	return &value
}
