package swauth

import "github.com/yyf330/gophercloud"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
