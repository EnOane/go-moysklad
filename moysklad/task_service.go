package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// TaskService
// Сервис для работы с задачами.
type TaskService interface {
	GetList(ctx context.Context, params *Params) (*List[Task], *resty.Response, error)
	Create(ctx context.Context, task *Task, params *Params) (*Task, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taskList []*Task, params *Params) (*[]Task, *resty.Response, error)
	DeleteMany(ctx context.Context, taskList []*Task) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Task, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, task *Task, params *Params) (*Task, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetNotes(ctx context.Context, taskId *uuid.UUID, params *Params) (*List[TaskNote], *resty.Response, error)
	CreateNote(ctx context.Context, taskId *uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	CreateNotes(ctx context.Context, taskId *uuid.UUID, taskNotes []*TaskNote) (*[]TaskNote, *resty.Response, error)
	GetNoteById(ctx context.Context, taskId, taskNoteId *uuid.UUID) (*TaskNote, *resty.Response, error)
	UpdateNote(ctx context.Context, taskId, taskNoteId *uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	DeleteNote(ctx context.Context, taskId, taskNoteId *uuid.UUID) (bool, *resty.Response, error)
}

type taskService struct {
	Endpoint
	endpointGetList[Task]
	endpointCreate[Task]
	endpointCreateUpdateMany[Task]
	endpointDeleteMany[Task]
	endpointDelete
	endpointGetById[Task]
	endpointUpdate[Task]
	endpointNamedFilter
}

func NewTaskService(client *Client) TaskService {
	e := NewEndpoint(client, "entity/task")
	return &taskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Task]{e},
		endpointCreate:           endpointCreate[Task]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Task]{e},
		endpointDeleteMany:       endpointDeleteMany[Task]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[Task]{e},
		endpointUpdate:           endpointUpdate[Task]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetNotes Запрос на получение списка всех комментариев данной Задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarii-zadachi
func (s *taskService) GetNotes(ctx context.Context, taskId *uuid.UUID, params *Params) (*List[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, taskId)
	return NewRequestBuilder[List[TaskNote]](s.client, path).SetParams(params).Get(ctx)
}

// CreateNote Запрос на создание нового комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-sozdat-kommentarij-zadachi
func (s *taskService) CreateNote(ctx context.Context, taskId *uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, taskId)
	return NewRequestBuilder[TaskNote](s.client, path).Post(ctx, taskNote)
}

// CreateNotes Запрос на создание нескольких комментариев к Задаче.
func (s *taskService) CreateNotes(ctx context.Context, taskId *uuid.UUID, taskNotes []*TaskNote) (*[]TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, taskId)
	return NewRequestBuilder[[]TaskNote](s.client, path).Post(ctx, taskNotes)
}

// GetNoteById Отдельный комментарий к Задаче с указанным id комментария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarij-k-zadache
func (s *taskService) GetNoteById(ctx context.Context, taskId, taskNoteId *uuid.UUID) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, taskId, taskNoteId)
	return NewRequestBuilder[TaskNote](s.client, path).Get(ctx)
}

// UpdateNote Запрос на обновление отдельного комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-izmenit-kommentarij-k-zadache
func (s *taskService) UpdateNote(ctx context.Context, taskId, taskNoteId *uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, taskId, taskNoteId)
	return NewRequestBuilder[TaskNote](s.client, path).Put(ctx, taskNote)
}

// DeleteNote Запрос на удаление отдельного комментария к Задаче с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-udalit-kommentarij
func (s *taskService) DeleteNote(ctx context.Context, taskId, taskNoteId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskId, taskNoteId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
