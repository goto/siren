package telemetry

import (
	"context"
	"fmt"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/xo/dburl"
	"go.opencensus.io/trace"
)

type PostgresTracer struct {
	dbSystem           string
	dbName             string
	dbUser             string
	dbAddr             string
	dbPort             string
	nrDataStoreSegment newrelic.DatastoreSegment
	span               *trace.Span
}

func NewPostgresTracer(url string) (*PostgresTracer, error) {
	u, err := dburl.Parse(url)
	if err != nil {
		return nil, err
	}
	return &PostgresTracer{
		dbSystem: "postgresql",
		dbName:   strings.TrimPrefix(u.EscapedPath(), "/"),
		dbUser:   u.User.Username(),
		dbAddr:   u.Hostname(),
		dbPort:   u.Port(),
	}, err
}

func (d *PostgresTracer) StartSpan(ctx context.Context, op string, tableName string, query string, spanAttributes ...trace.Attribute) (context.Context, *trace.Span) {
	nrTx := newrelic.FromContext(ctx)
	d.nrDataStoreSegment = newrelic.DatastoreSegment{
		Product:            newrelic.DatastorePostgres,
		DatabaseName:       d.dbName,
		PortPathOrID:       d.dbPort,
		Collection:         tableName,
		Operation:          op,
		Host:               d.dbAddr,
		StartTime:          nrTx.StartSegmentNow(),
		ParameterizedQuery: query,
	}

	// Refer https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/database.md
	traceCtx, span := trace.StartSpan(ctx, fmt.Sprintf("%s %s.%s", op, d.dbName, tableName), trace.WithSpanKind(trace.SpanKindClient))

	traceAttributes := []trace.Attribute{
		trace.StringAttribute("db.system", d.dbSystem),
		trace.StringAttribute("db.statement", query),
		trace.StringAttribute("db.user", d.dbUser),
		trace.StringAttribute("net.sock.peer.addr", d.dbAddr),
		trace.StringAttribute("net.peer.port", d.dbPort),
		trace.StringAttribute("db.name", d.dbName),
		trace.StringAttribute("db.operation", op),
		trace.StringAttribute("db.sql.table", tableName),
	}

	traceAttributes = append(traceAttributes, spanAttributes...)

	span.AddAttributes(
		traceAttributes...,
	)

	d.span = span

	return traceCtx, span
}

func (d *PostgresTracer) StopSpan() {
	d.nrDataStoreSegment.End()
	d.span.End()
}
