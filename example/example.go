package main

import "cloudcenter"
import "fmt"

func main() {

	/*
		Define new cloudcenter client
	*/

	client := cloudcenter.NewClient("<Username>", "<API Key>", "<https://example_cloudcenter_host>")

	/*
		Retrieve all users from CloudCenter
	*/

	users, err := client.GetUsers()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, user := range users {

			fmt.Println(user.Username)

		}
	}

	/*
		Retrieve all details about a user from CloudCenter
	*/

	user, err := client.GetUser(2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	/*
		Retrieve all tenants from CloudCenter
	*/

	tenants, err := client.GetTenants()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, tenant := range tenants {

			fmt.Println(tenant.Name)

		}
	}

	/*
		Retrieve all jobs from CloudCenter
	*/

	jobs, err := client.GetJobs()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, job := range jobs {

			fmt.Println(job.Name)

		}
	}

	/*
		Retrieve all application profiles from CloudCenter
	*/

	apps, err := client.GetApps()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, app := range apps {

			fmt.Println(app.Name)

		}
	}

	/*
		Retrieve all clouds from CloudCenter
	*/

	clouds, err := client.GetClouds()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, cloud := range clouds {

			fmt.Println(cloud.Name)

		}
	}

}
