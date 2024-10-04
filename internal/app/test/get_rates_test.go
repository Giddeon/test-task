package test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/api/test"
	"test/internal/entity"
)

type MockDepthApiClient struct {
	mock.Mock
}

func (m *MockDepthApiClient) GetV2Depth(ctx context.Context, market string) ([]byte, error) {
	args := m.Called(ctx, market)
	return args.Get(0).([]byte), args.Error(1)
}

type MockRq struct {
	mock.Mock
}

func (m *MockRq) Insert(ctx context.Context, rate *entity.Rate) error {
	args := m.Called(ctx, rate)
	return args.Error(0)
}

func TestGetRates_Success(t *testing.T) {
	mockClient := new(MockDepthApiClient)
	mockRq := new(MockRq)

	testResponse := GetRatesApiResponse{
		Timestamp: time.Now().Unix(),
		Asks: []struct {
			Price  string `json:"price"`
			Volume string `json:"volume"`
			Amount string `json:"amount"`
			Factor string `json:"factor"`
			Type   string `json:"type"`
		}{{Price: "100.50"}},
		Bids: []struct {
			Price  string `json:"price"`
			Volume string `json:"volume"`
			Amount string `json:"amount"`
			Factor string `json:"factor"`
			Type   string `json:"type"`
		}{{Price: "99.50"}},
	}
	responseBody, _ := json.Marshal(testResponse)

	mockClient.On("GetV2Depth", mock.Anything, "BTC-USD").Return(responseBody, nil)
	mockRq.On("Insert", mock.Anything, mock.AnythingOfType("*entity.Rate")).Return(nil)

	impl := &Implementation{
		depthApi: mockClient,
		rq:       mockRq,
	}

	req := &test.GetRatesRequest{Market: "BTC-USD"}
	ctx := context.Background()

	resp, err := impl.GetRates(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 100.50, resp.Ask)
	assert.Equal(t, 99.50, resp.Bid)

	mockClient.AssertExpectations(t)
	mockRq.AssertExpectations(t)
}
