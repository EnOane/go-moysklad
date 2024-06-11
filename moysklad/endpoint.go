package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

type mainService[E any, P any, M any, S any] struct {
	endpointGetOne[E]
	endpointGetList[E]
	endpointCreate[E]
	endpointCreateUpdateMany[E]
	endpointDeleteMany[E]
	endpointDelete
	endpointGetByID[E]
	endpointUpdate[E]
	endpointMetadata[M]
	endpointAttributes
	endpointNamedFilter
	endpointImages
	endpointSyncID[E]
	endpointAudit
	endpointPrintLabel
	endpointPositions[P]
	//endpointTemplate[E]
	endpointPublication
	endpointSettings[S]
	endpointGetOneAsync[E]
	endpointPrintTemplates
	endpointTrash
	endpointPrintDocument
	endpointAccounts
	endpointStates
	endpointFiles
}

func newMainService[E any, P any, M any, S any](e Endpoint) *mainService[E, P, M, S] {
	return &mainService[E, P, M, S]{
		endpointGetOne:           endpointGetOne[E]{e},
		endpointGetList:          endpointGetList[E]{e},
		endpointCreate:           endpointCreate[E]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[E]{e},
		endpointDeleteMany:       endpointDeleteMany[E]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[E]{e},
		endpointUpdate:           endpointUpdate[E]{e},
		endpointMetadata:         endpointMetadata[M]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointImages:           endpointImages{e},
		endpointSyncID:           endpointSyncID[E]{e},
		endpointAudit:            endpointAudit{e},
		endpointPrintLabel:       endpointPrintLabel{e},
		endpointPositions:        endpointPositions[P]{e},
		endpointPublication:      endpointPublication{e},
		endpointSettings:         endpointSettings[S]{e},
		endpointGetOneAsync:      endpointGetOneAsync[E]{e},
		endpointPrintTemplates:   endpointPrintTemplates{e},
		endpointTrash:            endpointTrash{e},
		endpointPrintDocument:    endpointPrintDocument{e},
		endpointAccounts:         endpointAccounts{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}

type endpointGetList[T any] struct{ Endpoint }

// GetList Запрос на получение списка объектов.
func (endpoint *endpointGetList[T]) GetList(ctx context.Context, params *Params) (*List[T], *resty.Response, error) {
	return NewRequestBuilder[List[T]](endpoint.client, endpoint.uri).SetParams(params).Get(ctx)
}

type endpointDelete struct{ Endpoint }

// Delete Запрос на удаление объекта по id.
func (endpoint *endpointDelete) Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointGetByID[T any] struct{ Endpoint }

// GetByID Запрос на получение объекта по id.
func (endpoint *endpointGetByID[T]) GetByID(ctx context.Context, id uuid.UUID, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params).Get(ctx)
}

type endpointGetOne[T any] struct{ Endpoint }

// Get Запрос (отдельный) на получение объекта. Например, ассортимент, контекст.
func (endpoint *endpointGetOne[T]) Get(ctx context.Context, params *Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](endpoint.client, endpoint.uri).SetParams(params).Get(ctx)
}

type endpointGetOneAsync[T any] struct{ Endpoint }

// GetAsync Запрос на асинхронное выполнение задачи.
// Первым возвращаемым аргументом является сервис для дальнейшей работы с конкретной асинхронной задачей.
func (endpoint *endpointGetOneAsync[T]) GetAsync(ctx context.Context) (AsyncResultService[T], *resty.Response, error) {
	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, endpoint.uri).
		SetParams(NewParams().withAsync()).
		Get(ctx)

	if err != nil {
		return nil, resp, nil
	}

	async := NewAsyncResultService[T](endpoint.client, resp)
	return async, resp, err
}

type endpointMetadata[T any] struct{ Endpoint }

