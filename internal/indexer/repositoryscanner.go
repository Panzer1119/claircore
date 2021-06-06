package indexer

import (
	"context"

	"github.com/Panzer1119/claircore"
)

type RepositoryScanner interface {
	VersionedScanner
	Scan(context.Context, *claircore.Layer) ([]*claircore.Repository, error)
}
