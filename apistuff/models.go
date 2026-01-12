package main

import "time"


// Structs in here 
// Attempt represents a login attempt
type Attempt struct {
	ID        int       `json:"id"`
	IP        string    `json:"ip"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

// Input types for Huma
type CreateAttemptInput struct {
	Body struct {
		IP       string `json:"ip"`
		Username string `json:"username"`
		Password string `json:"password"`
		Notes    string `json:"notes"`
	} `json:"body"`
}

type UpdateAttemptInput struct {
	ID   int `path:"id"`
	Body struct {
		IP       string `json:"ip"`
		Username string `json:"username"`
		Password string `json:"password"`
		Notes    string `json:"notes"`
	} `json:"body"`
}

type GetAttemptInput struct {
	ID int `path:"id"`
}

type DeleteAttemptInput struct {
	ID int `path:"id"`
}


type DeleteResponse struct {
	Status int    `json:"status"`          // HTTP status code
	Msg    string `json:"message,omitempty"` // optional human-readable message
}

type SearchAttemptsInput struct {
	From time.Time `query:"from" required:"true"`
	To   time.Time `query:"to" required:"true"`
}

type SearchAttemptsResponse struct {
	Attempts []Attempt `json:"attempts"`
}