// GetMetadata Получить метаданные объекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti
func (endpoint *endpointMetadata[T]) GetMetadata(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

type endpointTemplate[T MetaTyper] struct{ Endpoint }

// Template Получить предзаполненный стандартными полями объект.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-shablony-dokumentow
func (endpoint *endpointTemplate[T]) Template(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/new", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, nil)
}

// TemplateArg типы, которые могут быть использованы в качестве документа-основания
// при запросе на создание шаблона документа
//type TemplateArg interface {
//	InvoiceOutTemplateArg | SalesReturnTemplateArg | PurchaseReturnTemplateArg |
//		PaymentInTemplateArg | ProcessingOrderTemplateArg | PurchaseOrderTemplateArg |
//		PaymentOutTemplateArg | EnterTemplateArg | DemandTemplateArg | MoveTemplateArg |
//		CashInTemplateArg | CashOutTemplateArg | RetailDemandTemplateArg | LossTemplateArg |
//		InvoiceInTemplateArg | ProcessingTemplateArg
//}

//type endpointTemplateBasedOn[T MetaTyper, A TemplateArg] struct{ Endpoint }
//
//// TemplateBasedOn Получить предзаполненный стандартными полями объект на основании документа(-ов)
//func (endpoint *endpointTemplateBasedOn[T, A]) TemplateBasedOn(ctx context.Context, arg *A) (*T, *resty.Response, error) {
//	path := fmt.Sprintf("%s/new", endpoint.uri)
//	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, arg)
//}

type endpointCreate[T any] struct{ Endpoint }

// Create Запрос на создание объекта.
func (endpoint *endpointCreate[T]) Create(ctx context.Context, entity *T, params *Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](endpoint.client, endpoint.uri).SetParams(params).Post(ctx, entity)
}

// DeleteManyResponse объект ответа на запрос удаления нескольких элементов
type DeleteManyResponse []struct {
	Info      string    `json:"info"`
	ApiErrors ApiErrors `json:"errors"`
}

type endpointDeleteMany[T any] struct{ Endpoint }

// DeleteMany Запрос на удаление нескольких объектов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (endpoint *endpointDeleteMany[T]) DeleteMany(ctx context.Context, entities *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/delete", endpoint.uri)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, entities)
}

type endpointCreateUpdateMany[T any] struct{ Endpoint }

// CreateUpdateMany Запрос на создание и обновление нескольких объектов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sozdanie-i-obnowlenie-neskol-kih-ob-ektow
func (endpoint *endpointCreateUpdateMany[T]) CreateUpdateMany(ctx context.Context, entities []*T, params *Params) (*[]T, *resty.Response, error) {
	return NewRequestBuilder[[]T](endpoint.client, endpoint.uri).SetParams(params).Post(ctx, entities)
}

type endpointUpdate[T any] struct{ Endpoint }

// Update Запрос на обновление объекта.
func (endpoint *endpointUpdate[T]) Update(ctx context.Context, id uuid.UUID, entity *T, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params).Put(ctx, entity)
}

type endpointAccounts struct{ Endpoint }

// GetAccounts Получить все счета.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-scheta-kontragenta
func (endpoint *endpointAccounts) GetAccounts(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", endpoint.uri, id)
	return NewRequestBuilder[List[AgentAccount]](endpoint.client, path).Get(ctx)
}

// GetAccountByID Получить отдельный счёт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-schet-kontragenta
func (endpoint *endpointAccounts) GetAccountByID(ctx context.Context, id, accountID uuid.UUID) (*AgentAccount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts/%s", endpoint.uri, id, accountID)
	return NewRequestBuilder[AgentAccount](endpoint.client, path).Get(ctx)
}

// UpdateAccounts Изменить счета (списком).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico-izmenit-scheta-urlica
func (endpoint *endpointAccounts) UpdateAccounts(ctx context.Context, id uuid.UUID, accounts Slice[AgentAccount]) (*MetaArray[AgentAccount], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/accounts", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[AgentAccount]](endpoint.client, path).Post(ctx, accounts)
}

type endpointAttributes struct{ Endpoint }

// GetAttributes Получить все дополнительные поля для указанного типа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-poluchit-wse-dopolnitel-nye-polq-dlq-ukazannogo-tipa
func (endpoint *endpointAttributes) GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	return NewRequestBuilder[MetaArray[Attribute]](endpoint.client, path).Get(ctx)
}

// GetAttributeByID Получить дополнительное поле по id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-poluchit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Get(ctx)
}

// CreateAttribute Создать дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	return NewRequestBuilder[Attribute](endpoint.client, path).Post(ctx, attribute)
}

// CreateAttributes Создать несколько дополнительных полей.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-sozdat-dopolnitel-nye-polq
func (endpoint *endpointAttributes) CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes", endpoint.uri)
	// при передаче массива из 1-го доп поля сервис возвращает 1 доп поле, а не массив доп полей.
	// если количество передаваемых доп полей равняется 1, то дополнительно оборачиваем в срез.
	if len(attributeList) == 1 {
		attribute, resp, err := NewRequestBuilder[Attribute](endpoint.client, path).Post(ctx, attributeList[0])
		return &[]Attribute{*attribute}, resp, err
	}
	return NewRequestBuilder[[]Attribute](endpoint.client, path).Post(ctx, attributeList)
}

