package testing

import (
	"testing"
	"time"

	"github.com/yyf330/gophercloud/openstack/objectstorage/v1/accounts"
	th "github.com/yyf330/gophercloud/testhelper"
	fake "github.com/yyf330/gophercloud/testhelper/client"
)

func TestUpdateAccount(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateAccountSuccessfully(t)

	options := &accounts.UpdateOpts{Metadata: map[string]string{"gophercloud-test": "accounts"}}
	res := accounts.Update(fake.ServiceClient(), options)
	th.AssertNoErr(t, res.Err)

	expected := &accounts.UpdateHeader{
		Date: time.Date(2014, time.January, 17, 16, 9, 56, 0, time.UTC),
	}
	actual, err := res.Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetAccount(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetAccountSuccessfully(t)

	expectedMetadata := map[string]string{"Subject": "books", "Quota-Bytes": "42"}
	res := accounts.Get(fake.ServiceClient(), &accounts.GetOpts{})
	th.AssertNoErr(t, res.Err)
	actualMetadata, _ := res.ExtractMetadata()
	th.CheckDeepEquals(t, expectedMetadata, actualMetadata)
	_, err := res.Extract()
	th.AssertNoErr(t, err)

	var quotaBytes int64 = 42
	expected := &accounts.GetHeader{
		QuotaBytes:     &quotaBytes,
		ContainerCount: 2,
		ObjectCount:    5,
		BytesUsed:      14,
		Date:           time.Date(2014, time.January, 17, 16, 9, 56, 0, time.UTC),
	}
	actual, err := res.Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}

func TestGetAccountNoQuota(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetAccountNoQuotaSuccessfully(t)

	expectedMetadata := map[string]string{"Subject": "books"}
	res := accounts.Get(fake.ServiceClient(), &accounts.GetOpts{})
	th.AssertNoErr(t, res.Err)
	actualMetadata, _ := res.ExtractMetadata()
	th.CheckDeepEquals(t, expectedMetadata, actualMetadata)
	_, err := res.Extract()
	th.AssertNoErr(t, err)

	expected := &accounts.GetHeader{
		QuotaBytes:     nil,
		ContainerCount: 2,
		ObjectCount:    5,
		BytesUsed:      14,
		Date:           time.Date(2014, time.January, 17, 16, 9, 56, 0, time.UTC),
	}
	actual, err := res.Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}
