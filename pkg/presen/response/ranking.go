package response

type RankingListResponse struct {
	Ranks []UserGetResponse `json:"ranks"`
}
