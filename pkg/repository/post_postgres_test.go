package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	httpapi "github.com/tsarkovmi/http_api"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestPostPostgres_CreateWorker(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewPostPostgres(db) // инициализируем репозиторий

	type args struct {
		worker httpapi.Worker
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			input: args{
				worker: httpapi.Worker{
					Name:       "John Doe",
					Age:        30,
					Salary:     3000,
					Occupation: "Engineer",
				},
			},
			want: 1,
			mock: func(args args, id int) {
				// Ожидаем корректную вставку данных и возврат ID
				mock.ExpectQuery("INSERT INTO workers").
					WithArgs(args.worker.Name, args.worker.Age, args.worker.Salary, args.worker.Occupation).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
			},
		},
		{
			name: "Insert Error",
			input: args{
				worker: httpapi.Worker{
					Name:       "Jane Doe",
					Age:        28,
					Salary:     3500,
					Occupation: "Manager",
				},
			},
			mock: func(args args, id int) {
				// Симулируем ошибку при вставке данных
				mock.ExpectQuery("INSERT INTO workers").
					WithArgs(args.worker.Name, args.worker.Age, args.worker.Salary, args.worker.Occupation).
					WillReturnError(errors.New("insert error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.CreateWorker(tt.input.worker)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// Проверка выполнения всех ожиданий mock
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestPostPostgres_GetAllWorkers(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewPostPostgres(db) // инициализируем репозиторий

	tests := []struct {
		name    string
		mock    func()
		want    []httpapi.Worker
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "salary", "occupation"}).
					AddRow(1, "John Doe", 30, 3000.0, "Engineer").
					AddRow(2, "Jane Doe", 28, 3500.0, "Manager").
					AddRow(3, "Alice", 25, 3200.0, "Developer")

				mock.ExpectQuery("SELECT \\* FROM workers").
					WillReturnRows(rows)
			},
			want: []httpapi.Worker{
				{ID: 1, Name: "John Doe", Age: 30, Salary: 3000.0, Occupation: "Engineer"},
				{ID: 2, Name: "Jane Doe", Age: 28, Salary: 3500.0, Occupation: "Manager"},
				{ID: 3, Name: "Alice", Age: 25, Salary: 3200.0, Occupation: "Developer"},
			},
			wantErr: false,
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "salary", "occupation"})
				mock.ExpectQuery("SELECT \\* FROM workers").
					WillReturnRows(rows)
			},
			want:    []httpapi.Worker(nil),
			wantErr: false,
		},
		{
			name: "Query Error",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM workers").
					WillReturnError(errors.New("query error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAllWorkers()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestPostPostgres_FindWorkerByID(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewPostPostgres(db) // инициализируем репозиторий

	tests := []struct {
		name    string
		mock    func()
		input   int
		want    httpapi.Worker
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "salary", "occupation"}).
					AddRow(1, "John Doe", 30, 3000.0, "Engineer")

				mock.ExpectQuery("SELECT \\* FROM workers WHERE id=\\$1").
					WithArgs(1).WillReturnRows(rows)
			},
			input:   1,
			want:    httpapi.Worker{ID: 1, Name: "John Doe", Age: 30, Salary: 3000.0, Occupation: "Engineer"},
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "salary", "occupation"})
				mock.ExpectQuery("SELECT \\* FROM workers WHERE id=\\$1").
					WithArgs(404).WillReturnRows(rows)
			},
			input:   404,
			want:    httpapi.Worker{},
			wantErr: true,
		},
		{
			name: "Query Error",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM workers WHERE id=\\$1").
					WithArgs(1).WillReturnError(errors.New("query error"))
			},
			input:   1,
			want:    httpapi.Worker{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.FindWorkerByID(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
