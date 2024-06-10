package main

import (
	"fmt"
	"os"
	"text/template"

	"Report_Generation_System/data"
	"Report_Generation_System/funcs"
)

func loadTemplates() (*template.Template, error) {
	tmpl := template.New("").Funcs(funcs.FuncMap)
	tmpl, err := tmpl.ParseGlob("templates/*.tmpl")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func generateReport(tmpl *template.Template, data interface{}, templateName, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.ExecuteTemplate(file, templateName, data)
}

func main() {
	tmpl, err := loadTemplates()
	if err != nil {
		fmt.Println("Error loading templates:", err)
		return
	}

	salesData := data.SalesData{
		Month:      "June 2024",
		TotalSales: 12345.67,
		SalesByRegion: map[string]float64{
			"North": 4567.89,
			"South": 6789.01,
		},
	}

	employees := []data.EmployeePerformance{
		{"Alice", 7890.12, 7000},
		{"Bob", 6543.21, 8000},
	}

	inventory := []data.InventoryItem{
		{"Widget", 150, 200},
		{"Gadget", 80, 50},
	}

	err = generateReport(tmpl, salesData, "SalesReport", "reports/monthly_sales_report.txt")
	if err != nil {
		fmt.Println("Error generating sales report:", err)
	}

	err = generateReport(tmpl, employees, "PerformanceReport", "reports/employee_performance_report.txt")
	if err != nil {
		fmt.Println("Error generating performance report:", err)
	}

	err = generateReport(tmpl, inventory, "InventoryReport", "reports/inventory_report.txt")
	if err != nil {
		fmt.Println("Error generating inventory report:", err)
	}
}
