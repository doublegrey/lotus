package telegram

type Settings struct {
	Name              string   `json:"name" binding:"required"`
	Token             string   `json:"token" binding:"required"`
	SubscriptionToken string   `json:"subscriptionToken"`
	WhiteList         []string `json:"whitelist"`
	Users             []User   `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Chat     int    `json:"chat"`
	Apps     []struct {
		ID    string `json:"id"`
		Level string `json:"level"`
	} `json:"apps"`
}
