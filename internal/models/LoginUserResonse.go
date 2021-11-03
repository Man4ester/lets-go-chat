package models

type LoginUserResponse struct {
	Url string `json:"url" description:"A url for websoket API with a one-time token for starting chat"`
}
