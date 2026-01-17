package skinport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
	"github.com/madkingxxx/backend-test/internal/core/skinport"
	"github.com/madkingxxx/backend-test/internal/driven/ext_http/skinport/port"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

const (
	itemsEndpoint = "/v1/items?app_id=730&currency=USD"
)

// GetAllItems fetches items from Skinport API.
func (s *Sender) GetAllItems(ctx context.Context) ([]skinport.Item, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, s.baseURL+itemsEndpoint, nil)
	if err != nil {
		utils.Logger.Error(ctx, "failed to create request", zap.Error(err))
		return nil, err
	}

	// Skinport API supports Brotli compression and recommends using it.
	// https://docs.skinport.com/items#encoding
	request.Header.Set("Accept-Encoding", "br")

	resp, err := s.client.Do(request)
	if err != nil {
		utils.Logger.Error(ctx, "failed to get all items", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body

	switch resp.Header.Get("Content-Encoding") {
	case "br":
		reader = brotli.NewReader(resp.Body)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var errorResponse port.ErrorResponse
		if err := json.NewDecoder(reader).Decode(&errorResponse); err != nil {
			utils.Logger.Error(ctx, "failed to decode error response", zap.Error(err))
			return nil, err
		}

		utils.Logger.Error(ctx, "failed to get all items", zap.String("message", errorResponse.Errors[0].Msg))
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, errorResponse.Errors[0].Msg)
	}

	var result port.Items
	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		utils.Logger.Error(ctx, "failed to decode items", zap.Error(err))
		return nil, err
	}

	return result.ToCore(), nil
}
