package test

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"test/api/test"
	"test/internal/entity"
	"test/internal/metrics"
	"time"
)

type GetRatesApiResponse struct {
	Timestamp int64 `json:"timestamp"`
	Asks      []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Amount string `json:"amount"`
		Factor string `json:"factor"`
		Type   string `json:"type"`
	} `json:"asks"`
	Bids []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Amount string `json:"amount"`
		Factor string `json:"factor"`
		Type   string `json:"type"`
	} `json:"bids"`
}

func (i *Implementation) GetRates(ctx context.Context, in *test.GetRatesRequest) (*test.GetRatesResponse, error) {
	var err error
	defer func() {
		if err != nil {
			metrics.UnsuccessfulRequests.Inc()
		}
	}()

	resp, err := i.depthApi.GetV2Depth(ctx, in.Market)
	if err != nil {
		zap.L().Error("Get Depth API failed", zap.Error(err))
		return nil, err
	}

	unmarshalledResponse := &GetRatesApiResponse{}
	if err = json.Unmarshal(resp, unmarshalledResponse); err != nil {
		zap.L().Error("Unmarshal Depth API response failed", zap.Error(err))
		return nil, err
	}

	rate := &entity.Rate{
		Market:    in.Market,
		Ask:       0,
		Bid:       0,
		CreatedAt: time.Unix(unmarshalledResponse.Timestamp, 0),
	}

	if len(unmarshalledResponse.Asks) > 0 {
		rate.Ask, err = strconv.ParseFloat(unmarshalledResponse.Asks[0].Price, 64)
	}
	if len(unmarshalledResponse.Bids) > 0 {
		rate.Bid, err = strconv.ParseFloat(unmarshalledResponse.Bids[0].Price, 64)
	}

	if err = i.rq.Insert(ctx, rate); err != nil {
		zap.L().Error("Insert rate failed", zap.Error(err))
		return nil, err
	}

	metrics.SuccessfulRequests.Inc()

	return &test.GetRatesResponse{
		Ask:       rate.Ask,
		Bid:       rate.Bid,
		CreatedAt: timestamppb.New(rate.CreatedAt),
	}, nil
}
