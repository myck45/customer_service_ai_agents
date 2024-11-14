package request

type SemanticSearchReq struct {
	Query               string  `json:"query" binding:"required,min=1,max=255"`
	SimilarityThreshold float32 `json:"similarity_threshold" binding:"required,min=0,max=1"`
	MatchCount          int     `json:"match_count" binding:"required,min=1"`
	RestaurantID        uint    `json:"restaurant_id" binding:"required"`
}
