package route

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorAndrade/go-boilerplate/app/apiErrors"
	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_create(t *testing.T) {
	type args struct {
		ctn GetterDI
	}
	type fields struct {
		payload interface{}
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantErr    bool
		wantStatus int
		errType    apiErrors.ErrorType
	}{
		{
			name: "Save ok",
			fields: fields{
				payload: model.TodoList{Text: "qwerr"},
			},
			args: args{
				ctn: config.ContainerMock{
					GetFn: func(name string) interface{} {
						return &service.TodoListMock{
							CreateFn: func(c context.Context, tl *model.TodoList) error {
								return nil
							},
						}
					},
				},
			},
			wantErr:    false,
			wantStatus: http.StatusCreated,
		},
		{
			name: "Payload invalid",
			fields: fields{
				payload: "qwerr",
			},
			args: args{
				ctn: config.ContainerMock{
					GetFn: func(name string) interface{} {
						return &service.TodoListMock{
							CreateFn: func(c context.Context, tl *model.TodoList) error {
								return nil
							},
						}
					},
				},
			},
			wantErr: true,
			errType: apiErrors.BadRequest,
		},
		{
			name: "Error on create",
			fields: fields{
				payload: model.TodoList{Text: "qwerr"},
			},
			args: args{
				ctn: config.ContainerMock{
					GetFn: func(name string) interface{} {
						return &service.TodoListMock{
							CreateFn: func(c context.Context, tl *model.TodoList) error {
								return apiErrors.InternalError.New("got error")
							},
						}
					},
				},
			},
			wantErr: true,
			errType: apiErrors.InternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			b, err := json.Marshal(tt.fields.payload)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/todo-list", bytes.NewReader(b))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err = create(c, tt.args.ctn)
			if (err != nil) != tt.wantErr {
				t.Errorf("create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				assert.True(t, apiErrors.Is(err, tt.errType), fmt.Sprintf("expected type of error %v is eq %v", apiErrors.GetType(err), tt.errType))
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
			}
		})
	}
}
