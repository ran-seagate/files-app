package entities

import "time"

type FileDetails struct {
	Name         string    `json:"name"`
	Ext          string    `json:"ext"`
	Size         uint64    `json:"size"`
	CreationDate time.Time `json:"creation_date"`
}
