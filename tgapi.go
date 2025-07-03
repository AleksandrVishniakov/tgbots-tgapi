package tgbotstgapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AleksandrVishniakov/tgbots-tgapi/dto"
)

type API struct {
	httpClient *http.Client
}

func New() *API {
	return &API{
		httpClient: &http.Client{},
	}
}

func (api *API) GetMe(ctx context.Context, token string) (*dto.GetMeResponse, error) {
	const src = "API.GetMe"

	resp, err := request[any, dto.GetMeResponse](ctx, api.httpClient, token, dto.GetMeMethod, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) GetUpdates(ctx context.Context, token string, req *dto.GetUpdatesRequest) (*dto.GetUpdatesResponse, error) {
	const src = "API.GetUpdates"

	resp, err := request[dto.GetUpdatesRequest, dto.GetUpdatesResponse](ctx, api.httpClient, token, dto.GetUpdatesMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) GetMyName(ctx context.Context, token string, req *dto.GetMyNameRequest) (*dto.GetMyNameResponse, error) {
	const src = "API.GetMyName"

	resp, err := request[dto.GetMyNameRequest, dto.GetMyNameResponse](ctx, api.httpClient, token, dto.GetMyNameMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) SetMyName(ctx context.Context, token string, req *dto.SetMyNameRequest) (*dto.SetMyNameResponse, error) {
	const src = "API.SetMyName"

	resp, err := request[dto.SetMyNameRequest, dto.SetMyNameResponse](ctx, api.httpClient, token, dto.SetMyNameMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) GetMyDescription(ctx context.Context, token string, req *dto.GetMyDescriptionRequest) (*dto.GetMyDescriptionResponse, error) {
	const src = "API.GetMyDescription"

	resp, err := request[dto.GetMyDescriptionRequest, dto.GetMyDescriptionResponse](ctx, api.httpClient, token, dto.GetMyDescriptionMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) SetMyDescription(ctx context.Context, token string, req *dto.SetMyDescriptionRequest) (*dto.SetMyDescriptionResponse, error) {
	const src = "API.SetMyDescription"

	resp, err := request[dto.SetMyDescriptionRequest, dto.SetMyDescriptionResponse](ctx, api.httpClient, token, dto.SetMyDescriptionMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) GetMyShortDescription(ctx context.Context, token string, req *dto.GetMyShortDescriptionRequest) (*dto.GetMyShortDescriptionResponse, error) {
	const src = "API.GetMyShortDescription"

	resp, err := request[dto.GetMyShortDescriptionRequest, dto.GetMyShortDescriptionResponse](ctx, api.httpClient, token, dto.GetMyShortDescriptionMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}

func (api *API) SetMyShortDescription(ctx context.Context, token string, req *dto.SetMyShortDescriptionRequest) (*dto.SetMyShortDescriptionResponse, error) {
	const src = "API.SetMyDescription"

	resp, err := request[dto.SetMyShortDescriptionRequest, dto.SetMyShortDescriptionResponse](ctx, api.httpClient, token, dto.SetMyShortDescriptionMethod, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", src, err)
	}

	return resp, nil
}
