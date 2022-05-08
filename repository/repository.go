package repository

import "time"

type Service struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	ServiceType string    `json:"serviceType"`
	TypeDefs    string    `json:"tpeDefs"`
	IsAlive     bool      `json:"isAlive"`
	Version     string    `json:"version"`
	AddedTime   time.Time `json:"addedTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

func GetServices() []Service {
	var services = []Service{}

	return services
}
