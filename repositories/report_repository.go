package repositories

import (
	"database/sql"
	"time"

	"go-bootcamp/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport(startDate, endDate time.Time) (*models.Report, error) {
	report := &models.Report{}

	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	bestProductQuery := `
		SELECT 
			p.name,
			COALESCE(SUM(td.quantity), 0) as qty_terjual
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`
	
	var bestProduct models.BestProduct
	err = repo.db.QueryRow(bestProductQuery, startDate, endDate).Scan(&bestProduct.Nama, &bestProduct.QtyTerjual)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	
	if err != sql.ErrNoRows {
		report.ProdukTerlaris = &bestProduct
	}

	return report, nil
}
