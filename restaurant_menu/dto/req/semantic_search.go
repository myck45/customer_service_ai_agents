package req

// SemanticSearchReq represents the request to perform a semantic search
// @Description perform a semantic search on the menu items
type SemanticSearchReq struct {
	Query               string  `json:"query" binding:"required,min=1,max=255" example:"sopas para el frío" extensions:"x-order=0"` // Query is the search query
	SimilarityThreshold float32 `json:"similarity_threshold" binding:"required,min=0,max=1" example:"0.5" extensions:"x-order=1"`   // SimilarityThreshold is the threshold for the similarity
	MatchCount          int     `json:"match_count" binding:"required,min=1" example:"5" extensions:"x-order=2"`                    // MatchCount is the number of matches to return
	RestaurantID        uint    `json:"restaurant_id" binding:"required" example:"1" extensions:"x-order=3"`                        // RestaurantID is the ID of the restaurant to search in
}
