package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Thing Серийный номер
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-serijnyj-nomer
type Thing struct {
	AccountID   *string `json:"accountId,omitempty"`   // ID учётной записи
	Description *string `json:"description,omitempty"` // Описание Серийного номера
	ID          *string `json:"id,omitempty"`          // ID Серийного номера
	Meta        *Meta   `json:"meta,omitempty"`        // Метаданные о Серийном номере
	Name        *string `json:"name,omitempty"`        // Наименование Серийного номера
}

// GetAccountID возвращает ID учётной записи.
func (thing Thing) GetAccountID() string {
	return Deref(thing.AccountID)
}

// GetDescription возвращает Описание Серийного номера.
func (thing Thing) GetDescription() string {
	return Deref(thing.Description)
}

// GetID возвращает ID Серийного номера.
func (thing Thing) GetID() string {
	return Deref(thing.ID)
}

// GetMeta возвращает Метаданные о Серийном номере.
func (thing Thing) GetMeta() Meta {
	return Deref(thing.Meta)
}

// GetName возвращает Наименование Серийного номера.
func (thing Thing) GetName() string {
	return Deref(thing.Name)
}

// SetDescription устанавливает Описание Серийного номера.
func (thing *Thing) SetDescription(description string) *Thing {
	thing.Description = &description
	return thing
}

// SetMeta устанавливает Метаданные о Серийном номере.
func (thing *Thing) SetMeta(meta *Meta) *Thing {
	thing.Meta = meta
	return thing
}

// SetName устанавливает Наименование Серийного номера.
func (thing *Thing) SetName(name string) *Thing {
	thing.Name = &name
	return thing
}

// String реализует интерфейс [fmt.Stringer].
func (thing Thing) String() string {
	return Stringify(thing)
}

// MetaType возвращает код сущности.
func (Thing) MetaType() MetaType {
	return MetaTypeThing
}

// ThingService описывает методы сервиса для работы с серийными номерами.
type ThingService interface {
	// GetList выполняет запрос на получение списка серийных номеров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список серийных номеров.
	GetList(ctx context.Context, params ...func(*Params)) (*List[Thing], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех серийных номеров в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[Thing], *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного серийного номера по ID.
	// Принимает контекст, ID серийного номера и опционально объект параметров запроса Params.
	// Возвращает найденный серийный номер.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*Thing, *resty.Response, error)
}

const (
	EndpointThing = EndpointEntity + string(MetaTypeThing)
)

// NewThingService принимает [Client] и возвращает сервис для работы с серийными номерами.
func NewThingService(client *Client) ThingService {
	return newMainService[Thing, any, any, any](client, EndpointThing)
}
