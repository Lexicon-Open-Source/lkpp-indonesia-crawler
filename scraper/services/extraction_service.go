package services

import (
	"context"
	"lexicon/lkpp-indonesia-crawler/common"
	"lexicon/lkpp-indonesia-crawler/scraper/models"
)

func UpsertExtraction(extraction models.Extraction) error {
	context := context.Background()

	tx, err := common.Pool.Begin(context)
	if err != nil {
		return err
	}

	err = models.UpsertExtraction(context, tx, extraction)
	if err != nil {
		return err
	}

	tx.Commit(context)

	return nil
}
