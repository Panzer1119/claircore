// Importing this package registers default matchers via its init function.
package defaults

import (
	"context"
	"sync"
	"time"

	"github.com/Panzer1119/claircore/alpine"
	"github.com/Panzer1119/claircore/aws"
	"github.com/Panzer1119/claircore/debian"
	"github.com/Panzer1119/claircore/libvuln/driver"
	"github.com/Panzer1119/claircore/matchers/registry"
	"github.com/Panzer1119/claircore/oracle"
	"github.com/Panzer1119/claircore/photon"
	"github.com/Panzer1119/claircore/python"
	"github.com/Panzer1119/claircore/rhel"
	"github.com/Panzer1119/claircore/suse"
	"github.com/Panzer1119/claircore/ubuntu"
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
// matchers.
func Error() error {
	return regerr
}

// defaultMatchers is a variable containing
// all the matchers libvuln will use to match
// index records to vulnerabilities.
var defaultMatchers = []driver.Matcher{
	&alpine.Matcher{},
	&aws.Matcher{},
	&debian.Matcher{},
	&oracle.Matcher{},
	&photon.Matcher{},
	&python.Matcher{},
	&rhel.Matcher{},
	&suse.Matcher{},
	&ubuntu.Matcher{},
}

func inner(ctx context.Context) error {
	for _, m := range defaultMatchers {
		mf := driver.MatcherStatic(m)
		registry.Register(m.Name(), mf)
	}
	return nil
}
