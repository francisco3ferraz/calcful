package server

import (
	"log/slog"
	"net/http"

	"github.com/francisco3ferraz/calcful/internal/calc"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger) {
	rateLimiter := NewRateLimiter(1, 5)

	mux.HandleFunc("/add", rateLimiter.Middleware(calc.AddHandler(logger)))
	mux.HandleFunc("/sub", calc.SubHandler(logger))
	mux.HandleFunc("/mul", calc.MulHandler(logger))
	mux.HandleFunc("/div", calc.DivHandler(logger))
}
