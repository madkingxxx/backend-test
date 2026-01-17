package port

import "github.com/madkingxxx/backend-test/internal/core/skinport"

type ItemResponse struct {
	MarketHashName string  `json:"market_hash_name"`
	Currency       string  `json:"currency"`
	SuggestedPrice float64 `json:"suggested_price"`
	ItemPage       string  `json:"item_page"`
	MarketPage     string  `json:"market_page"`
	MinPrice       float64 `json:"min_price"`
	MaxPrice       float64 `json:"max_price"`
	MeanPrice      float64 `json:"mean_price"`
	MedianPrice    float64 `json:"median_price"`
	Quantity       int     `json:"quantity"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
}

func Convert(items []skinport.Item) []ItemResponse {
	var result []ItemResponse
	for _, item := range items {
		result = append(result, ItemResponse{
			MarketHashName: item.MarketHashName,
			Currency:       item.Currency,
			SuggestedPrice: item.SuggestedPrice,
			ItemPage:       item.ItemPage,
			MarketPage:     item.MarketPage,
			MinPrice:       item.MinPrice,
			MaxPrice:       item.MaxPrice,
			MeanPrice:      item.MeanPrice,
			MedianPrice:    item.MedianPrice,
			Quantity:       item.Quantity,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
		})
	}
	return result
}
