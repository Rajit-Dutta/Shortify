package service

import (
	"context"
	"errors"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/api/helpers"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/model"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/repository"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/utils"
)

func PublishURL(req model.Request, ctx context.Context) (model.ToBePublishedURL, error) {
	var urlId string

	//Check whether URL contains domain errors
	if !helpers.RemoveDomainErrors(req.URL) {
		return model.ToBePublishedURL{}, errors.New("URL contains domain errros")
	}
	req.URL = helpers.EnforceHTTP(req.URL)

	//validate URL
	if err := utils.ValidateURL(req.URL); err != nil {
		return model.ToBePublishedURL{}, errors.New("URL is not valid")
	}

	if req.Expiry == 0 {
		req.Expiry = 24
	}

	//Generate custom short
	if req.CustomShort == "" {
		urlId = utils.GenerateCustomShort(req.URL)
		err := repository.IsExistCustomShortURL(ctx, urlId)
		if err != nil {
			return model.ToBePublishedURL{}, err
		}
	} else {
		urlId = req.CustomShort
		//Check whether custom short already exists
		err := repository.IsExistCustomShortURL(ctx, urlId)
		if err != nil {
			return model.ToBePublishedURL{}, err
		}
	}

	//Create shortened URL
	shortenedURL := config.MustLoad().Domain + "/" + urlId
	//Set value URL with key id
	err := repository.SetShortenedURL(ctx, urlId, req.URL, req.Expiry)
	if err != nil {
		return model.ToBePublishedURL{}, err
	}
	return model.ToBePublishedURL{
		OriginalURL:  req.URL,
		ShortenedURL: shortenedURL,
		Expiry:       req.Expiry,
	}, nil
}
