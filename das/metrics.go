package das

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncfloat64"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/celestiaorg/celestia-node/header"
)

var (
	meter = global.MeterProvider().Meter("das")
)

type metrics struct {
	sampled       syncint64.Counter
	sampleTime    syncfloat64.Histogram
	getHeaderTime syncfloat64.Histogram
	newHead       syncint64.Counter
	lastSampledTS int64
}

func (sc *samplingCoordinator) initMetrics() error {
	sampled, err := meter.SyncInt64().Counter("das_sampled_headers_counter",
		instrument.WithDescription("sampled headers counter"))
	if err != nil {
		return err
	}

	sampleTime, err := meter.SyncFloat64().Histogram("das_sample_time_hist",
		instrument.WithDescription("duration of sampling a single header"))
	if err != nil {
		return err
	}

	getHeaderTime, err := meter.SyncFloat64().Histogram("das_get_header_time_hist",
		instrument.WithDescription("duration of getting header from header store"))
	if err != nil {
		return err
	}

	newHead, err := meter.SyncInt64().Counter("das_head_updated_counter",
		instrument.WithDescription("amount of times DAS'er advanced network head"))
	if err != nil {
		return err
	}

	lastSampledTS, err := meter.AsyncInt64().Gauge("das_latest_sampled_ts",
		instrument.WithDescription("latest sampled timestamp"))
	if err != nil {
		return err
	}

	busyWorkers, err := meter.AsyncInt64().Gauge("das_busy_workers_amount",
		instrument.WithDescription("number of active parallel workers in DAS'er"))
	if err != nil {
		return err
	}

	networkHead, err := meter.AsyncInt64().Gauge("das_network_head",
		instrument.WithDescription("most recent network head"))
	if err != nil {
		return err
	}

	sampledChainHead, err := meter.AsyncInt64().Gauge("das_sampled_chain_head",
		instrument.WithDescription("height of the sampled chain - all previous headers have been successfully sampled"))
	if err != nil {
		return err
	}

	sc.metrics = &metrics{
		sampled:       sampled,
		sampleTime:    sampleTime,
		getHeaderTime: getHeaderTime,
		newHead:       newHead,
	}

	err = meter.RegisterCallback(
		[]instrument.Asynchronous{
			lastSampledTS, busyWorkers, networkHead, sampledChainHead,
		},
		func(ctx context.Context) {
			stats, err := sc.stats(ctx)
			if err != nil {
				log.Errorf("observing stats: %s", err.Error())
			}

			busyWorkers.Observe(ctx, int64(len(stats.Workers)))
			networkHead.Observe(ctx, int64(stats.NetworkHead))
			sampledChainHead.Observe(ctx, int64(stats.SampledChainHead))

			if ts := atomic.LoadInt64(&sc.metrics.lastSampledTS); ts != 0 {
				lastSampledTS.Observe(ctx, ts)
			}
		},
	)

	if err != nil {
		return fmt.Errorf("regestering metrics callback: %w", err)
	}

	return nil
}

func (m *metrics) observeSample(ctx context.Context, h *header.ExtendedHeader, sampleTime time.Duration, err error) {
	if m == nil {
		return
	}
	m.sampleTime.Record(ctx, sampleTime.Seconds(),
		attribute.Bool("failed", err != nil),
		attribute.Int("header_width", len(h.DAH.RowsRoots)))
	m.sampled.Add(ctx, 1,
		attribute.Bool("failed", err != nil),
		attribute.Int("header_width", len(h.DAH.RowsRoots)))
	atomic.StoreInt64(&m.lastSampledTS, time.Now().UTC().Unix())
}

func (m *metrics) observeGetHeader(ctx context.Context, d time.Duration) {
	if m == nil {
		return
	}
	m.getHeaderTime.Record(ctx, d.Seconds())
}

func (m *metrics) observeNewHead(ctx context.Context) {
	if m == nil {
		return
	}
	m.newHead.Add(ctx, 1)
}
