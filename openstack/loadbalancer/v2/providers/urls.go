package providers

import "github.com/yyf330/gophercloud"

const (
	rootPath     = "lbaas"
	resourcePath = "providers"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}
