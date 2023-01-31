package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hungtg7/api-app/app/order/config"
	"github.com/hungtg7/api-app/app/order/repo"
)

type Service struct {
	repo *repo.OrderRepo

	config *config.Base
}

func NewService(cfg *config.Base, repo *repo.OrderRepo) *Service {
	return &Service{
		config: cfg,
		repo:   repo,
	}
}

func (s *Service) Close(ctx context.Context) {
	ctx.Done()
}

func (s *Service) GetOrder(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		filters, present := query["id"] //filters=["color", "price", "brand"]
		if !present || len(filters) == 0 {
			fmt.Println("filters not present")
		}
		w.WriteHeader(200)
		hostname, _ := os.Hostname()
		filters = append(filters, hostname)
		w.Write([]byte(strings.Join(filters, ",")))
	}
}
