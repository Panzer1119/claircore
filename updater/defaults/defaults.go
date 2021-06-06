// Package defaults sets updater defaults.
//
// Importing this package registers default updaters via its init function.
package defaults

import (
	"context"
	"sync"
	"time"

	"github.com/Panzer1119/claircore/alpine"
	"github.com/Panzer1119/claircore/aws"
	"github.com/Panzer1119/claircore/debian"
	"github.com/Panzer1119/claircore/libvuln/driver"
	"github.com/Panzer1119/claircore/oracle"
	"github.com/Panzer1119/claircore/photon"
	"github.com/Panzer1119/claircore/pyupio"
	"github.com/Panzer1119/claircore/rhel"
	"github.com/Panzer1119/claircore/suse"
	"github.com/Panzer1119/claircore/ubuntu"
	"github.com/Panzer1119/claircore/updater"
)

var (
	once   sync.Once
	regerr error
)

func init() {
	ctx, done := context.WithTimeout(context.Background(), 1*time.Minute)
	defer done()
	once.Do(func() { regerr = inner(ctx) })
}

// Error reports if an error was encountered when initializing the default
// updaters.
func Error() error {
	return regerr
}

func inner(ctx context.Context) error {
	rf, err := rhel.NewFactory(ctx, rhel.DefaultManifest)
	if err != nil {
		return err
	}
	updater.Register("rhel", rf)

	updater.Register("ubuntu", &ubuntu.Factory{Releases: ubuntu.Releases})
	updater.Register("alpine", driver.UpdaterSetFactoryFunc(alpine.UpdaterSet))
	updater.Register("aws", driver.UpdaterSetFactoryFunc(aws.UpdaterSet))
	updater.Register("debian", driver.UpdaterSetFactoryFunc(debian.UpdaterSet))
	updater.Register("oracle", driver.UpdaterSetFactoryFunc(oracle.UpdaterSet))
	updater.Register("photon", driver.UpdaterSetFactoryFunc(photon.UpdaterSet))
	updater.Register("pyupio", driver.UpdaterSetFactoryFunc(pyupio.UpdaterSet))
	updater.Register("suse", driver.UpdaterSetFactoryFunc(suse.UpdaterSet))

	return nil
}
