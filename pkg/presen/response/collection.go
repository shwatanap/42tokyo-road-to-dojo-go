package response

type CollectionListResponse struct {
	Collections []CollectionType `json:"collections"`
}

type CollectionType struct {
	CollectionID string `json:"collection_id"`
	Name         string `json:"name"`
	Rarity       int    `json:"rarity"`
	HasItem      bool   `json:"has_item"`
}
