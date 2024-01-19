package Cyberball

import "github.com/jung-kurt/gofpdf"

// GeneratePDF creates a PDF file with the given filename and content.
// It returns an error if the PDF generation fails.
func GeneratePDF(filename, content string) error {
	// Initialize a new PDF document.
	// "P" sets the orientation to Portrait.
	// "mm" is the unit of measurement (millimeters).
	// "A4" sets the size of the PDF.
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page to the PDF.
	pdf.AddPage()

	// Set the font for the PDF title.
	// "Arial" is the font type, "B" for bold, and 16 is the font size.
	pdf.SetFont("Arial", "B", 16)

	// Create a cell for the title and insert the filename.
	// 40 is the width of the cell, 10 is the height.
	pdf.Cell(40, 10, filename)

	// Move to the next line in the PDF, 20 units below the current line.
	pdf.Ln(20)

	// Set the font for the main content.
	// "Arial" is the font type, "" for normal style (not bold/italic), and 12 is the font size.
	pdf.SetFont("Arial", "", 12)

	// Add the main content as a multi-line cell.
	// 190 is the width, 10 is the height of each line.
	// "content" is the text to be added.
	// "0" is the border (0 for no border), "L" is the alignment (Left), and false means no fill.
	pdf.MultiCell(190, 10, content, "0", "L", false)

	// Output the PDF to a file with the given filename.
	// The filename is appended with ".pdf".
	// The function returns any error encountered during the file creation.
	err := pdf.OutputFileAndClose(filename + ".pdf")
	return err
}
