package mysql

import (
	"database/sql"
	"go_notes_snippet/pkg/models"
)

// NoteModel - Определяет тип который обертывает пул подключения sql.DB
type NoteModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *NoteModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *NoteModel) Get(id int) (*models.Note, error) {
	return nil, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *NoteModel) Latest() ([]*models.Note, error) {
	return nil, nil
}
