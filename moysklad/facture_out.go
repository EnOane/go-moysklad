package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureOut Счет-фактура выданный.
// Ключевое слово: factureout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	Organization    *Organization   `json:"organization,omitempty"`
	Deleted         *Timestamp      `json:"deleted,omitempty"`
	Applicable      *bool           `json:"applicable,omitempty"`
	AccountID       *uuid.UUID      `json:"accountId,omitempty"`
	Code            *string         `json:"code,omitempty"`
	Contract        *Contract       `json:"contract,omitempty"`
	Created         *Timestamp      `json:"created,omitempty"`
	Owner           *Employee       `json:"owner,omitempty"`
	Description     *string         `json:"description,omitempty"`
	ExternalCode    *string         `json:"externalCode,omitempty"`
	Files           *Files          `json:"files,omitempty"`
	Group           *Group          `json:"group,omitempty"`
	ID              *uuid.UUID      `json:"id,omitempty"`
	Printed         *bool           `json:"printed,omitempty"`
	Moment          *Timestamp      `json:"moment,omitempty"`
	Name            *string         `json:"name,omitempty"`
	PaymentDate     *Timestamp      `json:"paymentDate,omitempty"`
	Agent           *Counterparty   `json:"agent,omitempty"`
	Meta            *Meta           `json:"meta,omitempty"`
	Published       *bool           `json:"published,omitempty"`
	Rate            *Rate           `json:"rate,omitempty"`
	Shared          *bool           `json:"shared,omitempty"`
	State           *State          `json:"state,omitempty"`
	StateContractID *string         `json:"stateContractId,omitempty"`
	Sum             *float64        `json:"sum,omitempty"`
	SyncID          *uuid.UUID      `json:"syncId,omitempty"`
	Updated         *Timestamp      `json:"updated,omitempty"`
	Demands         Demands         `json:"demands,omitempty"`
	Payments        Payments        `json:"payments,omitempty"`
	Returns         PurchaseReturns `json:"returns,omitempty"`
	Consignee       *Counterparty   `json:"consignee,omitempty"`
	PaymentNumber   *string         `json:"paymentNumber,omitempty"`
	Attributes      Attributes      `json:"attributes,omitempty"`
}

func (factureOut FactureOut) GetOrganization() Organization {
	return Deref(factureOut.Organization)
}

func (factureOut FactureOut) GetDeleted() Timestamp {
	return Deref(factureOut.Deleted)
}

func (factureOut FactureOut) GetApplicable() bool {
	return Deref(factureOut.Applicable)
}

func (factureOut FactureOut) GetAccountID() uuid.UUID {
	return Deref(factureOut.AccountID)
}

func (factureOut FactureOut) GetCode() string {
	return Deref(factureOut.Code)
}

func (factureOut FactureOut) GetContract() Contract {
	return Deref(factureOut.Contract)
}

func (factureOut FactureOut) GetCreated() Timestamp {
	return Deref(factureOut.Created)
}

func (factureOut FactureOut) GetOwner() Employee {
	return Deref(factureOut.Owner)
}

func (factureOut FactureOut) GetDescription() string {
	return Deref(factureOut.Description)
}

func (factureOut FactureOut) GetExternalCode() string {
	return Deref(factureOut.ExternalCode)
}

func (factureOut FactureOut) GetFiles() Files {
	return Deref(factureOut.Files)
}

func (factureOut FactureOut) GetGroup() Group {
	return Deref(factureOut.Group)
}

func (factureOut FactureOut) GetID() uuid.UUID {
	return Deref(factureOut.ID)
}

func (factureOut FactureOut) GetPrinted() bool {
	return Deref(factureOut.Printed)
}

func (factureOut FactureOut) GetMoment() Timestamp {
	return Deref(factureOut.Moment)
}

func (factureOut FactureOut) GetName() string {
	return Deref(factureOut.Name)
}

func (factureOut FactureOut) GetPaymentDate() Timestamp {
	return Deref(factureOut.PaymentDate)
}

func (factureOut FactureOut) GetAgent() Counterparty {
	return Deref(factureOut.Agent)
}

func (factureOut FactureOut) GetMeta() Meta {
	return Deref(factureOut.Meta)
}

func (factureOut FactureOut) GetPublished() bool {
	return Deref(factureOut.Published)
}

func (factureOut FactureOut) GetRate() Rate {
	return Deref(factureOut.Rate)
}

func (factureOut FactureOut) GetShared() bool {
	return Deref(factureOut.Shared)
}

func (factureOut FactureOut) GetState() State {
	return Deref(factureOut.State)
}

func (factureOut FactureOut) GetStateContractId() string {
	return Deref(factureOut.StateContractID)
}

func (factureOut FactureOut) GetSum() float64 {
	return Deref(factureOut.Sum)
}

func (factureOut FactureOut) GetSyncID() uuid.UUID {
	return Deref(factureOut.SyncID)
}

func (factureOut FactureOut) GetUpdated() Timestamp {
	return Deref(factureOut.Updated)
}

