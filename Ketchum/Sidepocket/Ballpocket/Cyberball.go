package Cyberball

import "github.com/jung-kurt/gofpdf"

func GeneratePDF(filename, content string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, filename)
    pdf.Ln(20) // Move below the title
    pdf.SetFont("Arial", "", 12)
    pdf.MultiCell(190, 10, content, "0", "L", false)
    err := pdf.OutputFileAndClose(filename + ".pdf")
    return err
}
