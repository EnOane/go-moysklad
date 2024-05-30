package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

// AsyncService Сервис для работы с асинхронными задачами.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen
type AsyncService interface {
	GetStatuses(ctx context.Context, params *Params) (*List[Async], *resty.Response, error)
	GetStatusById(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *resty.Response, error)
}

type asyncService struct {
	Endpoint
}

func NewAsyncService(client *Client) AsyncService {
	e := NewEndpoint(client, "async")
	return &asyncService{e}
}

// GetStatuses Статусы Асинхронных задач.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-statusy-asinhronnyh-zadach
func (s *asyncService) GetStatuses(ctx context.Context, params *Params) (*List[Async], *resty.Response, error) {
	return NewRequestBuilder[List[Async]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetStatusById Получение статуса Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-statusa-asinhronnoj-zadachi
func (s *asyncService) GetStatusById(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *resty.Response, error) {
	path := fmt.Sprintf("async/%s", id)
	return NewRequestBuilder[Async](s.client, path).SetParams(params).Get(ctx)
}

type AsyncResultService[T any] interface {
	StatusURL() *url.URL
	ResultURL() *url.URL
	Check(ctx context.Context) (bool, *resty.Response, error)
	Result(ctx context.Context) (*T, *resty.Response, error)
	Cancel(ctx context.Context) (bool, *resty.Response, error)
}

type asyncResultService[T any] struct {
	req       *resty.Request
	statusURL *url.URL // URL статуса Асинхронной задачи.
	resultURL *url.URL // URL результата выполнения Асинхронной задачи.
}

// NewAsyncResultService Сервис для работы с асинхронной задачей.
func NewAsyncResultService[T any](req *resty.Request, resp *resty.Response) AsyncResultService[T] {
	statusUrlStr := resp.Header().Get("Location")
	resultUrlStr := resp.Header().Get("Content-Location")
	statusUrl, _ := url.Parse(statusUrlStr)
	resultUrl, _ := url.Parse(resultUrlStr)

	return &asyncResultService[T]{
		req:       req,
		statusURL: statusUrl,
		resultURL: resultUrl,
	}
}

// StatusURL возвращает URL проверки статуса асинхронной задачи.
func (s *asyncResultService[T]) StatusURL() *url.URL {
	return s.statusURL
}

// ResultURL возвращает URL результата выполнения асинхронной задачи.
func (s *asyncResultService[T]) ResultURL() *url.URL {
	return s.resultURL
}

// Check Проверяет статус асинхронной задачи.
// Если статус задачи = DONE, возвращает true, иначе false
func (s *asyncResultService[T]) Check(ctx context.Context) (bool, *resty.Response, error) {
	var async Async
	resp, err := s.req.SetContext(ctx).SetResult(async).Get(s.StatusURL().String())
	if err != nil {
		return false, resp, err
	}
	return async.State == AsyncStateDone, resp, nil
}

// Result Запрос на получение результата.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-rezul-tata-wypolneniq-asinhronnoj-zadachi
func (s *asyncResultService[T]) Result(ctx context.Context) (*T, *resty.Response, error) {
	var data = new(T)
	resp, err := s.req.SetContext(ctx).SetResult(data).Get(s.ResultURL().String())
	return data, resp, err
}

// Cancel Отмена Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-otmena-asinhronnoj-zadachi
func (s *asyncResultService[T]) Cancel(ctx context.Context) (bool, *resty.Response, error) {
	u, _ := s.StatusURL().Parse("/cancel")
	resp, err := s.req.SetContext(ctx).Post(u.String())
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}