func (factureOut FactureOut) GetDemands() Demands {
	return factureOut.Demands
}

func (factureOut FactureOut) GetPayments() Payments {
	return factureOut.Payments
}

func (factureOut FactureOut) GetReturns() PurchaseReturns {
	return factureOut.Returns
}

func (factureOut FactureOut) GetConsignee() Counterparty {
	return Deref(factureOut.Consignee)
}

func (factureOut FactureOut) GetPaymentNumber() string {
	return Deref(factureOut.PaymentNumber)
}

func (factureOut FactureOut) GetAttributes() Attributes {
	return factureOut.Attributes
}

func (factureOut *FactureOut) SetOrganization(organization *Organization) *FactureOut {
	factureOut.Organization = organization
	return factureOut
}

func (factureOut *FactureOut) SetApplicable(applicable bool) *FactureOut {
	factureOut.Applicable = &applicable
	return factureOut
}

func (factureOut *FactureOut) SetCode(code string) *FactureOut {
	factureOut.Code = &code
	return factureOut
}

func (factureOut *FactureOut) SetContract(contract *Contract) *FactureOut {
	factureOut.Contract = contract
	return factureOut
}

func (factureOut *FactureOut) SetOwner(owner *Employee) *FactureOut {
	factureOut.Owner = owner
	return factureOut
}

func (factureOut *FactureOut) SetDescription(description string) *FactureOut {
	factureOut.Description = &description
	return factureOut
}

func (factureOut *FactureOut) SetExternalCode(externalCode string) *FactureOut {
	factureOut.ExternalCode = &externalCode
	return factureOut
}

func (factureOut *FactureOut) SetFiles(files *Files) *FactureOut {
	factureOut.Files = files
	return factureOut
}

func (factureOut *FactureOut) SetGroup(group *Group) *FactureOut {
	factureOut.Group = group
	return factureOut
}

func (factureOut *FactureOut) SetMoment(moment *Timestamp) *FactureOut {
	factureOut.Moment = moment
	return factureOut
}

func (factureOut *FactureOut) SetName(name string) *FactureOut {
	factureOut.Name = &name
	return factureOut
}

func (factureOut *FactureOut) SetPaymentDate(paymentDate *Timestamp) *FactureOut {
	factureOut.PaymentDate = paymentDate
	return factureOut
}

func (factureOut *FactureOut) SetAgent(agent *Counterparty) *FactureOut {
	factureOut.Agent = agent
	return factureOut
}

func (factureOut *FactureOut) SetMeta(meta *Meta) *FactureOut {
	factureOut.Meta = meta
	return factureOut
}

func (factureOut *FactureOut) SetRate(rate *Rate) *FactureOut {
	factureOut.Rate = rate
	return factureOut
}

func (factureOut *FactureOut) SetShared(shared bool) *FactureOut {
	factureOut.Shared = &shared
	return factureOut
}

func (factureOut *FactureOut) SetState(state *State) *FactureOut {
	factureOut.State = state
	return factureOut
}

func (factureOut *FactureOut) SetStateContractID(stateContractID string) *FactureOut {
	factureOut.StateContractID = &stateContractID
	return factureOut
}

func (factureOut *FactureOut) SetSyncID(syncID *uuid.UUID) *FactureOut {
	factureOut.SyncID = syncID
	return factureOut
}

func (factureOut *FactureOut) SetDemands(demands Demands) *FactureOut {
	factureOut.Demands = demands
	return factureOut
}

func (factureOut *FactureOut) SetPayments(payments Payments) *FactureOut {
	factureOut.Payments = payments
	return factureOut
}

func (factureOut *FactureOut) SetReturns(returns PurchaseReturns) *FactureOut {
	factureOut.Returns = returns
	return factureOut
}

func (factureOut *FactureOut) SetConsignee(consignee *Counterparty) *FactureOut {
	factureOut.Consignee = consignee
	return factureOut
}

func (factureOut *FactureOut) SetPaymentNumber(paymentNumber string) *FactureOut {
	factureOut.PaymentNumber = &paymentNumber
	return factureOut
}

func (factureOut *FactureOut) SetAttributes(attributes Attributes) *FactureOut {
	factureOut.Attributes = attributes
	return factureOut
}

func (factureOut FactureOut) String() string {
	return Stringify(factureOut)
}

func (factureOut FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}

// FactureOutService
// Сервис для работы со счетами-фактурами выданными.
type FactureOutService interface {
	GetList(ctx context.Context, params *Params) (*List[FactureOut], *resty.Response, error)
	Create(ctx context.Context, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureOutList []*FactureOut, params *Params) (*[]FactureOut, *resty.Response, error)
	DeleteMany(ctx context.Context, factureOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*FactureOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*FactureOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewFactureOutService(client *Client) FactureOutService {
	e := NewEndpoint(client, "entity/factureout")
	return newMainService[FactureOut, any, MetadataAttributeSharedStates, any](e)
}
