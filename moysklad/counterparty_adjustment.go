package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"

	"time"
)

// CounterpartyAdjustment Корректировка взаиморасчетов.
//
// Код сущности: counterpartyadjustment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-wzaimoraschetow
type CounterpartyAdjustment struct {
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Корректировки взаиморасчетов
	Printed      *bool            `json:"printed,omitempty"`      // Напечатан ли документ
	AccountID    *string          `json:"accountId,omitempty"`    // ID учётной записи
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Files        *MetaArray[File] `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Applicable   *bool            `json:"applicable,omitempty"`   // Отметка о проведении
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Корректировки взаиморасчетов
	Created      *Timestamp       `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp       `json:"deleted,omitempty"`      // Момент последнего удаления Корректировки взаиморасчетов
	Description  *string          `json:"description,omitempty"`  // Комментарий Корректировки взаиморасчетов
	Name         *string          `json:"name,omitempty"`         // Наименование Корректировки взаиморасчетов
	Agent        *Agent           `json:"agent,omitempty"`        // Метаданные контрагента или сотрудника
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Корректировки взаиморасчетов
	Moment       *Timestamp       `json:"moment,omitempty"`       // Дата документа
	Organization *Organization    `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	ID           *string          `json:"id,omitempty"`           // ID Корректировки взаиморасчетов
	Published    *bool            `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Sum          *float64         `json:"sum,omitempty"`          // Сумма Корректировки взаиморасчетов в копейках
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (counterPartyAdjustment CounterpartyAdjustment) Clean() *CounterpartyAdjustment {
	if counterPartyAdjustment.Meta == nil {
		return nil
	}
	return &CounterpartyAdjustment{Meta: counterPartyAdjustment.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (counterPartyAdjustment CounterpartyAdjustment) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: counterPartyAdjustment.Meta}
}

// GetExternalCode возвращает Внешний код Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetExternalCode() string {
	return Deref(counterPartyAdjustment.ExternalCode)
}

// GetPrinted возвращает true, если документ напечатан.
func (counterPartyAdjustment CounterpartyAdjustment) GetPrinted() bool {
	return Deref(counterPartyAdjustment.Printed)
}

// GetAccountID возвращает ID учётной записи.
func (counterPartyAdjustment CounterpartyAdjustment) GetAccountID() string {
	return Deref(counterPartyAdjustment.AccountID)
}

// GetGroup возвращает Отдел сотрудника.
func (counterPartyAdjustment CounterpartyAdjustment) GetGroup() Group {
	return Deref(counterPartyAdjustment.Group)
}

// GetFiles возвращает Метаданные массива Файлов.
func (counterPartyAdjustment CounterpartyAdjustment) GetFiles() MetaArray[File] {
	return Deref(counterPartyAdjustment.Files)
}

// GetApplicable возвращает Отметку о проведении.
func (counterPartyAdjustment CounterpartyAdjustment) GetApplicable() bool {
	return Deref(counterPartyAdjustment.Applicable)
}

// GetUpdated возвращает Момент последнего обновления Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetUpdated() time.Time {
	return Deref(counterPartyAdjustment.Updated).Time()
}

// GetCreated возвращает Дату создания.
func (counterPartyAdjustment CounterpartyAdjustment) GetCreated() time.Time {
	return Deref(counterPartyAdjustment.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetDeleted() time.Time {
	return Deref(counterPartyAdjustment.Deleted).Time()
}

// GetDescription возвращает Комментарий Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetDescription() string {
	return Deref(counterPartyAdjustment.Description)
}

// GetName возвращает Наименование Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetName() string {
	return Deref(counterPartyAdjustment.Name)
}

// GetAgent возвращает Метаданные контрагента или сотрудника.
func (counterPartyAdjustment CounterpartyAdjustment) GetAgent() Agent {
	return Deref(counterPartyAdjustment.Agent)
}

// GetMeta возвращает Метаданные Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetMeta() Meta {
	return Deref(counterPartyAdjustment.Meta)
}

// GetMoment возвращает Дату документа.
func (counterPartyAdjustment CounterpartyAdjustment) GetMoment() time.Time {
	return Deref(counterPartyAdjustment.Moment).Time()
}

