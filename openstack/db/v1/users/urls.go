package users

import "github.com/yyf330/gophercloud"

func baseURL(c *gophercloud.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "users")
}

func userURL(c *gophercloud.ServiceClient, instanceID, userName string) string {
	return c.ServiceURL("instances", instanceID, "users", userName)
}
