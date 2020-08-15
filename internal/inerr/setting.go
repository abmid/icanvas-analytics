package inerr

import "errors"

var (
	ErrNoCanvasConfig error = errors.New("Please fill Canvas URL and Access Token")
)
