// +build acceptance

package v3

import (
	"testing"

	"github.com/yyf330/gophercloud/acceptance/clients"
	"github.com/yyf330/gophercloud/openstack"
	th "github.com/yyf330/gophercloud/testhelper"

	"github.com/yyf330/gophercloud"
	"github.com/yyf330/gophercloud/openstack/identity/v3/projects"
)

func TestReauthAuthResultDeadlock(t *testing.T) {
	clients.RequireAdmin(t)

	ao, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	ao.AllowReauth = true

	provider, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	provider.SetToken("this is not a valid token")

	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	pages, err := projects.List(client, nil).AllPages()
	th.AssertNoErr(t, err)
	_, err = projects.ExtractProjects(pages)
	th.AssertNoErr(t, err)
}
