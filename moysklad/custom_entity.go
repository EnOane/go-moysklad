package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CustomEntity Пользовательский справочник.
// Ключевое слово: customentity
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik
type CustomEntity struct {
	ID   *uuid.UUID `json:"id,omitempty"`   // ID Пользовательского справочника
	Meta *Meta      `json:"meta,omitempty"` // Метаданные Пользовательского справочника
	Name *string    `json:"name,omitempty"` // Наименование Пользовательского справочника
}

func (customEntity CustomEntity) Clean() *CustomEntity {
	return &CustomEntity{Meta: customEntity.Meta}
}

func (customEntity CustomEntity) GetID() uuid.UUID {
	return Deref(customEntity.ID)
}

func (customEntity CustomEntity) GetMeta() Meta {
	return Deref(customEntity.Meta)
}

func (customEntity CustomEntity) GetName() string {
	return Deref(customEntity.Name)
}

func (customEntity *CustomEntity) SetMeta(meta *Meta) *CustomEntity {
	customEntity.Meta = meta
	return customEntity
}

func (customEntity *CustomEntity) SetName(name string) *CustomEntity {
	customEntity.Name = &name
	return customEntity
}

func (customEntity CustomEntity) String() string {
	return Stringify(customEntity)
}

func (customEntity CustomEntity) MetaType() MetaType {
	return MetaTypeCustomEntity
}

// CustomEntityElement Элемент Пользовательского справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-jelementy-pol-zowatel-skogo-sprawochnika
type CustomEntityElement struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код элемента Пользовательского справочника
	Description  *string    `json:"description,omitempty"`  // Описание элемента Пользовательского справочника
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код элемента Пользовательского справочника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления элементе Пользовательского справочника
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
}

func (customEntityElement CustomEntityElement) GetAccountID() uuid.UUID {
	return Deref(customEntityElement.AccountID)
}

func (customEntityElement CustomEntityElement) GetCode() string {
	return Deref(customEntityElement.Code)
}

func (customEntityElement CustomEntityElement) GetDescription() string {
	return Deref(customEntityElement.Description)
}

func (customEntityElement CustomEntityElement) GetExternalCode() string {
	return Deref(customEntityElement.ExternalCode)
}

func (customEntityElement CustomEntityElement) GetID() uuid.UUID {
	return Deref(customEntityElement.ID)
}

func (customEntityElement CustomEntityElement) GetMeta() Meta {
	return Deref(customEntityElement.Meta)
}

func (customEntityElement CustomEntityElement) GetName() string {
	return Deref(customEntityElement.Name)
}

func (customEntityElement CustomEntityElement) GetUpdated() Timestamp {
	return Deref(customEntityElement.Updated)
}

func (customEntityElement CustomEntityElement) GetGroup() Group {
	return Deref(customEntityElement.Group)
}

func (customEntityElement CustomEntityElement) GetOwner() Employee {
	return Deref(customEntityElement.Owner)
}

func (customEntityElement CustomEntityElement) GetShared() bool {
	return Deref(customEntityElement.Shared)
}

func (customEntityElement *CustomEntityElement) SetCode(code string) *CustomEntityElement {
	customEntityElement.Code = &code
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetDescription(description string) *CustomEntityElement {
	customEntityElement.Description = &description
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetExternalCode(externalCode string) *CustomEntityElement {
	customEntityElement.ExternalCode = &externalCode
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetMeta(meta *Meta) *CustomEntityElement {
	customEntityElement.Meta = meta
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetName(name string) *CustomEntityElement {
	customEntityElement.Name = &name
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetGroup(group *Group) *CustomEntityElement {
	customEntityElement.Group = group
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetOwner(owner *Employee) *CustomEntityElement {
	customEntityElement.Owner = owner
	return customEntityElement
}

func (customEntityElement *CustomEntityElement) SetShared(shared bool) *CustomEntityElement {
	customEntityElement.Shared = &shared
	return customEntityElement
}

func (customEntityElement CustomEntityElement) String() string {
	return Stringify(customEntityElement)
}

// CustomEntityService
// Сервис для работы с пользовательскими справочниками.
type CustomEntityService interface {
	Create(ctx context.Context, customEntity *CustomEntity, params *Params) (*CustomEntity, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, customEntity *CustomEntity, params *Params) (*CustomEntity, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetElements(ctx context.Context, id uuid.UUID) (*List[CustomEntityElement], *resty.Response, error)
	CreateElement(ctx context.Context, id uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)
	DeleteElement(ctx context.Context, id, elementID uuid.UUID) (bool, *resty.Response, error)
	GetElementByID(ctx context.Context, id, elementID uuid.UUID) (*CustomEntityElement, *resty.Response, error)
	UpdateElement(ctx context.Context, id, elementID uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)
}

type customEntityService struct {
	Endpoint
	endpointCreate[CustomEntity]
	endpointUpdate[CustomEntity]
	endpointDelete
}

func NewCustomEntityService(client *Client) CustomEntityService {
	e := NewEndpoint(client, "entity/customentity")
	return &customEntityService{
		Endpoint:       e,
		endpointCreate: endpointCreate[CustomEntity]{e},
		endpointUpdate: endpointUpdate[CustomEntity]{e},
		endpointDelete: endpointDelete{e},
	}
}

// GetElements Получить элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alementy-sprawochnika
func (s *customEntityService) GetElements(ctx context.Context, id uuid.UUID) (*List[CustomEntityElement], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[List[CustomEntityElement]](s.client, path).Get(ctx)
}

// CreateElement Создать элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-sozdat-alement-sprawochnika
func (s *customEntityService) CreateElement(ctx context.Context, id uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Post(ctx, element)
}

// DeleteElement Удалить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-udalit-alement-sprawochnika
func (s *customEntityService) DeleteElement(ctx context.Context, id, elementID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetElementByID Получить отдельный элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alement
func (s *customEntityService) GetElementByID(ctx context.Context, id, elementID uuid.UUID) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementID)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Get(ctx)
}

// UpdateElement Изменить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-izmenit-alement
func (s *customEntityService) UpdateElement(ctx context.Context, id, elementID uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementID)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Put(ctx, element)
}
