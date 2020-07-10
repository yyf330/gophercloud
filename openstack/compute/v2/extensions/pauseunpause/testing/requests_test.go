package testing

import (
	"testing"

	"github.com/yyf330/gophercloud/openstack/compute/v2/extensions/pauseunpause"
	th "github.com/yyf330/gophercloud/testhelper"
	"github.com/yyf330/gophercloud/testhelper/client"
)

const serverID = "{serverId}"

func TestPause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockPauseServerResponse(t, serverID)

	err := pauseunpause.Pause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnpause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockUnpauseServerResponse(t, serverID)

	err := pauseunpause.Unpause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
