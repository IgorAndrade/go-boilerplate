package service

import (
	"context"
	"testing"

	"github.com/IgorAndrade/go-boilerplate/app/db/mock"
	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
	"github.com/globalsign/mgo"
)

func TestTodoListImp_Create(t *testing.T) {
	type fields struct {
		repository repository.TodoList
	}
	type args struct {
		ctx      context.Context
		todoList *model.TodoList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "save ok",
			args: args{
				ctx:      context.Background(),
				todoList: &model.TodoList{},
			},
			fields: fields{
				repository: &mock.TodoListMock{
					CreateFn: func(c context.Context, tl *model.TodoList) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Error when save",
			args: args{
				ctx:      context.Background(),
				todoList: &model.TodoList{},
			},
			fields: fields{
				repository: &mock.TodoListMock{
					CreateFn: func(c context.Context, tl *model.TodoList) error {
						return mgo.ErrCursor
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := NewTodoListService(tt.fields.repository)
			if err := tl.Create(tt.args.ctx, tt.args.todoList); (err != nil) != tt.wantErr {
				t.Errorf("TodoListImp.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
