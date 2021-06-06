package suse

import (
	"context"
	"testing"
	"time"

	"github.com/quay/zlog"

	"github.com/Panzer1119/claircore/libvuln/driver"
	"github.com/Panzer1119/claircore/test/integration"
)

func TestLiveDatabase(t *testing.T) {
	integration.Skip(t)
	ctx := zlog.Test(context.Background(), t)

	u, err := NewUpdater(EnterpriseServer15)
	if err != nil {
		t.Fatal(err)
	}

	tctx, done := context.WithTimeout(ctx, time.Minute)
	defer done()
	rc, _, err := u.Fetch(tctx, driver.Fingerprint(""))
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	tctx, done = context.WithTimeout(ctx, 8*time.Minute)
	defer done()
	vs, err := u.Parse(tctx, rc)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("found %d definitions", len(vs))
}
