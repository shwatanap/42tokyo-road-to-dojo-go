package response

type RankingListResponse struct {
	Ranks []User `json:"ranks"`
}

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	HighScore int    `json:"high_score"`
	Coin      int    `json:"coin"`
}
