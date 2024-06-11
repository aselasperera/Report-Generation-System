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
		TotalSales: 45000.10,
		SalesByRegion: map[string]float64{
			"North": 15000.05,
			"South": 30000.05,
		},
	}

	employees := []data.EmployeePerformance{
		{"Asela Perera", 7000.15, 6000},
		{"John Fernando", 6000.20, 8000},
	}

	inventory := []data.InventoryItem{
		{"Laptop", 150, 200},
		{"Mobile Phone", 80, 50},
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
