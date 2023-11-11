package services

import (
	"context"
	"math/rand"

	pg "github.com/goshathebusiness/kirleg-do-you-remember/pkg/db/postgres"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/repository"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/models"
)

type SentenceService struct {
	db *pg.DB
}

func NewSentenceService(db *pg.DB) *SentenceService {
	return &SentenceService{
		db: db,
	}
}

func (s *SentenceService) AddSentence(ctx context.Context, sentence *models.Sentence) error {
	err := repository.NewSentenceRepository(s.db).AddSentence(ctx, sentence)

	if err != nil {
		return err
	}

	return nil
}

func (s *SentenceService) DeleteSentenceByIDs(ctx context.Context, ids []uint64) error {
	err := repository.NewSentenceRepository(s.db).DeleteSentenceByIDs(ctx, ids)

	if err != nil {
		return err
	}

	return nil
}

func getRandomElement(arr []*models.Sentence) *models.Sentence {
	if len(arr) == 0 {
		return nil
	}
	index := rand.Intn(len(arr))
	return arr[index]
}

func (s *SentenceService) GetRandomSentence(ctx context.Context) (*models.Sentence, error) {
	sentences, err := repository.NewSentenceRepository(s.db).FetchSentences(ctx)
	if err != nil {
		return nil, err
	}

	return getRandomElement(sentences), nil
}
