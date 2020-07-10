package acls

import "github.com/yyf330/gophercloud"

func containerURL(client *gophercloud.ServiceClient, containerID string) string {
	return client.ServiceURL("containers", containerID, "acl")
}

func secretURL(client *gophercloud.ServiceClient, secretID string) string {
	return client.ServiceURL("secrets", secretID, "acl")
}
