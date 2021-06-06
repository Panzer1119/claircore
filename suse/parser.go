package suse

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/quay/goval-parser/oval"
	"github.com/quay/zlog"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/label"

	"github.com/Panzer1119/claircore"
	"github.com/Panzer1119/claircore/libvuln/driver"
	"github.com/Panzer1119/claircore/pkg/ovalutil"
)

var _ driver.Parser = (*Updater)(nil)

func (u *Updater) Parse(ctx context.Context, r io.ReadCloser) ([]*claircore.Vulnerability, error) {
	ctx = baggage.ContextWithValues(ctx,
		label.String("component", "suse/Updater.Parse"))
	zlog.Info(ctx).Msg("starting parse")
	defer r.Close()
	root := oval.Root{}
	if err := xml.NewDecoder(r).Decode(&root); err != nil {
		return nil, fmt.Errorf("suse: unable to decode OVAL document: %w", err)
	}
	zlog.Debug(ctx).Msg("xml decoded")
	protoVulns := func(def oval.Definition) ([]*claircore.Vulnerability, error) {
		return []*claircore.Vulnerability{
			&claircore.Vulnerability{
				Updater:            u.Name(),
				Name:               def.Title,
				Description:        def.Description,
				Links:              ovalutil.Links(def),
				Severity:           def.Advisory.Severity,
				NormalizedSeverity: NormalizeSeverity(def.Advisory.Severity),
				// each updater is configured to parse a suse release
				// specific xml database. we'll use the updater's release
				// to map the parsed vulnerabilities
				Dist: releaseToDist(u.release),
			}}, nil
	}
	vulns, err := ovalutil.RPMDefsToVulns(ctx, &root, protoVulns)
	if err != nil {
		return nil, err
	}
	return vulns, nil
}
