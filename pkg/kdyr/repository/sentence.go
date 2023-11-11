package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"

	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/models"
)

type SentenceRepository struct {
	db sqlx.ExtContext
}

func NewSentenceRepository(db sqlx.ExtContext) *SentenceRepository {
	return &SentenceRepository{db: db}
}

func (r *SentenceRepository) AddSentence(ctx context.Context, sentence *models.Sentence) error {
	query, args, err := sq.Insert("sentences").
		Columns("text").
		Values(sentence.Text).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *SentenceRepository) DeleteSentenceByIDs(ctx context.Context, ids []uint64) error {
	query, args, err := sq.Delete("sentences").
		Where(sq.Eq{"id": ids}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *SentenceRepository) FetchSentences(ctx context.Context) ([]*models.Sentence, error) {
	query, args, err := sq.
		Select("*").
		From("sentences").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var sentences []*models.Sentence
	err = sqlx.SelectContext(ctx, r.db, &sentences, query, args...)
	if err != nil {
		return nil, err
	}

	return sentences, nil
}
