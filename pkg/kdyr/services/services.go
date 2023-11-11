package services

import pg "github.com/goshathebusiness/kirleg-do-you-remember/pkg/db/postgres"

func NewServices(db *pg.DB) *Services {
	return &Services{
		SentenceSvc: NewSentenceService(db),
	}
}
