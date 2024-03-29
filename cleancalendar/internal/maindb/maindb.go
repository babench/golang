package maindb

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/otusteam/go/cleancalendar/internal/domain/models"
	"time"
)

// implements domain.interfaces.EventStorage
type PgEventStorage struct {
	db *sqlx.DB
}

func NewPgEventStorage(dsn string) (*PgEventStorage, error) {
	db, err := sqlx.Open("pgx", dsn) // *sql.DB
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgEventStorage{db: db}, nil
}

func (pges *PgEventStorage) SaveEvent(ctx context.Context, event *models.Event) error {
	query := `
		INSERT INTO events(id, owner, title, text, start_time, end_time)
		VALUES (:id, :owner, :title, :text, :start_time, :end_time)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         event.Id.String(),
		"owner":      event.Owner,
		"title":      event.Title,
		"text":       event.Text,
		"start_time": event.StartTime,
		"end_time":   event.EndTime,
	})
	return err
}

func (pges *PgEventStorage) GetEventById(ctx context.Context, id string) (*models.Event, error) {
	// TODO
	return nil, nil
}

func (pges *PgEventStorage) GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*models.Event {
	// TODO
	return nil
}
