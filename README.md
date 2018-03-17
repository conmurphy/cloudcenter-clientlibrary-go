# Cloudcenter Go Client Library

This is a Go Client Library used for accessing Cisco CloudCenter. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco CloudCenter 4.8.2 with Go version 1.9.3

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/overview.png)

## Quick Start

```golang
package main
import "github.com/cloudcenter-clientlibrary-go/cloudcenter”

/*
	Define new cloudcenter client
*/

client := cloudcenter.NewClient("cliqradmin", ”myAPIKey", "https://ccm.cloudcenter-address.com")

/*
	Create user
*/

newUser := cloudcenter.User {
	FirstName:   cloudcenter.String("client"),
	LastName:    cloudcenter.String("library"),
	Password:    cloudcenter.String("myPassword"),
	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
	CompanyName: cloudcenter.String("company"),
	PhoneNumber: cloudcenter.String("12345"),
	ExternalId:  cloudcenter.String("23456"),
	TenantId:    cloudcenter.String("1"),
}

user, err := client.AddUser(&newUser)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(”New user created. \n UserId: " + user.Id + ", User last name: " + user.LastName)
}
```

## Helper Functions

As per the following link, using the Marshal function from the encoding/json library treats false booleans as if they were nil values, and thus it omits them from the JSON response. To make a distinction between a non-existent boolean and false boolean we need to use a ```*bool``` in the struct. 

```golang
type User struct {
	Id                      *string `json:"id,omitempty"`
	FirstName               *string `json:"firstName,omitempty"`
	LastName                *string `json:"lastName,omitempty"`
	Password                *string `json:"password,omitempty"` 
	EmailAddr               *string `json:"emailAddr,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty"`
	TenantAdmin             *bool   `json:"tenantAdmin,omitempty"`
}
```
https://github.com/golang/go/issues/13284

Therefore in order to have a consistent experience all struct fields within this client library use pointers. This provides a way to differentiate between unset values, nil, and an intentional zero value, such as "", false, or 0. 

Helper functions have been created to simplify the creation of pointer types.

### Without 

```golang
firstName 	:= "client"
lastName 	:= "library"
password	:= "myPassword"
emailAddr	:= "clientlibrary@cloudcenter-address.com"
companyName	:= "company"
phoneNumber	:= "12345"
externalId	:= "23456"
tenantId	:= "1"


newUser := cloudcenter.User {
	FirstName:   &firstName,
	LastName:    &lastName,
	Password:    &password,
	EmailAddr:  &emailAddr,
	CompanyName: &companyName,
	PhoneNumber: &phoneNumber,
	ExternalId: &externalId,
	TenantId:    &tenantId,
}
```
### With

```golang
newUser := cloudcenter.User {
	FirstName:   cloudcenter.String("client"),
	LastName:    cloudcenter.String("library"),
	Password:    cloudcenter.String("myPassword"),
	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
	CompanyName: cloudcenter.String("company"),
	PhoneNumber: cloudcenter.String("12345"),
	ExternalId:  cloudcenter.String("23456"),
	TenantId:    cloudcenter.String("1"),
}
```

Reference: https://willnorris.com/2014/05/go-rest-apis-and-pointers

### Available Helper Functions

* cloudcenter.Bool()
* cloudcenter.Int()
* cloudcenter.Int64()
* cloudcenter.String()
* cloudcenter.Float32()
* cloudcenter.Float64()

## Sync and Async

* Synchronous APIs indicate that the program execution waits for a response to be returned by the API. The execution does not proceed until the call is completed. The real state of the API request is available in the response.

* Asynchronous APIs do not wait for the API call to complete. Program execution continues, and until the call completes,  you can issue GET requests to review the state after the submission, during the execution, and after the call completion

https://editor-docs.cloudcenter.cisco.com/display/40API/Synchronous+and+Asynchronous+APIs

Two options have been implemented in this library for each async API (example resource):

*  __AddCloudAccountSync__: Client library will make an asynchronous call and wait until the task is complete. Once complete it will return either the newly created object or an error message.
*  __AddCloudAccountAsync__: Client library will make an asynchronous call and will return the operationStatus of the call. The client library user will be required to monitor the operation status and once successful retrieve the newly created object. 

## Reference

* [ActionPolicies](ActionPolicies)
* Actions
* ActivationProfiles
* AgingPolicies
* Apps
* Bundles
* Client
* CloudAccounts
* CloudImageMapping
* CloudInstanceTypes
* CloudRegions
* CloudStorageTypes
* Clouds
* Contracts
* Environments
* Groups
* Images
* Jobs
* OperationStatus
* Phases
* Plans
* Projects
* Roles
* Services
* SuspensionPolicies
* Tenants
* Users
* VirtualMachines

### ActionPolicies

WARNING:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
