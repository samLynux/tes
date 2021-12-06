package model

import "gorm.io/gorm"

type Movie struct {
	Title  string `json:"title" gorm:"column:title"`
	Year   string `json:"year" gorm:"column:year"`
	ImdbID string `json:"imdbid" gorm:"column:imdbid"`
	Type   string `json:"type" gorm:"column:type"`
}

type Result struct {
	Results interface{} `json:"result" gorm:"column:result"`
	Page    string      `json:"page" gorm:"column:page"`
}

//SearchResponse is the struct of the response in a search
type SearchResponse struct {
	Search       []Movie
	Response     string `json:"response" gorm:"column:response"`
	TotalResults string `json:"totalresults" gorm:"column:totalresults"`
}

type Log struct {
	gorm.Model
	SearchWord string `json:"search_word" gorm:"column:search_word"`
	Page       string `json:"page" gorm:"column:page"`
}
