package services

import (
	"time"

	"go-bootcamp/models"
	"go-bootcamp/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.Report, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	return s.repo.GetReport(startOfDay, endOfDay)
}

func (s *ReportService) GetReportByDateRange(startDate, endDate time.Time) (*models.Report, error) {
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	return s.repo.GetReport(startDate, endDate.Add(1*time.Second))
}
