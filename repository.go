package main

import "time"

type Service struct {
	id          uint
	name        string
	url         string
	typeDefs    string
	isAlive     bool
	version     string
	addedTime   time.Time
	updatedTime time.Time
}

func getServices() []Service {
	var services = []Service{}

	return services
}
