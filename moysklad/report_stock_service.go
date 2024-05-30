package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ReportStockService
// Сервис для работы с отчётом "Остатки".
type ReportStockService interface {
	GetAll(ctx context.Context, params *Params) (*List[StockAll], *resty.Response, error)
	GetByStore(ctx context.Context, params *Params) (*List[StockByStore], *resty.Response, error)
	GetCurrentAll(ctx context.Context, params *Params) (*[]StockCurrentAll, *resty.Response, error)
	GetCurrentByStore(ctx context.Context, params *Params) (*[]StockCurrentByStore, *resty.Response, error)
	GetByOperationId(ctx context.Context, operationID *uuid.UUID, params *Params) (*List[StockByOperation], *resty.Response, error)
	GetAllAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockAll]], *resty.Response, error)
	GetByStoreAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockByStore]], *resty.Response, error)
}

type reportStockService struct {
	Endpoint
}

func NewReportStockService(client *Client) ReportStockService {
	e := NewEndpoint(client, "report/stock")
	return &reportStockService{e}
}

// GetAll Запрос на получение Расширенного отчета об остатках.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (s *reportStockService) GetAll(ctx context.Context, params *Params) (*List[StockAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", s.uri)
	return NewRequestBuilder[List[StockAll]](s.client, path).SetParams(params).Get(ctx)
}

// GetByStore Запрос на получение отчета "Остатки по складам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (s *reportStockService) GetByStore(ctx context.Context, params *Params) (*List[StockByStore], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", s.uri)
	return NewRequestBuilder[List[StockByStore]](s.client, path).SetParams(params).Get(ctx)
}

// GetCurrentAll Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (s *reportStockService) GetCurrentAll(ctx context.Context, params *Params) (*[]StockCurrentAll, *resty.Response, error) {
	path := fmt.Sprintf("%s/all/current", s.uri)
	return NewRequestBuilder[[]StockCurrentAll](s.client, path).SetParams(params).Get(ctx)
}

// GetCurrentByStore Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (s *reportStockService) GetCurrentByStore(ctx context.Context, params *Params) (*[]StockCurrentByStore, *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore/current", s.uri)
	return NewRequestBuilder[[]StockCurrentByStore](s.client, path).SetParams(params).Get(ctx)
}

// GetByOperationId
// Запрос на получение отчёта "Остатки по документу"
// Данный запрос работает со следующими типами документов:
// – Отгрузка
// – Заказ покупателя
// – Розничная продажа
// – Счет поставщика
// – Розничная продажа
// – Заказ поставщику
// – Приемка
// – Розничный возврат
// – Возврат поставщику
// – Возврат покупателя
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-dokumentu
func (s *reportStockService) GetByOperationId(ctx context.Context, operationId *uuid.UUID, params *Params) (*List[StockByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperation?operation.id=%s", s.uri, operationId)
	return NewRequestBuilder[List[StockByOperation]](s.client, path).SetParams(params).Get(ctx)
}

// GetAllAsync Запрос на получение Расширенного отчета об остатках (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (s *reportStockService) GetAllAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockAll]], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", s.uri)
	return NewRequestBuilder[List[StockAll]](s.client, path).SetParams(params).Async(ctx)
}

// GetByStoreAsync Запрос на получение отчета "Остатки по складам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (s *reportStockService) GetByStoreAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockByStore]], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", s.uri)
	return NewRequestBuilder[List[StockByStore]](s.client, path).SetParams(params).Async(ctx)
}
