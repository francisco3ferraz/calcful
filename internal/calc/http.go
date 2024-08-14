package calc

import (
	"log/slog"
	"net/http"

	"github.com/francisco3ferraz/calcful/pkg/jsonutil"
)

type inputForm struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func AddHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		input, err := jsonutil.Decode[inputForm](r)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		result := add(input.A, input.B)

		if err := jsonutil.Encode(w, r, http.StatusOK, result); err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		logger.Info("Added two numbers", slog.Float64("a", input.A), slog.Float64("b", input.B), slog.Float64("result", result))
	}
}

func SubHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		input, err := jsonutil.Decode[inputForm](r)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		result := sub(input.A, input.B)

		if err := jsonutil.Encode(w, r, http.StatusOK, result); err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		logger.Info("Subtracted two numbers", slog.Float64("a", input.A), slog.Float64("b", input.B), slog.Float64("result", result))
	}
}

func MulHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		input, err := jsonutil.Decode[inputForm](r)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		result := mul(input.A, input.B)

		if err := jsonutil.Encode(w, r, http.StatusOK, result); err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		logger.Info("Multiplied two numbers", slog.Float64("a", input.A), slog.Float64("b", input.B), slog.Float64("result", result))
	}
}

func DivHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		input, err := jsonutil.Decode[inputForm](r)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		result := div(input.A, input.B)

		if err := jsonutil.Encode(w, r, http.StatusOK, result); err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		logger.Info("Divided two numbers", slog.Float64("a", input.A), slog.Float64("b", input.B), slog.Float64("result", result))
	}
}
