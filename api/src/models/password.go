package models

type ChangePassword struct {
	Old string `json:"old"`
	New string `json:"new"`
}
