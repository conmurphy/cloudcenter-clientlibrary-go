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

### Without helper function

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
### With helper function

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

Two options have been implemented in this library for each async API (AddCloudAccount used as example below):

*  __AddCloudAccountSync__: Client library will make an asynchronous call and wait until the task is complete. Once complete it will return either the newly created object or an error message.
*  __AddCloudAccountAsync__: Client library will make an asynchronous call and will return the operationStatus of the call. The client library user will be required to monitor the operation status and once successful retrieve the newly created object. 

## Reference

- [Actionpolicies](#actionpolicies)
- [Actions](#actions)
- [Activationprofiles](#activationprofiles)
- [Agingpolicies](#agingpolicies)
- [Apps](#apps)
- [Bundles](#bundles)
- [Client](#client)
- [Cloudaccounts](#cloudaccounts)
- [Cloudimagemapping](#cloudimagemapping)
- [Cloudinstancetypes](#cloudinstancetypes)
- [Cloudregions](#cloudregions)
- [Cloudstoragetypes](#cloudstoragetypes)
- [Clouds](#clouds)
- [Contracts](#contracts)
- [Environments](#environments)
- [Groups](#groups)
- [Images](#images)
- [Jobs](#jobs)
- [Operationstatus](#operationstatus)
- [Phases](#phases)
- [Plans](#plans)
- [Projects](#projects)
- [Roles](#roles)
- [Services](#services)
- [Suspensionpolicies](#suspensionpolicies)
- [Tenants](#tenants)
- [Users](#users)
- [Virtualmachines](#virtualmachines)

### ActionPolicies

```go
type ActionPolicyAPIResponse struct {
	Resource       *string        
	Size           *int64         
	PageNumber     *int64         
	TotalElements  *int64         
	TotalPages     *int64         
	ActionPolicies []ActionPolicy 
}
```

```go
type ActionPolicy struct {
	Id          *string    
	Resource    *string    
	Perms       *[]string  
	Name        *string    
	Description *string    
	EntityType  *string    
	EventName   *string    
	Actions     *[]Actions 
	UserId      *string    
	Enabled     *bool      
	AutoEnable  *bool      
	ForceEnable *bool      
	Global      *bool      
}
```
#### GetActionPolicies

```go
func (s *Client) GetActionPolicies() ([]ActionPolicy, error)
```

#####Example

```go
actionPolicies, err := client.GetActionPolicies()

if err != nil {
	fmt.Println(err)
} else {
	for _, actionPolicy := range actionPolicies {

		fmt.Println("Id: " + actionPolicy.Id + ", Name: " + actionPolicy.Name)

	}
}
```

#### GetActionPolicy

```go
func (s *Client) GetActionPolicy(actionPolicyId int) (*ActionPolicy, error)
```

#####Example

```go
actionPolicy, err := client.GetActionPolicy(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Id: " + actionPolicy.Id + ", Name: " + actionPolicy.Name)
}
```

### Actions

```go
type ActionAPIResponse struct {
	Resource      *string  
	Size          *int64   
	PageNumber    *int64   
	TotalElements *int64   
	TotalPages    *int64   
	ActionJaxbs   []Action 
}
```

```go
type Action struct {
	Id                     *string                  
	Resource               *string                  
	Perms                  *[]string                
	Name                   *string                  
	Description            *string                  
	ActionType             *string                  
	LastUpdatedTime        *string                  
	TimeOut                *float64                 
	Enabled                *bool                    
	Encrypted              *bool                    
	Deleted                *bool                    
	SystemDefined          *bool                    
	BulkOperationSupported *bool                    
	IsAvailableToUser      *bool                    
	Owner                  *int64                   
	ActionParameters       *[]ActionParameter       
	ActionResourceMappings *[]ActionResourceMapping 
	ActionCustomParamSpecs *[]ActionCustomParamSpec 
}
```

```go
type ActionParameter struct {
	ParamName   *string 
	ParamValue  *string 
	CustomParam *bool   
	Required    *bool   
	Preference  *string 
}
```

```go
type ActionResourceMapping struct {
	Type                  *string                 
	ActionResourceFilters *[]ActionResourceFilter 
}
```

```go
type ActionCustomParamSpec struct {
	ParamName            *string              
	DisplayName          *string              
	HelpText             *string              
	Type                 *string              
	ValueList            *string              
	DefaultValue         *string              
	ConfirmValue         *string              
	PathSuffixValue      *string              
	UserVisible          *bool                
	UserEditable         *bool                
	SystemParam          *bool                
	ExampleValue         *string              
	DataUnit             *string              
	Optional             *bool                
	MultiselectSupported *bool                
	ValueConstraint      *ValueConstraint     
	Scope                *string              
	WebserviceListParams *WebserviceListParam 
	Preference           *string              
}
```

#### GetActions

```go
func (s *Client) GetActions() ([]Action, error)
```

#####Example

```go
actions, err := client.GetActions()

if err != nil {
	fmt.Println(err)
} else {
	for _, action := range actions {

		fmt.Println("Id: " + action.Id + ", Name: " + action.Name)

	}
}
```

#### GetAction

```go
func (s *Client) GetAction(id int) (*Action, error)
```

#####Example

```go
action, err := client.GetAction(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Id: " + action.Id + ", Name: " + action.Name)
}
```

### ActivationProfiles

```go
type ActivationProfile struct {
	Id                  *string           
	Name                *string           
	Description         *string           
	Resource            *string           
	TenantId            *int64            
	PlanId              *string           
	BundleId            *string           
	ContractId          *string           
	DepEnvId            *string           
	ActivateRegions     *[]ActivateRegion 
	ImportApps          *[]string         
	AgreeToContract     *bool             
	SendActivationEmail *bool             
}
```

```go
type ActivateRegion struct {
	RegionId string 
}
```

#### GetActivationProfiles

```go
func (s *Client) GetActivationProfiles(tenantId int) ([]ActivationProfile, error)
```

#####Example

```go
activationProfiles, err := client.GetActivationProfiles(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, activationProfile := range activationProfiles {

		fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)

	}
}
```

#### GetActivationProfile

```go
func (s *Client) GetActivationProfile(tenantId int, activationProfileId int) (*ActivationProfile, error)
```

#####Example

```go
activationProfile, err := client.GetActivationProfile(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)
}
```
### AgingPolicies
### Apps
### Bundles
### Client
### CloudAccounts
### CloudImageMapping
### CloudInstanceTypes
### CloudRegions
### CloudStorageTypes
### Clouds
### Contracts
### Environments
### Groups
### Images
### Jobs
### OperationStatus
### Phases
### Plans
### Projects
### Roles
### Services
### SuspensionPolicies
### Tenants
### Users

- [GetUsers](#getusers)
- [GetUser](#getuser)
- [GetUserByEmail](#getuserbyemail)
- [AddUser](#adduser)
- [UpdateUser](#updateuser)
- [DeleteUser](#deleteuser)

```go
type UserAPIResponse struct {
	Resource      string 
	Size          int    
	PageNumber    int    
	TotalElements int    
	TotalPages    int    
	Users         []User 
}
```

```go
type User struct {
	Id                      *string 
	Resource                *string 
	Username                *string  
	Password                *string  
	Enabled                 *bool   
	Type                    *string 
	FirstName               *string 
	LastName                *string 
	CompanyName             *string 
	EmailAddr               *string 
	EmailVerified           *bool   
	PhoneNumber             *string 
	ExternalId              *string 
	AccessKeys              *string 
	DisableReason           *string 
	AccountSource           *string 
	Status                  *string 
	Detail                  *string 
	ActivationData          *string 
	Created                 *int64  
	LastUpdated             *int64  
	CoAdmin                 *bool   
	TenantAdmin             *bool   
	ActivationProfileId     *string 
	HasSubscriptionPlanType *bool   
	TenantId                *string 
}
```

#### AddUser

```go
func (s *Client) AddUser(user *User) (*User, error)
```

##### __Required Fields__
* EmailAddr
* TenantId
* ActivateRegions (When creating user with activation)
* ActivationProfileId (required if using activation profile method)

##### Example

```golang

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

#### DeleteUser

```go
func (s *Client) DeleteUser(userId int) error
```
##### Example
```go
fmt.Println()
/*
	Delete user
*/

err := client.DeleteUser(6)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("User deleted")
}
```
### VirtualMachines


WARNING:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
