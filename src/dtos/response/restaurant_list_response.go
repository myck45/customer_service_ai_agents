package response

type RestaurantListResponse struct {
	Restaurants []RestaurantResponse `json:"restaurants"`
}
