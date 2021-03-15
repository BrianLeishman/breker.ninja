package main

import "time"

type User struct {
	UserID   string    `dynamo:"pk"`
	Email    string    `dynamo:"sk"`
	Password string    `dynamo:"pass"`
	AuthKey  []byte    `dynamo:"authKey"`
	Verified bool      `dynamo:"verified"`
	APIKey   string    `dynamo:"apiKey"`
	Name     string    `dynamo:"name"`
	Created  time.Time `dynamo:"created"`
}
