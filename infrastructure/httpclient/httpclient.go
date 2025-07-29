package httpclient

import (
	"br/com/agr/nfe/infrastructure/logger"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"github.com/failsafe-go/failsafe-go/fallback"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
	"go.elastic.co/apm/module/apmhttp/v2"
)

type AppResponse struct {
	Success bool       `json:"success"`
	Result  any        `json:"result"`
	Errors  []AppError `json:"errors"`
}

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func AppErr(code int, errMsg string) AppResponse {
	return AppResponse{
		Success: false,
		Result:  nil,
		Errors: []AppError{
			{
				Code:    code,
				Message: errMsg,
			},
		},
	}
}

func (ar *AppResponse) StringfyError() string {
	str := fmt.Sprintf("Success: %v", ar.Success)
	for _, err := range ar.Errors {
		str += fmt.Sprintf(" - Message: %s - Code: %d", err.Message, err.Code)
	}
	return str
}

func HttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 4096
	t.MaxIdleConnsPerHost = 4096
	t.TLSHandshakeTimeout = 0 * time.Second
	t.DisableKeepAlives = false
	t.IdleConnTimeout = 300 * time.Second
	client := apmhttp.WrapClient(&http.Client{Timeout: time.Duration(300) * time.Second,
		Transport: t,
	})

	return client
}

func HttpClientFailSafe(retryPolicy retrypolicy.RetryPolicy[*http.Response], fallBackPolicy fallback.Fallback[*http.Response]) *http.Client {
	var roundTripper http.RoundTripper
	if fallBackPolicy != nil {
		roundTripper = failsafehttp.NewRoundTripper(nil, fallBackPolicy, retryPolicy)
	} else {
		roundTripper = failsafehttp.NewRoundTripper(nil, retryPolicy)
	}

	client := apmhttp.WrapClient(&http.Client{Timeout: time.Duration(300) * time.Second,
		Transport: roundTripper,
	})

	return client
}

func NewRetryPolicy(ctx context.Context, maxRetry int, timeBackOff time.Duration) retrypolicy.RetryPolicy[*http.Response] {
	return retrypolicy.Builder[*http.Response]().
		HandleIf(func(response *http.Response, _ error) bool {
			return response.StatusCode >= 400
		}).
		WithBackoff(time.Second, timeBackOff*time.Second).
		WithMaxRetries(maxRetry).
		OnRetryScheduled(func(e failsafe.ExecutionScheduledEvent[*http.Response]) {
			logger.Info(ctx, fmt.Sprintf("Retry %d after delay of %d", e.Attempts(), e.Delay))
		}).Build()
}

func NewFallbackPolicy(ctx context.Context, resp *http.Response) fallback.Fallback[*http.Response] {
	return fallback.BuilderWithResult[*http.Response](resp).
		OnFallbackExecuted(func(e failsafe.ExecutionDoneEvent[*http.Response]) {
			logger.Info(ctx, "Fallback executed result")
		}).
		Build()
}

func WriteJSON(w http.ResponseWriter, status int, data any, r *http.Request, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	logger.Info(r.Context(), fmt.Sprintf("Status: %v Response: %s", status, string(out)))

	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Erro log Request: Method %s Path %s - Body %v", r.Method, r.URL.Path, err))
			} else {
				logger.Info(ctx, fmt.Sprintf("Request: Method %s Path %s - Body %s", r.Method, r.URL.Path, string(bodyBytes)))
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		next.ServeHTTP(w, r)
	})
}

type Error struct {
	StatusCode    int
	Err           error
	ErrorMessages []AppError
}

func NewError(statusCode int, err error, message AppError, messages ...AppError) *Error {
	e := &Error{
		StatusCode:    statusCode,
		Err:           err,
		ErrorMessages: []AppError{message},
	}

	if len(messages) > 0 {
		e.ErrorMessages = append(e.ErrorMessages, messages...)
	}

	return e
}

func NewMessage(message string, code int) AppError {
	return AppError{Message: message, Code: code}
}

func HasError(e *Error) bool {

	if e == nil {
		return false
	}

	return e.Err != nil || len(e.ErrorMessages) > 0
}

func (e *Error) AddMessage(message string, code int) {
	e.ErrorMessages = append(e.ErrorMessages, AppError{Message: message, Code: code})
}

func (e *Error) Stringfy() string {
	str := fmt.Sprintf("StatusCode: %d - Error: %v", e.StatusCode, e.Err)
	for _, err := range e.ErrorMessages {
		str += fmt.Sprintf(" - Message: %s - Code: %d", err.Message, err.Code)
	}
	return str
}

func (e *Error) ToAppResponse(result any) AppResponse {
	return AppResponse{
		Success: false,
		Result:  result,
		Errors:  e.ErrorMessages,
	}
}
