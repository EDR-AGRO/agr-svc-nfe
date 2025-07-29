package apm

import (
	"context"
	"net/http"

	"go.elastic.co/apm/v2"
)

type ApmTransaction struct {
	Ctx      context.Context
	apmTx    *apm.Transaction
	name     string
	spanType SpanType
}

type ApmSpan struct {
	ApmSpan  *apm.Span
	name     string
	spanType SpanType
}

type SpanType string

const (
	ORACLE_DATABASE     SpanType = "ORACLE_DATABASE"
	SQL_SERVER_DATABASE SpanType = "SQL_SERVER_DATABASE"
	MONGO_DB            SpanType = "MONGO_DB"
	REDIS_DATABASE      SpanType = "REDIS_DATABASE"
	KAFKA_CONSUMER      SpanType = "KAFKA_CONSUMER"
	KAFKA_PRODUCER      SpanType = "KAFKA_PRODUCER"
	HTTP_REQUEST        SpanType = "HTTP_REQUEST"
)

func SetDatabase(instance, stmt, tp, user string) apm.DatabaseSpanContext {
	return apm.DatabaseSpanContext{
		Instance:  instance,
		Statement: stmt,
		Type:      tp,
		User:      user,
	}
}

func SetDestination(name, resource string) apm.DestinationServiceSpanContext {
	return apm.DestinationServiceSpanContext{
		Name:     name,
		Resource: resource,
	}
}

func StartTransaction(ctx context.Context, name string, spantype SpanType) *ApmTransaction {
	t := &ApmTransaction{
		name:     name,
		spanType: spantype,
	}
	tx := apm.DefaultTracer().StartTransaction(name, string(spantype))
	t.Ctx = apm.ContextWithTransaction(ctx, tx)
	t.apmTx = tx

	return t
}

func (t *ApmTransaction) EndTransaction() {
	t.apmTx.End()
}

func StartHttpTransaction(r *http.Request, name string) *ApmTransaction {
	ctx := r.Context()
	t := &ApmTransaction{
		Ctx:      ctx,
		name:     name,
		spanType: HTTP_REQUEST,
	}
	tx := apm.TransactionFromContext(ctx)
	t.apmTx = tx
	return t
}

func (t *ApmTransaction) StartSpan(name string, spantype SpanType) *ApmSpan {
	span := t.apmTx.StartSpan(name, string(spantype), nil)
	return &ApmSpan{
		name:     name,
		spanType: spantype,
		ApmSpan:  span,
	}
}

func (t *ApmSpan) EndSpan() {
	t.ApmSpan.End()
}

func Flush() {
	apm.DefaultTracer().Flush(nil)
}
