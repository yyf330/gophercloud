package ruletypes

import "github.com/yyf330/gophercloud"

func listRuleTypesURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}

func getRuleTypeURL(c *gophercloud.ServiceClient, name string) string {
	return c.ServiceURL("qos", "rule-types", name)
}
