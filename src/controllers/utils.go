package controllers

type wsMessage struct {
	Namespace string `json:"namespace"`
	Url       string `json:"url"`
}
