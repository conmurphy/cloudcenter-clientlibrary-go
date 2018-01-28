package main

import "cloudcenterclient"
import "fmt"

func main() {
	client := cloudcenterclient.NewClient("<Username>", "<API Key>", "<CCM URL>")

	users, _ := client.GetUsers()

	for _, user := range users {

		fmt.Println(user.Username)

	}

	tenants, _ := client.GetTenants()

	for _, tenant := range tenants {

		fmt.Println(tenant.Name)

	}

}
