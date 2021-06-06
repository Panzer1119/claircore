package debian

import (
	"context"

	"github.com/Panzer1119/claircore/libvuln/driver"
)

var debianReleases = []Release{
	Buster,
	Jessie,
	Stretch,
	Wheezy,
}

func UpdaterSet(_ context.Context) (driver.UpdaterSet, error) {
	us := driver.NewUpdaterSet()
	for _, release := range debianReleases {
		u := NewUpdater(release)
		err := us.Add(u)
		if err != nil {
			return us, err
		}
	}
	return us, nil
}
