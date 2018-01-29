package main

import "cloudcenterclient"
import "fmt"

func main() {
	client := cloudcenterclient.NewClient("<Username>", "<API Key>", "<https://example_cloudcenter_host>")

	users, err := client.GetUsers()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, user := range users {

			fmt.Println(user.Username)

		}
	}

	tenants, err := client.GetTenants()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, tenant := range tenants {

			fmt.Println(tenant.Name)

		}
	}

	jobs, err := client.GetJobs()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, job := range jobs {

			fmt.Println(job.Name)

		}
	}

}