// UpdateAttribute Изменить дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-izmenit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[Attribute](endpoint.client, path).Put(ctx, attribute)
}

// DeleteAttribute Удалить дополнительное поле.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-noe-pole-udalit-dopolnitel-noe-pole
func (endpoint *endpointAttributes) DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteAttributes Удалить несколько дополнительных полей.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-dopolnitel-nye-polq-suschnostej-udalit-dopolnitel-nye-polq
func (endpoint *endpointAttributes) DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/attributes/delete", endpoint.uri)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, attributeList)
}

type endpointAudit struct{ Endpoint }

// GetAudit Запрос на получение событий по сущности с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-suschnosti
func (endpoint *endpointAudit) GetAudit(ctx context.Context, id uuid.UUID, params *Params) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/audit", endpoint.uri, id)
	return NewRequestBuilder[List[AuditEvent]](endpoint.client, path).SetParams(params).Get(ctx)
}

type endpointFiles struct{ Endpoint }

// GetFiles Получить список Файлов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-poluchit-spisok-fajlow-operacii-nomenklatury-zadachi-ili-kontragenta
func (endpoint *endpointFiles) GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[File]](endpoint.client, path).Get(ctx)
}

// CreateFile Добавить Файл.
func (endpoint *endpointFiles) CreateFile(ctx context.Context, id uuid.UUID, file *File) (*[]File, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[[]File](endpoint.client, path).Get(ctx)
}

// UpdateFiles Добавить/обновить Файлы.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-dobawit-fajly-k-operacii-nomenklature-ili-kontragentu
func (endpoint *endpointFiles) UpdateFiles(ctx context.Context, id uuid.UUID, files []*File) (*[]File, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files", endpoint.uri, id)
	return NewRequestBuilder[[]File](endpoint.client, path).Post(ctx, files)
}

// DeleteFile Удалить Файл.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-fajly-udalit-fajl
func (endpoint *endpointFiles) DeleteFile(ctx context.Context, id, fileID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files/%s", endpoint.uri, id, fileID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeleteFiles Удалить несколько Файлов.
func (endpoint *endpointFiles) DeleteFiles(ctx context.Context, id uuid.UUID, files *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/files/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, files)
}

type endpointImages struct{ Endpoint }

// GetImages Получить список Изображений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-poluchit-spisok-izobrazhenij-towara-komplekta-i-modifikacii
func (endpoint *endpointImages) GetImages(ctx context.Context, id uuid.UUID) (*MetaArray[Image], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[Image]](endpoint.client, path).Get(ctx)
}

// CreateImage Добавить Изображение.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-dobawit-izobrazhenie-k-towaru-komplektu-ili-modifikacii
func (endpoint *endpointImages) CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*[]*Image, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[[]*Image](endpoint.client, path).Post(ctx, image)
}

// UpdateImages Изменение Изображений (списком).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-izmenenie-spiska-izobrazhenij-u-towara-komplekta-ili-modifikacii
func (endpoint *endpointImages) UpdateImages(ctx context.Context, id uuid.UUID, images []*Image) (*[]Image, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images", endpoint.uri, id)
	return NewRequestBuilder[[]Image](endpoint.client, path).Post(ctx, images)
}

// DeleteImage Удалить Изображение.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-izobrazhenie
func (endpoint *endpointImages) DeleteImage(ctx context.Context, id, imageID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images/%s", endpoint.uri, id, imageID)
	return NewRequestBuilder[[]Image](endpoint.client, path).Delete(ctx)
}

// DeleteImages Удалить несколько Изображений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-izobrazhenie-udalit-gruppu-izobrazhenij
func (endpoint *endpointImages) DeleteImages(ctx context.Context, id uuid.UUID, images *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/images/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, images)
}

type endpointNamedFilter struct{ Endpoint }

// GetNamedFilters Получить список фильтров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-spisok-fil-trow
func (endpoint *endpointNamedFilter) GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter", endpoint.uri)
	return NewRequestBuilder[List[NamedFilter]](endpoint.client, path).SetParams(params).Get(ctx)
}

// GetNamedFilterByID Получить отдельный фильтр по id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try-poluchit-fil-tr-po-id
func (endpoint *endpointNamedFilter) GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error) {
	path := fmt.Sprintf("%s/namedfilter/%s", endpoint.uri, id)
	return NewRequestBuilder[NamedFilter](endpoint.client, path).Get(ctx)
}

