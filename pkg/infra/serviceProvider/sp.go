package serviceprovider

import (
	"rabbitmqtest/pkg/domain"
	"rabbitmqtest/pkg/infra/db/repos"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/betalixt/gingorr"
	"go.uber.org/zap"
)

// Handles transient and scoped dependencies
// Should be created for each incoming request
type ServiceProvider struct {
	// trace info
	tid string
	pid string
	rid string
	flg string
	// core
	traceCore *trace.AppInsightsCore
	// derived
	lgr    *zap.Logger
	tracer *trace.AppInsightsTrace
}

var _ domain.IServiceProvider = (*ServiceProvider)(nil)
var _ gingorr.IServiceProvider = (*ServiceProvider)(nil)

// Scoped services
func (prov *ServiceProvider) GetLogger() *zap.Logger {
	return prov.lgr
}

func (prov *ServiceProvider) GetTracer() *trace.AppInsightsTrace {
	if prov.tracer == nil {
		prov.tracer = trace.NewAppInsightsTrace(
			prov.traceCore,
			prov.tid,
			prov.pid,
			prov.rid,
		)
	}
	return prov.tracer
}

// Transient services
func (prov *ServiceProvider) GetRabbitMqRepo() domain.IRabbitMqRepository {
	return repos.NewRabbitMqRepo()
}

// Construcutor
func NewServiceProvider(
	ver string,
	tid string,
	pid string,
	rid string,
	lgr *zap.Logger,
	trCore *trace.AppInsightsCore,
) *ServiceProvider {
	return &ServiceProvider{
		tid: tid,
		pid: pid,
		rid: rid,
		flg: "00",
		lgr: lgr.With(
			zap.String("tid", tid),
			zap.String("pid", pid),
			zap.String("rid", rid),
		),
		traceCore: trCore,
	}
}
