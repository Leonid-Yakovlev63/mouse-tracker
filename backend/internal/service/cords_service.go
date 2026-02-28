package service

import (
	"github.com/Leonid-Yakovlev63/mouse-tracker/backend/internal/domain"
)

type (
	CordsService interface {
		ProcessCords(cords domain.Cords)
	}
	cordsService struct{}
)

func NewCordsService() CordsService {
	return &cordsService{}
}

func (s *cordsService) ProcessCords(cords domain.Cords) {
	cords.PrintCords()
}
