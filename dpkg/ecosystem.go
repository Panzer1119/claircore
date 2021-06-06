package dpkg

import (
	"context"

	"github.com/Panzer1119/claircore/debian"
	"github.com/Panzer1119/claircore/internal/indexer"
	"github.com/Panzer1119/claircore/internal/indexer/linux"
	"github.com/Panzer1119/claircore/ubuntu"
)

// NewEcosystem provides the set of scanners and coalescers for the dpkg ecosystem
func NewEcosystem(ctx context.Context) *indexer.Ecosystem {
	return &indexer.Ecosystem{
		PackageScanners: func(ctx context.Context) ([]indexer.PackageScanner, error) {
			return []indexer.PackageScanner{&Scanner{}}, nil
		},
		DistributionScanners: func(ctx context.Context) ([]indexer.DistributionScanner, error) {
			return []indexer.DistributionScanner{
				&debian.DistributionScanner{},
				&ubuntu.DistributionScanner{},
			}, nil
		},
		RepositoryScanners: func(ctx context.Context) ([]indexer.RepositoryScanner, error) {
			return []indexer.RepositoryScanner{}, nil
		},
		Coalescer: func(ctx context.Context) (indexer.Coalescer, error) {
			return linux.NewCoalescer(), nil
		},
	}
}
