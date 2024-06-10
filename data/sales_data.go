package data

type SalesData struct {
	Month         string
	TotalSales    float64
	SalesByRegion map[string]float64
}
