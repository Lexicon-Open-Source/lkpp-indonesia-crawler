package services

import (
	"context"
	"lexicon/lkpp-indonesia-crawler/common"
	"lexicon/lkpp-indonesia-crawler/crawler/models"
)

func UpsetUrl(urlFrontiers []models.UrlFrontier) error {
	ctx := context.Background()

	tx, err := common.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	err = models.UpsertUrlFrontier(ctx, tx, urlFrontiers)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}

func UpdateUrlFrontierStatus(urlId string, status int) error {
	ctx := context.Background()

	tx, err := common.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	err = models.UpdateUrlFrontierStatus(ctx, tx, urlId, status)
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}

func GetUnscrapedUrl() ([]models.UrlFrontier, error) {
	ctx := context.Background()

	tx, err := common.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	urls, err := models.GetUrlFrontiersUnscraped(ctx, tx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}
