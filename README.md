# Cloudcenter Go Client Library

This is a Go Client Library used for accessing Cisco CloudCenter. 

It is currently a Proof of Concept and has been developed and tested against Cisco CloudCenter 4.8.2 with Go version 1.9.3

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/overview.png)

## Quick Start

```golang
package main
import "github.com/cloudcenter-clientlibrary-go/cloudcenter”

/*
	Define new cloudcenter client
*/

client := cloudcenter.NewClient("cliqradmin", ”myAPIKey", "https://ccm.dcloud.cisco.com")

/*
	Create user
*/

newUser := cloudcenter.User{
	TenantId:    "1",
	FirstName:   "client",
	LastName:    "library",
	Password:    "myPassword",
	EmailAddr:   "clientlibrary@cloudcenter.com",
	CompanyName: "Company",
	PhoneNumber: "12345",
	ExternalId:  "23456",
}

user, err := client.AddUser(&newUser)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(”New user created. \n UserId: " + user.Id + ", User last name: " + user.LastName)
}
```
