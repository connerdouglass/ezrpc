package main

import "context"

type sumRequest struct {
	Numbers []int `json:"numbers"`
}

type sumResponse struct {
	Sum int `json:"sum"`
}

// Sum calculates the sum of a slice of numbers
func Sum(ctx context.Context, req *sumRequest) (*sumResponse, error) {
	var sum int
	for _, number := range req.Numbers {
		sum += number
	}
	return &sumResponse{Sum: sum}, nil
}
