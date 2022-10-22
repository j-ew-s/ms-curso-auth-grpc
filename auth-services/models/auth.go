package authModels

import "time"

type Authorization struct {
	Username   string
	LoginLimit time.Time
}
