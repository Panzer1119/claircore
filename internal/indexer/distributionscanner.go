package indexer

import (
	"context"

	"github.com/Panzer1119/claircore"
)

type DistributionScanner interface {
	VersionedScanner
	Scan(context.Context, *claircore.Layer) ([]*claircore.Distribution, error)
}
