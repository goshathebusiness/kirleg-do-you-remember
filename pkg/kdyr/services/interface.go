package services

import (
	"context"

	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/models"
)

type Services struct {
	SentenceSvc SentenceSvc
}

type SentenceSvc interface {
	GetRandomSentence(ctx context.Context) (*models.Sentence, error)
	AddSentence(ctx context.Context, sentence *models.Sentence) error
	DeleteSentenceByIDs(ctx context.Context, ids []uint64) error
}
