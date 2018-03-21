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
//  		FirstName:   cloudcenter.String("client"),
//  		LastName:    cloudcenter.String("library"),
//  		Password:    cloudcenter.String("myPassword"),
//  		EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter.com"),
//  		CompanyName: cloudcenter.String("company"),
//  		PhoneNumber: cloudcenter.String("12345"),
//  		ExternalId:  cloudcenter.String("23456"),
//  		TenantId:    cloudcenter.String("1"),
//  	}
//
//  user, err := client.AddUser(&newUser)
//
//  if err != nil {
//	  fmt.Println(err)
//  } else {
//	  fmt.Println(”New user created. \n UserId: " + user.Id + ", Username: " + user.LastName)
//  }
//
// Helper Functions
//
// As per the following link, using the Marshal function from the encoding/json library treats false booleans as if they were nil values, and thus it omits them from the JSON response. To make a distinction between a non-existent boolean and false boolean we need to use a ```*bool``` in the struct.
//
// type User struct {
// 	Id                      *string `json:"id,omitempty"`
// 	FirstName               *string `json:"firstName,omitempty"`
// 	LastName                *string `json:"lastName,omitempty"`
// 	Password                *string `json:"password,omitempty"`
// 	EmailAddr               *string `json:"emailAddr,omitempty"`
// 	Enabled                 *bool   `json:"enabled,omitempty"`
// 	TenantAdmin             *bool   `json:"tenantAdmin,omitempty"`
// }
// https://github.com/golang/go/issues/13284
//
// Therefore in order to have a consistent experience all struct fields within this client library use pointers. This provides a way to differentiate between unset values, nil, and an intentional zero value, such as "", false, or 0.
//
// Helper functions have been created to simplify the creation of pointer types.
//
// Without
//
// firstName 	:= "client"
// lastName 	:= "library"//
// password	    := "myPassword"
// emailAddr	:= "clientlibrary@cloudcenter-address.com"
// companyName	:= "company"
// phoneNumber	:= "12345"
// externalId	:= "23456"
// tenantId	:= "1"
//
// newUser := cloudcenter.User {
// 	FirstName:   &firstName,
// 	LastName:    &lastName,
// 	Password:    &password,
// 	EmailAddr:  &emailAddr,
// 	CompanyName: &companyName,
// 	PhoneNumber: &phoneNumber,
// 	ExternalId: &externalId,
// 	TenantId:    &tenantId,
// }
//
// With
//
// newUser := cloudcenter.User {
// 	FirstName:   cloudcenter.String("client"),
// 	LastName:    cloudcenter.String("library"),
// 	Password:    cloudcenter.String("myPassword"),
// 	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
// 	CompanyName: cloudcenter.String("company"),
// 	PhoneNumber: cloudcenter.String("12345"),
// 	ExternalId:  cloudcenter.String("23456"),
// 	TenantId:    cloudcenter.String("1"),
// }
//
// Reference: https://willnorris.com/2014/05/go-rest-apis-and-pointers
package cloudcenter

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
)

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

//modified from unexported nonzero function in the validtor package
//https://github.com/go-validator/validator/blob/v2/builtins.go
func nonzero(v interface{}) bool {
	st := reflect.ValueOf(v)
	nonZeroValue := false
	switch st.Kind() {
	case reflect.Ptr, reflect.Interface:
		nonZeroValue = st.IsNil()
	case reflect.Invalid:
		nonZeroValue = true // always invalid
	case reflect.Struct:
		nonZeroValue = false // always valid since only nil pointers are empty
	default:
		return true
	}

	if nonZeroValue {
		return true
	}
	return false
}
