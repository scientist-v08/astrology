package models

import "time"

type BirthChartReq struct {
	BirthDateTime time.Time `json:"birthDateTime" binding:"required"`
	Latitude      float64   `json:"latitude" binding:"required,gte=-90,lte=90"`
	Longitude     float64   `json:"longitude" binding:"required,gte=-180,lte=180"`
}