type endpointPositions[T any] struct{ Endpoint }

// GetPositions Получить все позиции документа.
func (endpoint *endpointPositions[T]) GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[T], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[T]](endpoint.client, path).SetParams(params).Get(ctx)
}

// GetPositionByID Получение отдельной позиции.
func (endpoint *endpointPositions[T]) GetPositionByID(ctx context.Context, id, positionID uuid.UUID, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params).Get(ctx)
}

// UpdatePosition Обновление позиции.
func (endpoint *endpointPositions[T]) UpdatePosition(ctx context.Context, id, positionID uuid.UUID, position *T, params *Params) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[T](endpoint.client, path).SetParams(params).Put(ctx, position)
}

// CreatePosition Создание позиции документа.
func (endpoint *endpointPositions[T]) CreatePosition(ctx context.Context, id uuid.UUID, position *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[T](endpoint.client, path).Post(ctx, position)
}

// CreatePositions Массово создаёт позиции документа.
func (endpoint *endpointPositions[T]) CreatePositions(ctx context.Context, id uuid.UUID, positions []*T) (*[]T, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions", endpoint.uri, id)
	return NewRequestBuilder[[]T](endpoint.client, path).Post(ctx, positions)
}

// DeletePosition Удаляет позицию документа.
func (endpoint *endpointPositions[T]) DeletePosition(ctx context.Context, id, positionID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s", endpoint.uri, id, positionID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

// DeletePositionList запрос на удаление нескольких позиций документа.
func (endpoint *endpointPositions[T]) DeletePositionList(ctx context.Context, id uuid.UUID, list *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/delete", endpoint.uri, id)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, list)
}

// GetPositionTrackingCodes Получить Коды маркировки позиции документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-poluchit-kody-markirowki-pozicii-dokumenta
func (endpoint *endpointPositions[T]) GetPositionTrackingCodes(ctx context.Context, id, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", endpoint.uri, id, positionID)
	return NewRequestBuilder[MetaArray[TrackingCode]](endpoint.client, path).Get(ctx)
}

// CreateOrUpdatePositionTrackingCodes Массовое создание и обновление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-sozdanie-i-obnowlenie-kodow-markirowki
func (endpoint *endpointPositions[T]) CreateOrUpdatePositionTrackingCodes(ctx context.Context, id, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*[]TrackingCode, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes", endpoint.uri, id, positionID)
	return NewRequestBuilder[[]TrackingCode](endpoint.client, path).Post(ctx, trackingCodes)
}

// DeletePositionTrackingCodes Массовое удаление Кодов маркировки.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-massowoe-udalenie-kodow-markirowki
func (endpoint *endpointPositions[T]) DeletePositionTrackingCodes(ctx context.Context, id, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/positions/%s/trackingCodes/delete", endpoint.uri, id, positionID)
	return NewRequestBuilder[DeleteManyResponse](endpoint.client, path).Post(ctx, trackingCodes)
}

type endpointPrintDocument struct{ Endpoint }

// PrintDocument Запрос на печать документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-pechat-dokumentow-zapros-na-pechat
func (endpoint *endpointPrintDocument) PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", endpoint.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintDocumentArg)

	if err != nil {
		return nil, resp, err
	}

	file, err := GetFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPrintLabel struct{ Endpoint }

// PrintLabel Запрос на печать этикеток и ценников.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pechat-atiketok-i-cennikow
func (endpoint *endpointPrintLabel) PrintLabel(ctx context.Context, id uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/export", endpoint.uri, id)

	_, resp, err := NewRequestBuilder[PrintFile](endpoint.client, path).
		SetHeader(headerGetContent, "true").
		Post(ctx, PrintLabelArg)

	if err != nil {
		return nil, resp, err
	}

	file, err := GetFileFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}
	return file, resp, err
}

type endpointPublication struct{ Endpoint }

// GetPublications Запрос на получение списка Публикаций по указанному документу.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikacii
func (endpoint *endpointPublication) GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication", endpoint.uri, id)
	return NewRequestBuilder[MetaArray[Publication]](endpoint.client, path).Get(ctx)
}

// GetPublicationByID Запрос на получение Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-poluchit-publikaciu
func (endpoint *endpointPublication) GetPublicationByID(ctx context.Context, id, publicationID uuid.UUID) (*Publication, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", endpoint.uri, id, publicationID)
	return NewRequestBuilder[Publication](endpoint.client, path).Get(ctx)
}

