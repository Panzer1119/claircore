package rpm

import (
	"context"

	"github.com/Panzer1119/claircore/aws"
	"github.com/Panzer1119/claircore/internal/indexer"
	"github.com/Panzer1119/claircore/internal/indexer/linux"
	"github.com/Panzer1119/claircore/oracle"
	"github.com/Panzer1119/claircore/photon"
	"github.com/Panzer1119/claircore/suse"
)

// NewEcosystem provides the set of scanners and coalescers for the rpm ecosystem
func NewEcosystem(ctx context.Context) *indexer.Ecosystem {
	return &indexer.Ecosystem{
		PackageScanners: func(ctx context.Context) ([]indexer.PackageScanner, error) {
			return []indexer.PackageScanner{&Scanner{}}, nil
		},
		DistributionScanners: func(ctx context.Context) ([]indexer.DistributionScanner, error) {
			return []indexer.DistributionScanner{
				&aws.DistributionScanner{},
				&oracle.DistributionScanner{},
				&suse.DistributionScanner{},
				&photon.DistributionScanner{},
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
