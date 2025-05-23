package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"

	"time"
)

// CustomerOrder Заказ покупателя.
//
// Код сущности: customerorder
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq
type CustomerOrder struct {
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Project               *NullValue[Project]               `json:"project,omitempty"`               // Метаданные проекта
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Moves                 Slice[Move]                       `json:"moves,omitempty"`                 // Массив ссылок на связанные перемещения
	Code                  *string                           `json:"code,omitempty"`                  // Код Заказа покупателя
	Agent                 *Counterparty                     `json:"agent,omitempty"`                 // Метаданные контрагента
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления Заказа покупателя
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата отгрузки
	Description           *string                           `json:"description,omitempty"`           // Комментарий Заказа покупателя
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код Заказа покупателя
	Files                 *MetaArray[File]                  `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *string                           `json:"id,omitempty"`                    // ID Заказа покупателя
	InvoicedSum           *float64                          `json:"invoicedSum,omitempty"`           // Сумма счетов покупателю
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные Заказа покупателя
	Name                  *string                           `json:"name,omitempty"`                  // Наименование Заказа покупателя
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Метаданные владельца (Сотрудника)
	PayedSum              *float64                          `json:"payedSum,omitempty"`              // Сумма входящих платежей по Заказу
	Positions             *MetaArray[CustomerOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа покупателя
	AccountID             *string                           `json:"accountId,omitempty"`             // ID учётной записи
	Contract              *NullValue[Contract]              `json:"contract,omitempty"`              // Метаданные договора
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *NullValue[Rate]                  `json:"rate,omitempty"`                  // Валюта
	ReservedSum           *float64                          `json:"reservedSum,omitempty"`           // Сумма товаров в резерве
	SalesChannel          *NullValue[SalesChannel]          `json:"salesChannel,omitempty"`          // Метаданные канала продаж
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	ShipmentAddress       *string                           `json:"shipmentAddress,omitempty"`       // Адрес доставки Заказа покупателя
	ShipmentAddressFull   *Address                          `json:"shipmentAddressFull,omitempty"`   // Адрес доставки Заказа покупателя с детализацией по отдельным полям
	ShippedSum            *float64                          `json:"shippedSum,omitempty"`            // Сумма отгруженного
	State                 *NullValue[State]                 `json:"state,omitempty"`                 // Метаданные статуса заказа
	Store                 *NullValue[Store]                 `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *float64                          `json:"sum,omitempty"`                   // Сумма Заказа в установленной валюте
	SyncID                *string                           `json:"syncId,omitempty"`                // ID синхронизации
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления Заказа покупателя
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *float64                          `json:"vatSum,omitempty"`                // Сумма НДС
	Prepayments           Slice[Prepayment]                 `json:"prepayments,omitempty"`           // Массив ссылок на связанные предоплаты
	PurchaseOrders        Slice[PurchaseOrder]              `json:"purchaseOrders,omitempty"`        // Массив ссылок на связанные заказы поставщикам
	Demands               Slice[Demand]                     `json:"demands,omitempty"`               // Массив ссылок на связанные отгрузки
	Payments              Slice[Payment]                    `json:"payments,omitempty"`              // Массив ссылок на связанные платежи
	InvoicesOut           Slice[InvoiceOut]                 `json:"invoicesOut,omitempty"`           // Массив ссылок на связанные счета покупателям
	TaxSystem             TaxSystem                         `json:"taxSystem,omitempty"`             // Код системы налогообложения
	Attributes            Slice[Attribute]                  `json:"attributes,omitempty"`            // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (customerOrder CustomerOrder) Clean() *CustomerOrder {
	if customerOrder.Meta == nil {
		return nil
	}
	return &CustomerOrder{Meta: customerOrder.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (customerOrder CustomerOrder) AsOperation() *Operation {
	return newOperation(customerOrder)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (customerOrder CustomerOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: customerOrder.Meta}
}

// AsOperationIn реализует интерфейс [OperationInConverter].
func (commissionReportOut CommissionReportOut) AsOperationIn() *Operation {
	return commissionReportOut.AsOperation()
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (customerOrder CustomerOrder) GetOrganizationAccount() AgentAccount {
	return Deref(customerOrder.OrganizationAccount)
}

// GetProject возвращает Метаданные проекта.
func (customerOrder CustomerOrder) GetProject() Project {
	return Deref(customerOrder.Project).getValue()
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (customerOrder CustomerOrder) GetAgentAccount() AgentAccount {
	return Deref(customerOrder.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (customerOrder CustomerOrder) GetApplicable() bool {
	return Deref(customerOrder.Applicable)
}

// GetMoves возвращает Массив ссылок на связанные перемещения.
func (customerOrder CustomerOrder) GetMoves() Slice[Move] {
	return customerOrder.Moves
}

// GetCode возвращает Код Заказа покупателя.
func (customerOrder CustomerOrder) GetCode() string {
	return Deref(customerOrder.Code)
}

// GetAgent возвращает Метаданные контрагента.
func (customerOrder CustomerOrder) GetAgent() Counterparty {
	return Deref(customerOrder.Agent)
}

// GetCreated возвращает Дату создания.
func (customerOrder CustomerOrder) GetCreated() time.Time {
	return Deref(customerOrder.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Заказа покупателя.
func (customerOrder CustomerOrder) GetDeleted() time.Time {
	return Deref(customerOrder.Deleted).Time()
}

// GetDeliveryPlannedMoment возвращает Планируемую дата отгрузки.
func (customerOrder CustomerOrder) GetDeliveryPlannedMoment() time.Time {
	return Deref(customerOrder.DeliveryPlannedMoment).Time()
}

// GetDescription возвращает Комментарий Заказа покупателя.
func (customerOrder CustomerOrder) GetDescription() string {
	return Deref(customerOrder.Description)
}

// GetExternalCode возвращает Внешний код Заказа покупателя.
func (customerOrder CustomerOrder) GetExternalCode() string {
	return Deref(customerOrder.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (customerOrder CustomerOrder) GetFiles() MetaArray[File] {
	return Deref(customerOrder.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (customerOrder CustomerOrder) GetGroup() Group {
	return Deref(customerOrder.Group)
}

// GetID возвращает ID Заказа покупателя.
func (customerOrder CustomerOrder) GetID() string {
	return Deref(customerOrder.ID)
}

// GetInvoicedSum возвращает Сумму счетов покупателю.
func (customerOrder CustomerOrder) GetInvoicedSum() float64 {
	return Deref(customerOrder.InvoicedSum)
}

// GetMeta возвращает Метаданные Заказа покупателя.
func (customerOrder CustomerOrder) GetMeta() Meta {
	return Deref(customerOrder.Meta)
}

// GetName возвращает Наименование Заказа покупателя.
func (customerOrder CustomerOrder) GetName() string {
	return Deref(customerOrder.Name)
}

// GetMoment возвращает Дату документа.
func (customerOrder CustomerOrder) GetMoment() time.Time {
	return Deref(customerOrder.Moment).Time()
}

// GetOrganization возвращает Метаданные юрлица.
func (customerOrder CustomerOrder) GetOrganization() Organization {
	return Deref(customerOrder.Organization)
}

// GetPrinted возвращает true, если документ напечатан.
func (customerOrder CustomerOrder) GetPrinted() bool {
	return Deref(customerOrder.Printed)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (customerOrder CustomerOrder) GetOwner() Employee {
	return Deref(customerOrder.Owner)
}

// GetPayedSum возвращает Оплаченную сумму.
func (customerOrder CustomerOrder) GetPayedSum() float64 {
	return Deref(customerOrder.PayedSum)
}

// GetPositions возвращает Метаданные позиций Заказа покупателя.
func (customerOrder CustomerOrder) GetPositions() MetaArray[CustomerOrderPosition] {
	return Deref(customerOrder.Positions)
}

// GetAccountID возвращает ID учётной записи.
func (customerOrder CustomerOrder) GetAccountID() string {
	return Deref(customerOrder.AccountID)
}

// GetContract возвращает Метаданные договора.
func (customerOrder CustomerOrder) GetContract() Contract {
	return Deref(customerOrder.Contract).getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (customerOrder CustomerOrder) GetPublished() bool {
	return Deref(customerOrder.Published)
}

// GetRate возвращает Валюту.
func (customerOrder CustomerOrder) GetRate() Rate {
	return Deref(customerOrder.Rate).getValue()
}

// GetReservedSum возвращает Сумму товаров в резерве.
func (customerOrder CustomerOrder) GetReservedSum() float64 {
	return Deref(customerOrder.ReservedSum)
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (customerOrder CustomerOrder) GetSalesChannel() SalesChannel {
	return Deref(customerOrder.SalesChannel).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (customerOrder CustomerOrder) GetShared() bool {
	return Deref(customerOrder.Shared)
}

// GetShipmentAddress возвращает Адрес доставки Заказа покупателя.
func (customerOrder CustomerOrder) GetShipmentAddress() string {
	return Deref(customerOrder.ShipmentAddress)
}

// GetShipmentAddressFull возвращает Адрес доставки Заказа покупателя с детализацией по отдельным полям.
func (customerOrder CustomerOrder) GetShipmentAddressFull() Address {
	return Deref(customerOrder.ShipmentAddressFull)
}

// GetShippedSum возвращает Сумму отгруженного.
func (customerOrder CustomerOrder) GetShippedSum() float64 {
	return Deref(customerOrder.ShippedSum)
}

// GetState возвращает Метаданные статуса Заказа покупателя.
func (customerOrder CustomerOrder) GetState() State {
	return Deref(customerOrder.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (customerOrder CustomerOrder) GetStore() Store {
	return Deref(customerOrder.Store).getValue()
}

// GetSum возвращает Сумму Заказа в установленной валюте.
func (customerOrder CustomerOrder) GetSum() float64 {
	return Deref(customerOrder.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (customerOrder CustomerOrder) GetSyncID() string {
	return Deref(customerOrder.SyncID)
}

// GetPrepayments возвращает Массив ссылок на связанные предоплаты.
func (customerOrder CustomerOrder) GetPrepayments() Slice[Prepayment] {
	return customerOrder.Prepayments
}

// GetUpdated возвращает Момент последнего обновления Заказа покупателя.
func (customerOrder CustomerOrder) GetUpdated() time.Time {
	return Deref(customerOrder.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (customerOrder CustomerOrder) GetVatEnabled() bool {
	return Deref(customerOrder.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (customerOrder CustomerOrder) GetVatIncluded() bool {
	return Deref(customerOrder.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (customerOrder CustomerOrder) GetVatSum() float64 {
	return Deref(customerOrder.VatSum)
}

// GetPurchaseOrders возвращает Массив ссылок на связанные заказы поставщикам.
func (customerOrder CustomerOrder) GetPurchaseOrders() Slice[PurchaseOrder] {
	return customerOrder.PurchaseOrders
}

// GetDemands возвращает Массив ссылок на связанные отгрузки.
func (customerOrder CustomerOrder) GetDemands() Slice[Demand] {
	return customerOrder.Demands
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (customerOrder CustomerOrder) GetPayments() Slice[Payment] {
	return customerOrder.Payments
}

// GetInvoicesOut возвращает Массив ссылок на связанные счета покупателям.
func (customerOrder CustomerOrder) GetInvoicesOut() Slice[InvoiceOut] {
	return customerOrder.InvoicesOut
}

// GetTaxSystem возвращает Код системы налогообложения.
func (customerOrder CustomerOrder) GetTaxSystem() TaxSystem {
	return customerOrder.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (customerOrder CustomerOrder) GetAttributes() Slice[Attribute] {
	return customerOrder.Attributes
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (customerOrder *CustomerOrder) SetOrganizationAccount(organizationAccount *AgentAccount) *CustomerOrder {
	if organizationAccount != nil {
		customerOrder.OrganizationAccount = organizationAccount.Clean()
	}
	return customerOrder
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetProject(project *Project) *CustomerOrder {
	customerOrder.Project = NewNullValue(project)
	return customerOrder
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (customerOrder *CustomerOrder) SetAgentAccount(agentAccount *AgentAccount) *CustomerOrder {
	if agentAccount != nil {
		customerOrder.AgentAccount = agentAccount.Clean()
	}
	return customerOrder
}

// SetApplicable устанавливает Отметку о проведении.
func (customerOrder *CustomerOrder) SetApplicable(applicable bool) *CustomerOrder {
	customerOrder.Applicable = &applicable
	return customerOrder
}

// SetMoves устанавливает Массив ссылок на связанные перемещения.
//
// Принимает множество объектов [Move].
func (customerOrder *CustomerOrder) SetMoves(moves ...*Move) *CustomerOrder {
	customerOrder.Moves.Push(moves...)
	return customerOrder
}

// SetCode устанавливает Код Заказа покупателя.
func (customerOrder *CustomerOrder) SetCode(code string) *CustomerOrder {
	customerOrder.Code = &code
	return customerOrder
}

// SetAgent устанавливает Метаданные Контрагента.
func (customerOrder *CustomerOrder) SetAgent(agent *Counterparty) *CustomerOrder {
	if agent != nil {
		customerOrder.Agent = agent.Clean()
	}
	return customerOrder
}

// SetDeliveryPlannedMoment устанавливает Планируемую дату отгрузки.
func (customerOrder *CustomerOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment time.Time) *CustomerOrder {
	customerOrder.DeliveryPlannedMoment = NewTimestamp(deliveryPlannedMoment)
	return customerOrder
}

// SetDescription устанавливает Комментарий Заказа покупателя.
func (customerOrder *CustomerOrder) SetDescription(description string) *CustomerOrder {
	customerOrder.Description = &description
	return customerOrder
}

// SetExternalCode устанавливает Внешний код Заказа покупателя.
func (customerOrder *CustomerOrder) SetExternalCode(externalCode string) *CustomerOrder {
	customerOrder.ExternalCode = &externalCode
	return customerOrder
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (customerOrder *CustomerOrder) SetFiles(files ...*File) *CustomerOrder {
	customerOrder.Files = NewMetaArrayFrom(files)
	return customerOrder
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (customerOrder *CustomerOrder) SetGroup(group *Group) *CustomerOrder {
	if group != nil {
		customerOrder.Group = group.Clean()
	}
	return customerOrder
}

// SetMeta устанавливает Метаданные Заказа покупателя.
func (customerOrder *CustomerOrder) SetMeta(meta *Meta) *CustomerOrder {
	customerOrder.Meta = meta
	return customerOrder
}

// SetName устанавливает Наименование Заказа покупателя.
func (customerOrder *CustomerOrder) SetName(name string) *CustomerOrder {
	customerOrder.Name = &name
	return customerOrder
}

// SetMoment устанавливает Дату документа.
func (customerOrder *CustomerOrder) SetMoment(moment time.Time) *CustomerOrder {
	customerOrder.Moment = NewTimestamp(moment)
	return customerOrder
}

// SetOrganization устанавливает Метаданные юрлица.
func (customerOrder *CustomerOrder) SetOrganization(organization *Organization) *CustomerOrder {
	if organization != nil {
		customerOrder.Organization = organization.Clean()
	}
	return customerOrder
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (customerOrder *CustomerOrder) SetOwner(owner *Employee) *CustomerOrder {
	if owner != nil {
		customerOrder.Owner = owner.Clean()
	}
	return customerOrder
}

// SetPositions устанавливает Метаданные позиций Заказа покупателя.
//
// Принимает множество объектов [CustomerOrderPosition].
func (customerOrder *CustomerOrder) SetPositions(positions ...*CustomerOrderPosition) *CustomerOrder {
	customerOrder.Positions = NewMetaArrayFrom(positions)
	return customerOrder
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetContract(contract *Contract) *CustomerOrder {
	customerOrder.Contract = NewNullValue(contract)
	return customerOrder
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetRate(rate *Rate) *CustomerOrder {
	customerOrder.Rate = NewNullValue(rate)
	return customerOrder
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetSalesChannel(salesChannel *SalesChannel) *CustomerOrder {
	customerOrder.SalesChannel = NewNullValue(salesChannel)
	return customerOrder
}

// SetShared устанавливает флаг общего доступа.
func (customerOrder *CustomerOrder) SetShared(shared bool) *CustomerOrder {
	customerOrder.Shared = &shared
	return customerOrder
}

// SetShipmentAddress устанавливает Адрес доставки Заказа покупателя.
func (customerOrder *CustomerOrder) SetShipmentAddress(shipmentAddress string) *CustomerOrder {
	customerOrder.ShipmentAddress = &shipmentAddress
	return customerOrder
}

// SetShipmentAddressFull устанавливает Адрес доставки Заказа покупателя с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetShipmentAddressFull(shipmentAddressFull *Address) *CustomerOrder {
	if shipmentAddressFull == nil {
		customerOrder.SetShipmentAddress("")
	} else {
		customerOrder.ShipmentAddressFull = shipmentAddressFull
	}
	return customerOrder
}

// SetState устанавливает Метаданные статуса Заказа покупателя.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetState(state *State) *CustomerOrder {
	customerOrder.State = NewNullValue(state)
	return customerOrder
}

// SetStore устанавливает Метаданные склада.
//
// Передача nil передаёт сброс значения (null).
func (customerOrder *CustomerOrder) SetStore(store *Store) *CustomerOrder {
	customerOrder.Store = NewNullValue(store)
	return customerOrder
}

// SetSyncID устанавливает ID синхронизации.
func (customerOrder *CustomerOrder) SetSyncID(syncID string) *CustomerOrder {
	customerOrder.SyncID = &syncID
	return customerOrder
}

// SetPrepayments устанавливает Массив ссылок на связанные предоплаты.
//
// Принимает множество объектов [Prepayment].
func (customerOrder *CustomerOrder) SetPrepayments(prepayments ...*Prepayment) *CustomerOrder {
	customerOrder.Prepayments.Push(prepayments...)
	return customerOrder
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Заказа покупателя.
func (customerOrder *CustomerOrder) SetVatEnabled(vatEnabled bool) *CustomerOrder {
	customerOrder.VatEnabled = &vatEnabled
	return customerOrder
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (customerOrder *CustomerOrder) SetVatIncluded(vatIncluded bool) *CustomerOrder {
	customerOrder.VatIncluded = &vatIncluded
	return customerOrder
}

// SetPurchaseOrders устанавливает Массив ссылок на связанные заказы поставщикам.
//
// Принимает множество объектов [PurchaseOrder].
func (customerOrder *CustomerOrder) SetPurchaseOrders(purchaseOrders ...*PurchaseOrder) *CustomerOrder {
	customerOrder.PurchaseOrders.Push(purchaseOrders...)
	return customerOrder
}

// SetDemands устанавливает Массив ссылок на связанные отгрузки.
//
// Принимает множество объектов [Demand].
func (customerOrder *CustomerOrder) SetDemands(demands ...*Demand) *CustomerOrder {
	customerOrder.Demands.Push(demands...)
	return customerOrder
}

// SetPayments устанавливает Массив ссылок на связанные платежи.
//
// Принимает множество объектов [Payment].
func (customerOrder *CustomerOrder) SetPayments(payments ...PaymentConverter) *CustomerOrder {
	customerOrder.Payments = NewPaymentsFrom(payments)
	return customerOrder
}

// SetInvoicesOut устанавливает Массив ссылок на связанные счета покупателям.
//
// Принимает множество объектов [InvoiceOut].
func (customerOrder *CustomerOrder) SetInvoicesOut(invoicesOut ...*InvoiceOut) *CustomerOrder {
	customerOrder.InvoicesOut.Push(invoicesOut...)
	return customerOrder
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (customerOrder *CustomerOrder) SetTaxSystem(taxSystem TaxSystem) *CustomerOrder {
	customerOrder.TaxSystem = taxSystem
	return customerOrder
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (customerOrder *CustomerOrder) SetAttributes(attributes ...*Attribute) *CustomerOrder {
	customerOrder.Attributes.Push(attributes...)
	return customerOrder
}

// String реализует интерфейс [fmt.Stringer].
func (customerOrder CustomerOrder) String() string {
	return Stringify(customerOrder)
}

// MetaType возвращает код сущности.
func (CustomerOrder) MetaType() MetaType {
	return MetaTypeCustomerOrder
}

// Update shortcut
func (customerOrder *CustomerOrder) Update(ctx context.Context, client *Client, params ...func(*Params)) (*CustomerOrder, *resty.Response, error) {
	return NewCustomerOrderService(client).Update(ctx, customerOrder.GetID(), customerOrder, params...)
}

// Create shortcut
func (customerOrder *CustomerOrder) Create(ctx context.Context, client *Client, params ...func(*Params)) (*CustomerOrder, *resty.Response, error) {
	return NewCustomerOrderService(client).Create(ctx, customerOrder, params...)
}

// Delete shortcut
func (customerOrder *CustomerOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCustomerOrderService(client).Delete(ctx, customerOrder)
}

// CustomerOrderPosition Позиция Заказа покупателя.
//
// Код сущности: customerorderposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq-zakazy-pokupatelej-pozicii-zakaza-pokupatelq
type CustomerOrderPosition struct {
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *string             `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	AccountID  *string             `json:"accountId,omitempty"`  // ID учётной записи
	Reserve    *float64            `json:"reserve,omitempty"`    // Резерв данной позиции
	Shipped    *float64            `json:"shipped,omitempty"`    // Доставлено
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	TaxSystem  TaxSystem           `json:"taxSystem,omitempty"`  // Код системы налогообложения
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (customerOrderPosition CustomerOrderPosition) GetQuantity() float64 {
	return Deref(customerOrderPosition.Quantity)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (customerOrderPosition CustomerOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(customerOrderPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
func (customerOrderPosition CustomerOrderPosition) GetDiscount() float64 {
	return Deref(customerOrderPosition.Discount)
}

// GetID возвращает ID позиции.
func (customerOrderPosition CustomerOrderPosition) GetID() string {
	return Deref(customerOrderPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (customerOrderPosition CustomerOrderPosition) GetPack() Pack {
	return Deref(customerOrderPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (customerOrderPosition CustomerOrderPosition) GetPrice() float64 {
	return Deref(customerOrderPosition.Price)
}

// GetAccountID возвращает ID учётной записи.
func (customerOrderPosition CustomerOrderPosition) GetAccountID() string {
	return Deref(customerOrderPosition.AccountID)
}

// GetReserve возвращает Резерв данной позиции.
func (customerOrderPosition CustomerOrderPosition) GetReserve() float64 {
	return Deref(customerOrderPosition.Reserve)
}

// GetShipped возвращает Доставлено.
func (customerOrderPosition CustomerOrderPosition) GetShipped() float64 {
	return Deref(customerOrderPosition.Shipped)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (customerOrderPosition CustomerOrderPosition) GetVat() int {
	return Deref(customerOrderPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (customerOrderPosition CustomerOrderPosition) GetVatEnabled() bool {
	return Deref(customerOrderPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (customerOrderPosition CustomerOrderPosition) GetStock() Stock {
	return Deref(customerOrderPosition.Stock)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (customerOrderPosition CustomerOrderPosition) GetTaxSystem() TaxSystem {
	return customerOrderPosition.TaxSystem
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (customerOrderPosition *CustomerOrderPosition) SetQuantity(quantity float64) *CustomerOrderPosition {
	customerOrderPosition.Quantity = &quantity
	return customerOrderPosition
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (customerOrderPosition *CustomerOrderPosition) SetAssortment(assortment AssortmentConverter) *CustomerOrderPosition {
	if assortment != nil {
		customerOrderPosition.Assortment = assortment.AsAssortment()
	}
	return customerOrderPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (customerOrderPosition *CustomerOrderPosition) SetDiscount(discount float64) *CustomerOrderPosition {
	customerOrderPosition.Discount = &discount
	return customerOrderPosition
}

// SetPack устанавливает Упаковку Товара.
func (customerOrderPosition *CustomerOrderPosition) SetPack(pack *Pack) *CustomerOrderPosition {
	customerOrderPosition.Pack = pack
	return customerOrderPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (customerOrderPosition *CustomerOrderPosition) SetPrice(price float64) *CustomerOrderPosition {
	customerOrderPosition.Price = &price
	return customerOrderPosition
}

// SetReserve устанавливает Резерв данной позиции.
func (customerOrderPosition *CustomerOrderPosition) SetReserve(reserve float64) *CustomerOrderPosition {
	customerOrderPosition.Reserve = &reserve
	return customerOrderPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (customerOrderPosition *CustomerOrderPosition) SetVat(vat int) *CustomerOrderPosition {
	customerOrderPosition.Vat = &vat
	return customerOrderPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (customerOrderPosition *CustomerOrderPosition) SetVatEnabled(vatEnabled bool) *CustomerOrderPosition {
	customerOrderPosition.VatEnabled = &vatEnabled
	return customerOrderPosition
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (customerOrderPosition *CustomerOrderPosition) SetTaxSystem(taxSystem TaxSystem) *CustomerOrderPosition {
	customerOrderPosition.TaxSystem = taxSystem
	return customerOrderPosition
}

// String реализует интерфейс [fmt.Stringer].
func (customerOrderPosition CustomerOrderPosition) String() string {
	return Stringify(customerOrderPosition)
}

// MetaType возвращает код сущности.
func (CustomerOrderPosition) MetaType() MetaType {
	return MetaTypeCustomerOrderPosition
}

// EventNote Лента событий.
//
// Код сущности: eventnote.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-lenta-sobytij
type EventNote struct {
	Meta        Meta      `json:"meta,omitempty"`        // Метаданные События
	Created     Timestamp `json:"created,omitempty"`     // Момент создания События
	Description string    `json:"description,omitempty"` // Текст События
	ID          string    `json:"id,omitempty"`          // ID События
	AccountID   string    `json:"accountId,omitempty"`   // ID учётной записи
	Author      Employee  `json:"author,omitempty"`      // Метаданные Сотрудника - создателя События
}

type customerOrderService struct {
	Endpoint
	endpointGetList[CustomerOrder]
	endpointCreate[CustomerOrder]
	endpointCreateUpdateMany[CustomerOrder]
	endpointGetByID[CustomerOrder]
	endpointUpdate[CustomerOrder]
	endpointMetadata[MetaAttributesStatesSharedWrapper]
	endpointPositions[CustomerOrderPosition]
	endpointAttributes
	endpointPublication
	endpointStates
	endpointFiles
	endpointPrintTemplates
	endpointSyncID[CustomerOrder]
	endpointDeleteByID
	endpointDelete[CustomerOrder]
	endpointNamedFilter
	endpointDeleteMany[CustomerOrder]
	endpointTrash
	endpointTemplate[CustomerOrder]
	endpointEvaluate[CustomerOrder]
}

// CustomerOrderService описывает методы сервис для работы с заказами покупателя.
type CustomerOrderService interface {
	// GetList выполняет запрос на получение списка заказов покупателей.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[CustomerOrder], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех заказов покупателей в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[CustomerOrder], *resty.Response, error)

	// Create выполняет запрос на создание заказа покупателя.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента (покупателя))
	// Принимает контекст, заказ покупателя и опционально объект параметров запроса Params.
	// Возвращает созданный заказ покупателя.
	Create(ctx context.Context, customerOrder *CustomerOrder, params ...func(*Params)) (*CustomerOrder, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение заказов покупателей.
	// Изменяемые заказы покупателей должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список заказов покупателей и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых заказов покупателей.
	CreateUpdateMany(ctx context.Context, customerOrderList Slice[CustomerOrder], params ...func(*Params)) (*Slice[CustomerOrder], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление заказов покупателей.
	// Принимает контекст и множество заказов покупателей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CustomerOrder) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление заказа покупателя по ID.
	// Принимает контекст и ID заказа покупателя.
	// Возвращает «true» в случае успешного удаления заказа покупателя.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление заказа покупателя.
	// Принимает контекст и заказ покупателя.
	// Возвращает «true» в случае успешного удаления заказа покупателя.
	Delete(ctx context.Context, entity *CustomerOrder) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного заказа покупателя по ID.
	// Принимает контекст, ID заказа покупателя и опционально объект параметров запроса Params.
	// Возвращает найденный заказ покупателя.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*CustomerOrder, *resty.Response, error)

	// Update выполняет запрос на изменение заказа покупателя.
	// Принимает контекст, заказ покупателя и опционально объект параметров запроса Params.
	// Возвращает изменённый заказ покупателя.
	Update(ctx context.Context, id string, customerOrder *CustomerOrder, params ...func(*Params)) (*CustomerOrder, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных заказов покупателей.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id string, params ...func(*Params)) (*List[CustomerOrderPosition], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id string, params ...func(*Params)) (*Slice[CustomerOrderPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id string, positionID string, params ...func(*Params)) (*CustomerOrderPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id string, positionID string, position *CustomerOrderPosition, params ...func(*Params)) (*CustomerOrderPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id string, position *CustomerOrderPosition, params ...func(*Params)) (*CustomerOrderPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id string, positions ...*CustomerOrderPosition) (*Slice[CustomerOrderPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id string, positionID string) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id string, positions ...*CustomerOrderPosition) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id string, positionID string) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id string) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id string, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id string) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id string) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id string, publicationID string) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон (CustomTemplate или EmbeddedTemplate)
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id string, template TemplateConverter) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает «true» в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id string, publicationID string) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID string) (*CustomerOrder, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID string) (bool, *resty.Response, error)

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

	// GetEmbeddedTemplateList выполняет запрос на получение списка встроенных шаблонов заказа покупателя.
	// Принимает контекст.
	// Возвращает объект List.
	GetEmbeddedTemplateList(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)

	// GetEmbeddedTemplateByID выполняет запрос на получение отдельного встроенного шаблона заказа покупателя по ID.
	// Принимает контекст и ID встроенного шаблона заказа покупателя.
	// Возвращает найденный встроенный шаблон заказа покупателя.
	GetEmbeddedTemplateByID(ctx context.Context, id string) (*EmbeddedTemplate, *resty.Response, error)

	// GetCustomTemplateList выполняет запрос на получение списка пользовательских шаблонов заказа покупателя.
	// Принимает контекст.
	// Возвращает объект List.
	GetCustomTemplateList(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)

	// GetCustomTemplateByID выполняет запрос на получение отдельного пользовательского шаблона заказа покупателя по ID.
	// Принимает контекст и ID пользовательского шаблона заказа покупателя.
	// Возвращает найденный пользовательский шаблон заказа покупателя.
	GetCustomTemplateByID(ctx context.Context, id string) (*CustomTemplate, *resty.Response, error)

	// GetStateByID выполняет запрос на получение статуса документа по ID.
	// Принимает контекст и ID статуса.
	// Возвращает найденный статус.
	GetStateByID(ctx context.Context, id string) (*State, *resty.Response, error)

	// CreateState выполняет запрос на создание статуса документа.
	// Принимает контекст и статус.
	// Возвращает созданный статус.
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)

	// UpdateState выполняет запрос на изменение статуса документа.
	// Принимает контекст, ID статуса и статус.
	// Возвращает изменённый статус.
	UpdateState(ctx context.Context, id string, state *State) (*State, *resty.Response, error)

	// CreateUpdateStateMany выполняет запрос на массовое создание и/или изменение статусов документа.
	// Принимает контекст и множество статусов.
	// Возвращает список созданных и/или изменённых статусов.
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)

	// DeleteState выполняет запрос на удаление статуса документа.
	// Принимает контекст и ID статуса.
	// Возвращает «true» в случае успешного удаления статуса.
	DeleteState(ctx context.Context, id string) (bool, *resty.Response, error)

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

	// GetNoteList выполняет запрос на получение ленты событий.
	// Принимает контекст и ID заказа покупателя.
	// Возвращает объект List.
	GetNoteList(ctx context.Context, id string) (*List[EventNote], *resty.Response, error)

	// GetNoteByID выполняет запрос на получение события по ID.
	// Принимает контекст, ID заказа покупателя и ID события.
	// Возвращает найденное событие.
	GetNoteByID(ctx context.Context, id string, noteID string) (*EventNote, *resty.Response, error)

	// AddNote выполняет запрос на добавление события Заказа покупателя.
	// Принимает контекст, ID заказа покупателя и текст события.
	// Возвращает объект [EventNote].
	AddNote(ctx context.Context, id string, description string) (*EventNote, *resty.Response, error)

	// UpdateNote выполняет запрос на обновление события Заказа покупателя.
	// Принимает контекст ID заказа покупателя, ID события и текст события.
	// Возвращает объект [EventNote].
	UpdateNote(ctx context.Context, id string, noteID string, description string) (*EventNote, *resty.Response, error)

	// DeleteNoteByID выполняет запрос на удаление события Заказа покупателя по ID.
	// Принимает контекст, ID заказа покупателя и ID события.
	// Возвращает «true» в случае успешного удаления события.
	DeleteNoteByID(ctx context.Context, id string, noteID string) (bool, *resty.Response, error)

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, entity *CustomerOrder, evaluate ...Evaluate) (*CustomerOrder, *resty.Response, error)
}

const (
	EndpointCustomerOrder        = EndpointEntity + string(MetaTypeCustomerOrder)
	EndpointCustomerOrderNotes   = EndpointCustomerOrder + "/%s/notes"
	EndpointCustomerOrderNotesID = EndpointCustomerOrderNotes + "/%s"
)

func (service *customerOrderService) GetNoteList(ctx context.Context, id string) (*List[EventNote], *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomerOrderNotes, id)
	return NewRequestBuilder[List[EventNote]](service.client, path).Get(ctx)
}

func (service *customerOrderService) GetNoteByID(ctx context.Context, id string, noteID string) (*EventNote, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomerOrderNotesID, id, noteID)
	return NewRequestBuilder[EventNote](service.client, path).Get(ctx)
}

func (service *customerOrderService) AddNote(ctx context.Context, id string, description string) (*EventNote, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomerOrderNotes, id)
	note := EventNote{Description: description}

	return NewRequestBuilder[EventNote](service.client, path).Post(ctx, note)
}

func (service *customerOrderService) UpdateNote(ctx context.Context, id string, noteID string, description string) (*EventNote, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomerOrderNotesID, id, noteID)
	note := EventNote{Description: description}

	return NewRequestBuilder[EventNote](service.client, path).Put(ctx, note)
}

func (service *customerOrderService) DeleteNoteByID(ctx context.Context, id string, noteID string) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomerOrderNotesID, id, noteID)
	_, resp, err := NewRequestBuilder[EventNote](service.client, path).Delete(ctx)

	return resp.StatusCode() == http.StatusNoContent, resp, err
}

// NewCustomerOrderService принимает [Client] и возвращает сервис для работы с заказами покупателя.
func NewCustomerOrderService(client *Client) CustomerOrderService {
	e := NewEndpoint(client, EndpointCustomerOrder)
	return &customerOrderService{
		endpointGetList:          endpointGetList[CustomerOrder]{e},
		endpointCreate:           endpointCreate[CustomerOrder]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[CustomerOrder]{e},
		endpointGetByID:          endpointGetByID[CustomerOrder]{e},
		endpointUpdate:           endpointUpdate[CustomerOrder]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesStatesSharedWrapper]{e},
		endpointPositions:        endpointPositions[CustomerOrderPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointPublication:      endpointPublication{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointPrintTemplates:   endpointPrintTemplates{e},
		endpointSyncID:           endpointSyncID[CustomerOrder]{e},
		endpointDeleteByID:       endpointDeleteByID{e},
		endpointDelete:           endpointDelete[CustomerOrder]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointDeleteMany:       endpointDeleteMany[CustomerOrder]{e},
	}
}