// GetOrganization возвращает Метаданные юрлица.
func (counterPartyAdjustment CounterpartyAdjustment) GetOrganization() Organization {
	return Deref(counterPartyAdjustment.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (counterPartyAdjustment CounterpartyAdjustment) GetOwner() Employee {
	return Deref(counterPartyAdjustment.Owner)
}

// GetID возвращает ID Корректировки взаиморасчетов.
func (counterPartyAdjustment CounterpartyAdjustment) GetID() string {
	return Deref(counterPartyAdjustment.ID)
}

// GetPublished возвращает true, если документ опубликован.
func (counterPartyAdjustment CounterpartyAdjustment) GetPublished() bool {
	return Deref(counterPartyAdjustment.Published)
}

// GetShared возвращает флаг общего доступа.
func (counterPartyAdjustment CounterpartyAdjustment) GetShared() bool {
	return Deref(counterPartyAdjustment.Shared)
}

// GetSum возвращает Сумму Корректировки взаиморасчетов в копейках.
func (counterPartyAdjustment CounterpartyAdjustment) GetSum() float64 {
	return Deref(counterPartyAdjustment.Sum)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (counterPartyAdjustment CounterpartyAdjustment) GetAttributes() Slice[Attribute] {
	return counterPartyAdjustment.Attributes
}

// SetExternalCode устанавливает Внешний код Корректировки взаиморасчетов.
func (counterPartyAdjustment *CounterpartyAdjustment) SetExternalCode(externalCode string) *CounterpartyAdjustment {
	counterPartyAdjustment.ExternalCode = &externalCode
	return counterPartyAdjustment
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (counterPartyAdjustment *CounterpartyAdjustment) SetGroup(group *Group) *CounterpartyAdjustment {
	counterPartyAdjustment.Group = group.Clean()
	return counterPartyAdjustment
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (counterPartyAdjustment *CounterpartyAdjustment) SetFiles(files ...*File) *CounterpartyAdjustment {
	counterPartyAdjustment.Files = NewMetaArrayFrom(files)
	return counterPartyAdjustment
}

// SetApplicable устанавливает Отметку о проведении.
func (counterPartyAdjustment *CounterpartyAdjustment) SetApplicable(applicable bool) *CounterpartyAdjustment {
	counterPartyAdjustment.Applicable = &applicable
	return counterPartyAdjustment
}

// SetDescription устанавливает Комментарий Корректировки взаиморасчетов.
func (counterPartyAdjustment *CounterpartyAdjustment) SetDescription(description string) *CounterpartyAdjustment {
	counterPartyAdjustment.Description = &description
	return counterPartyAdjustment
}

// SetName устанавливает Наименование Корректировки взаиморасчетов.
func (counterPartyAdjustment *CounterpartyAdjustment) SetName(name string) *CounterpartyAdjustment {
	counterPartyAdjustment.Name = &name
	return counterPartyAdjustment
}

// SetAgent устанавливает Метаданные контрагента.
//
// Принимает [Counterparty] или [Employee].
func (counterPartyAdjustment *CounterpartyAdjustment) SetAgent(agent AgentEmployeeConverter) *CounterpartyAdjustment {
	if agent != nil {
		counterPartyAdjustment.Agent = agent.AsEmployeeAgent()
	}
	return counterPartyAdjustment
}

// SetMeta устанавливает Метаданные Корректировки взаиморасчетов.
func (counterPartyAdjustment *CounterpartyAdjustment) SetMeta(meta *Meta) *CounterpartyAdjustment {
	counterPartyAdjustment.Meta = meta
	return counterPartyAdjustment
}

// SetMoment устанавливает Дату документа.
func (counterPartyAdjustment *CounterpartyAdjustment) SetMoment(moment time.Time) *CounterpartyAdjustment {
	counterPartyAdjustment.Moment = NewTimestamp(moment)
	return counterPartyAdjustment
}

// SetOrganization устанавливает Метаданные юрлица.
func (counterPartyAdjustment *CounterpartyAdjustment) SetOrganization(organization *Organization) *CounterpartyAdjustment {
	if organization != nil {
		counterPartyAdjustment.Organization = organization.Clean()
	}
	return counterPartyAdjustment
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (counterPartyAdjustment *CounterpartyAdjustment) SetOwner(owner *Employee) *CounterpartyAdjustment {
	if owner != nil {
		counterPartyAdjustment.Owner = owner.Clean()
	}
	return counterPartyAdjustment
}

// SetShared устанавливает флаг общего доступа.
func (counterPartyAdjustment *CounterpartyAdjustment) SetShared(shared bool) *CounterpartyAdjustment {
	counterPartyAdjustment.Shared = &shared
	return counterPartyAdjustment
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (counterPartyAdjustment *CounterpartyAdjustment) SetAttributes(attributes ...*Attribute) *CounterpartyAdjustment {
	counterPartyAdjustment.Attributes.Push(attributes...)
	return counterPartyAdjustment
}

// String реализует интерфейс [fmt.Stringer].
func (counterPartyAdjustment CounterpartyAdjustment) String() string {
	return Stringify(counterPartyAdjustment)
}

// MetaType возвращает код сущности.
func (CounterpartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterpartyAdjustment
}

// Update shortcut
func (counterPartyAdjustment *CounterpartyAdjustment) Update(ctx context.Context, client *Client, params ...func(*Params)) (*CounterpartyAdjustment, *resty.Response, error) {
	return NewCounterPartyAdjustmentService(client).Update(ctx, counterPartyAdjustment.GetID(), counterPartyAdjustment, params...)
}

// Create shortcut
func (counterPartyAdjustment *CounterpartyAdjustment) Create(ctx context.Context, client *Client, params ...func(*Params)) (*CounterpartyAdjustment, *resty.Response, error) {
	return NewCounterPartyAdjustmentService(client).Create(ctx, counterPartyAdjustment, params...)
}

// Delete shortcut
func (counterPartyAdjustment *CounterpartyAdjustment) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCounterPartyAdjustmentService(client).Delete(ctx, counterPartyAdjustment)
}

// CounterPartyAdjustmentService описывает методы сервиса для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	// GetList выполняет запрос на получение списка корректировок взаиморасчётов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[CounterpartyAdjustment], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех корректировок взаиморасчётов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[CounterpartyAdjustment], *resty.Response, error)

	// Create выполняет запрос на создание корректировки взаиморасчётов.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента или сотрудника)
	// Принимает контекст, корректировку взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает созданную корректировку взаиморасчётов.
	Create(ctx context.Context, counterPartyAdjustment *CounterpartyAdjustment, params ...func(*Params)) (*CounterpartyAdjustment, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение корректировок взаиморасчётов.
	// Изменяемые корректировки взаиморасчётов должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список корректировок взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых корректировок взаиморасчётов.
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList Slice[CounterpartyAdjustment], params ...func(*Params)) (*Slice[CounterpartyAdjustment], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление корректировок взаиморасчётов.
	// Принимает контекст и множество корректировок взаиморасчётов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CounterpartyAdjustment) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление корректировки взаиморасчётов по ID.
	// Принимает контекст и ID корректировки взаиморасчётов.
	// Возвращает «true» в случае успешного удаления корректировки взаиморасчётов.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление корректировки взаиморасчётов.
	// Принимает контекст и корректировку взаиморасчётов.
	// Возвращает «true» в случае успешного удаления корректировки взаиморасчётов.
	Delete(ctx context.Context, entity *CounterpartyAdjustment) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной корректировки взаиморасчётов по ID.
	// Принимает контекст, ID корректировки взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает найденную корректировку взаиморасчётов.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*CounterpartyAdjustment, *resty.Response, error)

	// Update выполняет запрос на изменение корректировки взаиморасчётов.
	// Принимает контекст, корректировку взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает изменённый корректировку взаиморасчётов.
	Update(ctx context.Context, id string, counterPartyAdjustment *CounterpartyAdjustment, params ...func(*Params)) (*CounterpartyAdjustment, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных корректировок взаиморасчётов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...func(*Params)) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id string) (*NamedFilter, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id string) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id string) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id string, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id string, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id string, fileID string) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id string, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

const (
	EndpointCounterpartyAdjustment = EndpointEntity + string(MetaTypeCounterpartyAdjustment)
)

// NewCounterPartyAdjustmentService принимает [Client] и возвращает сервис для работы с корректировками баланса контрагента.
func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	return newMainService[CounterpartyAdjustment, any, MetaAttributesStatesSharedWrapper, any](client, EndpointCounterpartyAdjustment)
}
