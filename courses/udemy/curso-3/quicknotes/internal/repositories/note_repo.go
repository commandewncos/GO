package repositories

import (
	"context"
	"math/big"
	"time"

	"github.com/Wilson-cmd/quicknotes/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
	GetById(id int) (*models.Note, error)
	CreateNote(title, content, color string) (*models.Note, error)
	UpdateNote(id int, title, content, color string) (*models.Note, error)
	Delete(id int) error
}

func NewNoteRepository(dbpool *pgxpool.Pool) NoteRepository {
	return &noteRepository{db: dbpool}

}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) List() ([]models.Note, error) {
	sql := `SELECT * FROM notes;`
	var list []models.Note
	rows, err := nr.db.Query(context.Background(), sql)

	if err != nil {
		return list, newRepositoryError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt)

		if err != nil {
			return list, newRepositoryError(err)
		}

		list = append(list, note)
	}

	return list, nil
}

func (nr *noteRepository) GetById(id int) (*models.Note, error) {
	query := `SELECT * FROM notes WHERE id = $1`
	row := nr.db.QueryRow(context.Background(), query, id)
	var note models.Note

	e := row.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt)

	if e != nil {
		return &note, newRepositoryError(e)
	} else {

		return &note, nil
	}

}

func (nr *noteRepository) CreateNote(title, content, color string) (*models.Note, error) {
	var (
		note  models.Note
		query string
	)

	query = `INSERT INTO notes (title, content, color)
			VALUES ($1, $2, $3)
			RETURNING id, created_at;`

	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Color = pgtype.Text{String: color, Valid: true}
	row := nr.db.QueryRow(context.Background(), query, note.Title, note.Content, note.Color)

	e := row.Scan(&note.Id, &note.CreatedAt)
	if e != nil {
		return &note, newRepositoryError(e)

	} else {

		return &note, nil
	}
}

func (nr *noteRepository) UpdateNote(id int, title, content, color string) (*models.Note, error) {

	var (
		note  models.Note
		query string
	)

	if len(title) > 0 {
		note.Title = pgtype.Text{String: title, Valid: true}
	}
	if len(content) > 0 {
		note.Content = pgtype.Text{String: content, Valid: true}
	}

	if len(color) > 0 {
		note.Color = pgtype.Text{String: color, Valid: true}
	}

	note.UpdatedAt = pgtype.Date{Time: time.Now(), Valid: true}
	note.Id = pgtype.Numeric{Int: big.NewInt(int64(id)), Valid: true}
	query = `
			UPDATE notes 
			
			SET 
				title = COALESCE($1, title),
				content = COALESCE($2, content),
				color = COALESCE($3, color),
				updated_at = COALESCE($4, updated_at)
			
			WHERE id = $5`

	row := nr.db.QueryRow(context.Background(), query, note.Title, note.Content, note.Color, note.UpdatedAt, note.Id)

	e := row.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.UpdatedAt)

	if e != nil {
		return &note, newRepositoryError(e)
	} else {

		return &note, nil
	}

}

func (nr *noteRepository) Delete(id int) error {

	query := `DELETE FROM notes WHERE id = $1`

	_, e := nr.db.Exec(context.Background(), query, id)

	if e != nil {
		return newRepositoryError(e)
	}

	return nil

}
