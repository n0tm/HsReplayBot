package config

type Config struct {
	Bot struct {
		AccessToken string `json:"access_token"`
		GroupId     int    `json:"group_id"`
	}
}
