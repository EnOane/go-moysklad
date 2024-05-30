package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProjectService
// Сервис для работы с проектами.
type ProjectService interface {
	GetList(ctx context.Context, params *Params) (*List[Project], *resty.Response, error)
	Create(ctx context.Context, project *Project, params *Params) (*Project, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, projectList []*Project, params *Params) (*[]Project, *resty.Response, error)
	DeleteMany(ctx context.Context, projectList []*Project) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Project, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, project *Project, params *Params) (*Project, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewProjectService(client *Client) ProjectService {
	e := NewEndpoint(client, "entity/project")
	return newMainService[Project, any, MetadataAttributeShared, any](e)
}
