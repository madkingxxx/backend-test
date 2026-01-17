package port

import "github.com/madkingxxx/backend-test/internal/core/skinport"

// Item - external HTTP item representation for parsing JSON responses
type Item struct {
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
type Items []Item

func (irl Items) ToCore() []skinport.Item {
	var result []skinport.Item
	for _, item := range irl {
		result = append(result, skinport.Item{
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

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	ID  string `json:"id"`
	Msg string `json:"message"`
}
