package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

type QueryParams struct {
	Group string `form:"group"`
	Song  string `form:"song"`
	Limit int    `form:"limit"`
	Offset int   `form:"offset"`
	Verse int `form:"verse"`
}