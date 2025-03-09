package models

type Request struct{
	Id uint64 `json:"id"`
	Status string `json:"status"`
}