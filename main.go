package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

type ContractData struct {
	// Second party information (to be filled)
	SecondPartyName    string
	SecondPartyType    string
	SecondPartyAddress string
	SecondPartyRep     string
	SecondPartyTitle   string
	PharmacyName       string
	SIANumber          string

	// Contract duration
	Duration  string
	StartDate string
	EndDate   string
}

func main() {
	// Sample data - replace with actual values
	contractData := ContractData{
		SecondPartyName:    "PT Apotek Sehat Bersama",
		SecondPartyType:    "Perseroan Terbatas",
		SecondPartyAddress: "Jl. Kesehatan No. 123, Jakarta Pusat",
		SecondPartyRep:     "Dr. Ahmad Farmasi",
		SecondPartyTitle:   "Direktur",
		PharmacyName:       "Apotek Sehat",
		SIANumber:          "SIA/001/2024/DKI",
		Duration:           "3 (tiga)",
		StartDate:          "01 Agustus 2025",
		EndDate:            "01 Agustus 2028",
	}

	err := generateContractPDF(contractData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF contract generated successfully: pharmacy_contract.pdf")
}

func generateContractPDF(data ContractData) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	// pdf.SetAutoPageBreak(true, 20)
	// ==== HEADER (centered) ====
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Times", "B", 12)
		// Move cursor to top margin
		pdf.SetY(5)
		// Full width cell with centered text
		pdf.CellFormat(0, 5, "PERJANJIAN KERJASAMA", "", 2, "C", false, 0, "")
		pdf.CellFormat(0, 5, "APOTEK DARING", "", 2, "C", false, 0, "")
		pdf.CellFormat(0, 5, "No. SP/XXX/LGL-PPP/XXX/2025", "", 2, "C", false, 0, "")
		// Line under header
		pdf.Line(20, 20, 190, 20)
		// pdf.Ln(10)
	})

	// ==== FOOTER (centered page number) ====
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15) // 15 mm from bottom
		pdf.SetFont("Times", "I", 10)
		pdf.CellFormat(0, 10,
			fmt.Sprintf("Halaman %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.SetMargins(10, 20, 10)

	// Add font for Indonesian characters
	pdf.AddPage()
	pdf.SetFont("Times", "B", 14)

	// Introduction

	pdf.SetFont("Times", "", 11)
	introText := `Perjanjian Kerja Sama Apotek Daring ini (selanjutnya disebut sebagai "Perjanjian") dibuat dan ditandatangani pada hari ini Jumat, tanggal satu bulan Agustus tahun dua ribu dua puluh lima (01-08-2025), oleh dan antara :`

	pdf.MultiCell(0, 6, introText, "", "", false)
	pdf.Ln(5)

	// First Party
	pdf.SetFont("Times", "B", 11)
	pdf.Cell(10, 6, "1.")
	pdf.Cell(0, 6, "PT Perintis Pelayanan Paripurna")
	pdf.Ln(6)

	pdf.SetFont("Times", "", 10)
	firstPartyText := `   suatu perseroan terbatas yang didirikan berdasarkan dan tunduk pada hukum Negara Republik Indonesia, beralamat di Grand ITC Permata Hijau Kantor Emerald, Jl.Letjen Soepeno Arteri Permata Hijau E No.26, Grogol Utara, Kebayoran Lama, Kota Jakarta Selatan, Provinsi DKI Jakarta. Dalam hal ini diwakili oleh Yasinta Yulian Hendrata selaku Presiden Direktur, oleh karenanya sah bertindak untuk dan atas nama PT Perintis Pelayanan Paripurna (selanjutnya disebut sebagai "PIHAK PERTAMA").`

	pdf.MultiCell(0, 5, firstPartyText, "", "", false)
	pdf.Ln(3)

	pdf.SetFont("Times", "", 11)
	pdf.Cell(0, 6, "Dan")
	pdf.Ln(8)

	// Second Party
	pdf.SetFont("Times", "B", 11)
	pdf.Cell(10, 6, "2.")
	pdf.Cell(0, 6, data.SecondPartyName)
	pdf.Ln(6)

	pdf.SetFont("Times", "", 10)
	secondPartyText := fmt.Sprintf(`   suatu %s yang didirikan berdasarkan dan tunduk pada hukum Negara Republik Indonesia beralamat di %s. Dalam hal ini diwakili oleh %s selaku %s, oleh karenanya sah bertindak untuk dan atas nama %s sebagai pengelola Apotek %s dengan Nomor Surat Ijin Apotek (SIA) : %s (selanjutnya disebut sebagai "PIHAK KEDUA").`,
		data.SecondPartyType, data.SecondPartyAddress, data.SecondPartyRep, data.SecondPartyTitle, data.SecondPartyName, data.PharmacyName, data.SIANumber)

	pdf.MultiCell(0, 5, secondPartyText, "", "", false)
	pdf.Ln(5)

	// Parties definition
	pdf.SetFont("Times", "", 11)
	partiesText := `PIHAK PERTAMA dan PIHAK KEDUA secara bersama-sama disebut sebagai "PARA PIHAK" dan secara masing-masing disebut sebagai "Pihak". PARA PIHAK sepakat untuk menandatangani Perjanjian ini dengan ketentuan dan syarat-syarat sebagai berikut :`
	pdf.MultiCell(0, 6, partiesText, "", "", false)
	pdf.Ln(8)

	pdf.AddPage()

	// Article 1 - Introduction
	addArticle(pdf, "PASAL 1", "PENDAHULUAN")

	articles1 := []string{
		"PIHAK PERTAMA adalah agregator platform digital apotek yang bernama Apotek Daring dengan ini setuju untuk bekerjasama dengan PIHAK KEDUA.",
		"PIHAK KEDUA adalah apotek mitra yang bergabung dalam program kerja sama Apotek Daring.",
		"PARA PIHAK sepakat untuk melakukan program kerja sama Apotek Daring dan pengelolaan produk secara digital yang berhubungan dengan distribusi farmasi.",
	}
	addNumberedItems(pdf, articles1)

	// Article 2 - Duration
	addArticle(pdf, "PASAL 2", "JANGKA WAKTU")

	articles2 := []string{
		fmt.Sprintf("Perjanjian ini berlaku selama %s tahun terhitung sejak tanggal %s sampai dengan tanggal %s.", data.Duration, data.StartDate, data.EndDate),
		"Jangka waktu dapat diperpanjang dengan salah satu Pihak mengajukan perpanjangan kepada Pihak lainnya paling lambat 1 (satu) bulan sebelum berakhirnya jangka waktu. Perjanjian ini efektif diperpanjang setelah ada kesepakatan secara tertulis dari PARA PIHAK.",
		"Perjanjian ini akan tetap berlaku dan mengikat PARA PIHAK walaupun masing-masing Pihak mengalami perubahan anggaran dasar, perubahan susunan pemegang saham dan/atau Pihak pengendali, perubahan susunan pengurus, perubahan nama, perubahan status badan hukum dan/atau hadirnya ahli waris dari salah satu Pihak, termasuk apabila terjadinya penggabungan dan/atau peleburan.",
	}
	addNumberedItems(pdf, articles2)

	// Add remaining articles (abbreviated for space)
	addRemainingArticles(pdf)

	// Signature section
	addSignatureSection(pdf, data)

	return pdf.OutputFileAndClose("pharmacy_contract.pdf")
}

func addArticle(pdf *gofpdf.Fpdf, articleNum, articleTitle string) {

	pdf.SetFont("Times", "B", 12)

	// Center the text manually
	pageWidth, _ := pdf.GetPageSize()
	left, _, right, _ := pdf.GetMargins()
	textWidth := pageWidth - left - right

	// pdf.Cell(textWidth, 8, articleNum)
	pdf.MultiCell(textWidth, 8, articleNum, "", "C", false)
	// pdf.Ln(8)
	// pdf.Cell(textWidth, 8, articleTitle)
	pdf.MultiCell(textWidth, 8, articleTitle, "", "C", false)
	pdf.Ln(3)
}

func addNumberedItems(pdf *gofpdf.Fpdf, items []string) {
	pdf.SetFont("Times", "", 10)
	for i, item := range items {
		pdf.Cell(10, 6, fmt.Sprintf("%d.", i+1))
		pdf.MultiCell(0, 6, item, "", "", false)
		pdf.Ln(2)
	}
	pdf.Ln(5)
}

func addRemainingArticles(pdf *gofpdf.Fpdf) {
	// Article 3 - Rights and Obligations (abbreviated)
	addArticle(pdf, "PASAL 3", "HAK DAN KEWAJIBAN PARA PIHAK")

	obligations := []string{
		"PIHAK PERTAMA berwenang memastikan bahwa PIHAK KEDUA memiliki lokasi usaha yang sah sesuai dengan ketentuan peraturan perundang-undangan yang berlaku.",
		"PIHAK PERTAMA akan membantu PIHAK KEDUA dalam proses pendaftaran apotek mitra ke dalam platform e-commerce dan memberikan dukungan terhadap pelaksanaan digital marketing.",
		"PIHAK KEDUA wajib menjamin akurasi, kelengkapan, dan keabsahan data yang dimasukkan ke dalam sistem yang digunakan.",
		"PIHAK KEDUA wajib untuk menjaga integritas dalam penggunaan platform e-commerce dan melaporkan setiap kendala kepada PIHAK PERTAMA.",
		"PIHAK KEDUA wajib untuk mengikuti program pelatihan dan evaluasi yang diselenggarakan secara berkala oleh PIHAK PERTAMA.",
	}
	addNumberedItems(pdf, obligations[:5]) // Show first 5 items

	// Article 4 - Confidentiality
	pdf.AddPage()
	addArticle(pdf, "PASAL 4", "KERAHASIAAN")
	confidentiality := []string{
		"PARA PIHAK menjamin bahwa semua informasi rahasia terkait Perjanjian akan dijaga kerahasiaannya dari Pihak ketiga.",
		"Masing-masing Pihak akan menggunakan Informasi Rahasia milik Pihak lainnya hanya untuk tujuan pelaksanaan Perjanjian ini.",
		"Kewajiban untuk menjaga kerahasiaan akan terus berlaku, baik selama jangka waktu Perjanjian maupun setelah pengakhiran Perjanjian.",
	}
	addNumberedItems(pdf, confidentiality)

	// Article 5 - Sanctions and Termination
	addArticle(pdf, "PASAL 5", "SANKSI DAN PENGAKHIRAN")
	sanctions := []string{
		"PARA PIHAK dapat mengakhiri perjanjian ini dengan memberikan pemberitahuan secara tertulis paling lambat 30 (tiga puluh) hari sebelum tanggal efektif Perjanjian diakhiri.",
		"Perjanjian ini dapat berakhir apabila salah satu Pihak telah dinyatakan lalai dan wanprestasi atau terjadi Keadaan Kahar (Force Majeure).",
		"Apabila PIHAK KEDUA terbukti melakukan penyalahgunaan sistem akan dikenakan sanksi berupa sanksi administratif, pemblokiran akses, dan/atau pemutusan kerja sama.",
	}
	addNumberedItems(pdf, sanctions)

	// Article 6 - Force Majeure
	addArticle(pdf, "PASAL 6", "KEADAAN KAHAR (FORCE MAJEURE)")
	forceMajeure := []string{
		"Keadaan Kahar adalah suatu keadaan yang terjadi karena kehendak Tuhan atau di luar kekuasaan manusia yang menyebabkan terhentinya pelaksanaan Perjanjian ini.",
		"Pihak yang mengalami Keadaan Kahar berkewajiban memberitahukan secara tertulis kepada Pihak lainnya dalam waktu 14 (empat belas) hari kalender.",
		"Masing-masing Pihak berhak untuk merundingkan kembali hak dan kewajiban sebagai dampak dari Keadaan Kahar.",
	}
	addNumberedItems(pdf, forceMajeure)

	// Article 7 - Dispute Resolution
	addArticle(pdf, "PASAL 7", "PENYELESAIAN SENGKETA")
	dispute := []string{
		"Setiap perselisihan akan diselesaikan terlebih dahulu melalui musyawarah mufakat dalam jangka waktu 30 (tiga puluh) hari kalender.",
		"Apabila musyawarah tidak berhasil, PARA PIHAK sepakat untuk menyelesaikan berdasarkan hukum Indonesia dan memilih domisili hukum di Pengadilan Negeri Jakarta Selatan.",
	}
	addNumberedItems(pdf, dispute)

	pdf.AddPage()
	// Article 8 - Closing
	addArticle(pdf, "PASAL 8", "PENUTUP")
	closing := []string{
		"Perjanjian ini dibuat, ditafsirkan, dilaksanakan dan tunduk berdasarkan hukum dan peraturan perundang-undangan yang berlaku di Negara Republik Indonesia.",
		"Segala hal yang belum diatur akan dituangkan dalam perjanjian tambahan (addendum) yang merupakan satu kesatuan dengan Perjanjian ini.",
		"Dengan menandatangani Perjanjian ini, PARA PIHAK dianggap telah membaca, memahami, dan menyetujui seluruh Perjanjian ini tanpa paksaan.",
	}
	addNumberedItems(pdf, closing)
}

func addSignatureSection(pdf *gofpdf.Fpdf, data ContractData) {
	pdf.SetFont("Times", "", 11)

	// closingText := `Perjanjian ini dibuat serta ditandatangani oleh PARA PIHAK, pada tempat dan tanggal tersebut pada awal Perjanjian ini, dibuat rangkap 2 (dua), bermeterai cukup serta memiliki kekuatan hukum yang sama.`
	closingText := `Perjanjian ini dibuat serta ditandatangani oleh PARA PIHAK, pada tempat dan tanggal tersebut pada awal Perjanjian ini, dibuat rangkap 2 (dua) serta memiliki kekuatan hukum yang sama.`
	pdf.MultiCell(0, 6, closingText, "", "", false)
	pdf.Ln(20)

	// Signature table using simple approach
	pdf.SetFont("Times", "B", 12)

	// Create a simple table layout
	pdf.Cell(95, 10, "PIHAK PERTAMA")
	pdf.Cell(95, 10, "PIHAK KEDUA")
	pdf.Ln(10)

	// Signature spaces (empty cells for signatures)
	pdf.Cell(95, 40, "")
	pdf.Cell(95, 40, "")
	pdf.Ln(40)

	// Names
	pdf.SetFont("Times", "B", 11)
	pdf.Cell(95, 10, "Yasinta Yulian Hendrata")
	pdf.Cell(95, 10, data.SecondPartyRep)
	pdf.Ln(10)

	pdf.SetFont("Times", "", 10)
	pdf.Cell(95, 8, "Presiden Direktur")
	pdf.Cell(95, 8, data.SecondPartyTitle)
	pdf.Ln(8)
}
