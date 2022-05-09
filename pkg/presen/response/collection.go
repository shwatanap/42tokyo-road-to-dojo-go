package response

type CollectionListResponse struct {
	Collections []CollectionItem `json:"collections"`
}

type CollectionItem struct {
	CollectionID string `json:"collection_id"`
	Name         string `json:"name"`
	Rarity       int    `json:"rarity"`
	HasItem      bool   `json:"has_item"`
}