// Publish Запрос на публикацию документа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-sozdat-publikaciu
func (endpoint *endpointPublication) Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error) {
	publication := new(Publication).SetTemplate(template)
	path := fmt.Sprintf("%s/%s/publication", endpoint.uri, id)
	return NewRequestBuilder[Publication](endpoint.client, path).Post(ctx, publication)
}

// DeletePublication Запрос на удаление Публикации с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow-udalit-publikaciu
func (endpoint *endpointPublication) DeletePublication(ctx context.Context, id, publicationID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/publication/%s", endpoint.uri, id, publicationID)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSettings[T any] struct{ Endpoint }

// GetSettings Запрос на получение настроек справочника.
func (endpoint *endpointSettings[T]) GetSettings(ctx context.Context) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// UpdateSettings Изменить настройки справочника.
func (endpoint *endpointSettings[T]) UpdateSettings(ctx context.Context, settings *T) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", endpoint.uri)
	return NewRequestBuilder[T](endpoint.client, path).Put(ctx, settings)
}

type endpointStates struct{ Endpoint }

// GetStateByID Запрос на получение статуса по id.
func (endpoint *endpointStates) GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Get(ctx)
}

// CreateState Создать новый статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-sozdat-status
func (endpoint *endpointStates) CreateState(ctx context.Context, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states", endpoint.uri)
	return NewRequestBuilder[State](endpoint.client, path).Post(ctx, state)
}

// UpdateState Изменить существующий статус.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-izmenit-status
func (endpoint *endpointStates) UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[State](endpoint.client, path).Put(ctx, state)
}

// CreateOrUpdateStates Массовое создание и обновление Статусов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-massowoe-sozdanie-i-obnowlenie-statusow
func (endpoint *endpointStates) CreateOrUpdateStates(ctx context.Context, states []*State) (*[]State, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states", endpoint.uri)
	return NewRequestBuilder[[]State](endpoint.client, path).Post(ctx, states)
}

// DeleteState Запрос на удаление Статуса с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-udalit-status
func (endpoint *endpointStates) DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/states/%s", endpoint.uri, id)
	return NewRequestBuilder[any](endpoint.client, path).Delete(ctx)
}

type endpointSyncID[T any] struct{ Endpoint }

// GetBySyncID Запрос на получение объекта по syncID.
func (endpoint *endpointSyncID[T]) GetBySyncID(ctx context.Context, syncID uuid.UUID) (*T, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Get(ctx)
}

// DeleteBySyncID Запрос на удаление объекта по syncID.
func (endpoint *endpointSyncID[T]) DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/syncid/%s", endpoint.uri, syncID)
	return NewRequestBuilder[T](endpoint.client, path).Delete(ctx)
}

type endpointPrintTemplates struct{ Endpoint }

// GetEmbeddedTemplates Запрос на получение информации о стандартных шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-standartnyh-shablonow
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate", endpoint.uri)
	return NewRequestBuilder[List[EmbeddedTemplate]](endpoint.client, path).Get(ctx)
}

// GetEmbeddedTemplateByID Запрос на получение информации об отдельном стандартном шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-standartnyj-shablon
func (endpoint *endpointPrintTemplates) GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", endpoint.uri, id)
	return NewRequestBuilder[EmbeddedTemplate](endpoint.client, path).Get(ctx)
}

// GetCustomTemplates Запрос на получение информации о пользовательских шаблонах печатных форм для указанного типа сущности.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-spisok-pol-zowatel-skih-shablonow
func (endpoint *endpointPrintTemplates) GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/customtemplate", endpoint.uri)
	return NewRequestBuilder[List[CustomTemplate]](endpoint.client, path).Get(ctx)
}

// GetCustomTemplateByID Запрос на получение информации об отдельном пользовательском шаблоне печатной формы для указанного типа сущности по его id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-otdel-nyj-pol-zowatel-skij-shablon
func (endpoint *endpointPrintTemplates) GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/embeddedtemplate/%s", endpoint.uri, id)
	return NewRequestBuilder[CustomTemplate](endpoint.client, path).Get(ctx)
}

type endpointTrash struct{ Endpoint }

// MoveToTrash Запрос на перемещение документа с указанным id в корзину.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-udalenie-w-korzinu
func (endpoint *endpointTrash) MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/trash", endpoint.uri, id)
	_, resp, err := NewRequestBuilder[any](endpoint.client, path).Post(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}
