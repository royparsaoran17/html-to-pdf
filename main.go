package main

import (
	"fmt"
	"html-to-pdf/pdfgenerator"
)

type Company struct {
	Title       string
	Description string
	Company     string
	Contact     string
	Country     string
	Offices     []Office
}

type Office struct {
	Name     string
	Location string
}

func main() {
	r := pdfgenerator.NewRequestPdf("")

	//html template path
	templatePath := "template/sample.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	//html template data
	templateData := Company{
		Title:       "Company Profile",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Privy",
		Contact:     "(021) 193112",
		Country:     "Indonesia",
		Offices: []Office{{
			Name:     "Paris",
			Location: "Jalan Parangtriris",
		}, {
			Name:     "Jakal",
			Location: "Jalan Kaliurang",
		}},
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {

		// Generate PDF with custom arguments
		args := []string{"no-pdf-compression"}

		// Generate PDF
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
