// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/yyf330/gophercloud/acceptance/clients"
	"github.com/yyf330/gophercloud/acceptance/tools"
	"github.com/yyf330/gophercloud/openstack/blockstorage/extensions/schedulerstats"
	th "github.com/yyf330/gophercloud/testhelper"
)

func TestSchedulerStatsList(t *testing.T) {
	clients.RequireAdmin(t)
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")

	blockClient, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	listOpts := schedulerstats.ListOpts{
		Detail: true,
	}

	allPages, err := schedulerstats.List(blockClient, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allStats, err := schedulerstats.ExtractStoragePools(allPages)
	th.AssertNoErr(t, err)

	for _, stat := range allStats {
		tools.PrintResource(t, stat)
	}
}
