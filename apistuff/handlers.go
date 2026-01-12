
package main

import (
	"context"
	// "fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// Define response wrappers to ensure Huma serializes to the JSON body

// STILL NEED to do a couple moire


type AttemptResponseBody struct {
	Body Attempt
}

type SearchAttemptsResponseBody struct {
	Body struct {
		Attempts []Attempt `json:"attempts"`
	}
}

func RegisterRoutes(api huma.API) {
	// --- POST: Create Attempt ---
	huma.Post(api, "/attempts", func(ctx context.Context, input *CreateAttemptInput) (*AttemptResponseBody, error) {
		var attempt Attempt
		err := DB.QueryRow(ctx,
			`INSERT INTO attempts (ip, username, password, notes, created_at)
			 VALUES ($1, $2, $3, $4, $5)
			 RETURNING id, ip, username, password, notes, created_at`,
			input.Body.IP, input.Body.Username, input.Body.Password, input.Body.Notes, time.Now(),
		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
		
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to insert attempt", err)
		}
		return &AttemptResponseBody{Body: attempt}, nil
	})

	// --- GET: Single Attempt ---
	huma.Get(api, "/attempts/{id}", func(ctx context.Context, input *GetAttemptInput) (*AttemptResponseBody, error) {
		var attempt Attempt
		err := DB.QueryRow(ctx,
			`SELECT id, ip, username, password, notes, created_at
			 FROM attempts WHERE id = $1`, input.ID,
		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
		
		if err != nil {
			return nil, huma.Error404NotFound("attempt not found", err)
		}
		return &AttemptResponseBody{Body: attempt}, nil
	})

	// --- PUT: Update Attempt ---
	huma.Put(api, "/attempts/{id}", func(ctx context.Context, input *UpdateAttemptInput) (*AttemptResponseBody, error) {
		var attempt Attempt
		err := DB.QueryRow(ctx,
			`UPDATE attempts SET ip=$1, username=$2, password=$3, notes=$4
			 WHERE id=$5
			 RETURNING id, ip, username, password, notes, created_at`,
			input.Body.IP, input.Body.Username, input.Body.Password, input.Body.Notes, input.ID,
		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
		
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to update attempt", err)
		}
		return &AttemptResponseBody{Body: attempt}, nil
	})

	// --- DELETE: Remove Attempt ---
	huma.Delete(api, "/attempts/{id}", func(ctx context.Context, input *DeleteAttemptInput) (*struct{}, error) {
		_, err := DB.Exec(ctx, `DELETE FROM attempts WHERE id=$1`, input.ID)
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to delete attempt", err)
		}
		return nil, nil // Returns 204 No Content
	})

	// --- GET: Search/List Attempts ---
	huma.Get(api, "/attempts", func(ctx context.Context, input *SearchAttemptsInput) (*SearchAttemptsResponseBody, error) {
		rows, err := DB.Query(ctx,
			`SELECT id, ip, username, password, notes, created_at
			 FROM attempts
			 WHERE created_at BETWEEN $1 AND $2
			 ORDER BY created_at`,
			input.From, input.To,
		)
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to search attempts", err)
		}
		defer rows.Close()

		var attempts []Attempt
		for rows.Next() {
			var a Attempt
			if err := rows.Scan(&a.ID, &a.IP, &a.Username, &a.Password, &a.Notes, &a.CreatedAt); err != nil {
				return nil, err
			}
			attempts = append(attempts, a)
		}

		resp := &SearchAttemptsResponseBody{}
		resp.Body.Attempts = attempts
		return resp, nil
	})
}






// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/danielgtaylor/huma/v2"
// )

// func RegisterRoutes(api huma.API) {
// 	huma.Post(api, "/attempts", func(ctx context.Context, input *CreateAttemptInput) (*Attempt, error) {
// 		var attempt Attempt
// 		err := DB.QueryRow(ctx,
// 			`INSERT INTO attempts (ip, username, password, notes, created_at)
// 			 VALUES ($1, $2, $3, $4, $5)
// 			 RETURNING id, ip, username, password, notes, created_at`,
// 			input.Body.IP,
// 			input.Body.Username,
// 			input.Body.Password,
// 			input.Body.Notes,
// 			time.Now(),
// 		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to insert attempt: %w", err)
// 		}
// 		return &attempt, nil
// 	})

// 	huma.Get(api, "/attempts/{id}", func(ctx context.Context, input *GetAttemptInput) (*Attempt, error) {
// 		var attempt Attempt
// 		err := DB.QueryRow(ctx,
// 			`SELECT id, ip, username, password, notes, created_at
// 			 FROM attempts
// 			 WHERE id = $1`, input.ID,
// 		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
// 		if err != nil {
// 			return nil, fmt.Errorf("attempt not found: %w", err)
// 		}
// 		return &attempt, nil
// 	})

// 	huma.Put(api, "/attempts/{id}", func(ctx context.Context, input *UpdateAttemptInput) (*Attempt, error) {
// 		var attempt Attempt
// 		err := DB.QueryRow(ctx,
// 			`UPDATE attempts SET ip=$1, username=$2, password=$3, notes=$4
// 			 WHERE id=$5
// 			 RETURNING id, ip, username, password, notes, created_at`,
// 			input.Body.IP,
// 			input.Body.Username,
// 			input.Body.Password,
// 			input.Body.Notes,
// 			input.ID,
// 		).Scan(&attempt.ID, &attempt.IP, &attempt.Username, &attempt.Password, &attempt.Notes, &attempt.CreatedAt)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to update attempt: %w", err)
// 		}
// 		return &attempt, nil
// 	})
// 	huma.Delete(api, "/attempts/{id}", func(ctx context.Context, input *DeleteAttemptInput) (*DeleteResponse, error) {
//     _, err := DB.Exec(ctx, `DELETE FROM attempts WHERE id=$1`, input.ID)
//     if err != nil {
//         return nil, fmt.Errorf("failed to delete attempt: %w", err)
//     }

//     return &DeleteResponse{
//         Status: 200,              // HTTP 200 OK
//         Msg:    "deleted",
//     }, nil
// })

// huma.Get(api, "/attempts", func(ctx context.Context, input *SearchAttemptsInput) (*SearchAttemptsResponse, error) {
// 	rows, err := DB.Query(ctx,
// 		`SELECT id, ip, username, password, notes, created_at
// 		 FROM attempts
// 		 WHERE created_at BETWEEN $1 AND $2
// 		 ORDER BY created_at`,
// 		input.From,
// 		input.To,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to search attempts: %w", err)
// 	}
// 	defer rows.Close()

// 	var attempts []Attempt
// 	for rows.Next() {
// 		var a Attempt
// 		if err := rows.Scan(&a.ID, &a.IP, &a.Username, &a.Password, &a.Notes, &a.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		attempts = append(attempts, a)
// 	}

// 	return &SearchAttemptsResponse{Attempts: attempts}, nil
// })

// }
