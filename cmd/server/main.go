package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type InvoiceResult struct {
	InvoiceNo string `json:"invoiceNo"`
	PCC       string `json:"pcc"`
	BPPS      string `json:"bpps"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
	Date      string `json:"date"`
	Note      string `json:"note"`
	Result    string `json:"result"`
}

type PageData struct {
	Title string
}

func main() {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = tmpl.Execute(w, PageData{Title: "Fatura Mutabakat Ekranı"})
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var invoiceNo string

		if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			var body struct {
				InvoiceNo string `json:"invoiceNo"`
			}
			_ = json.NewDecoder(r.Body).Decode(&body)
			invoiceNo = strings.TrimSpace(body.InvoiceNo)
		} else {
			_ = r.ParseForm()
			invoiceNo = strings.TrimSpace(r.FormValue("invoice_no"))
		}

		result := findInvoice(invoiceNo)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	})

	log.Println("Fatura Mutabakat Ekranı started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func findInvoice(invoiceNo string) InvoiceResult {
	switch strings.ToUpper(invoiceNo) {
	case "":
		return InvoiceResult{
			InvoiceNo: invoiceNo,
			Result:    "error",
			Note:      "Fatura numarası boş bırakılamaz.",
		}
	case "FD60351N9DD0EA":
		return InvoiceResult{
			InvoiceNo: invoiceNo,
			PCC:       "Var",
			BPPS:      "Var",
			Amount:    "325,50 TL",
			Status:    "Uyumlu",
			Date:      "22.03.2026",
			Note:      "Her iki sistemde kayıt bulundu.",
			Result:    "success",
		}
	case "FD60351NA6DA85":
		return InvoiceResult{
			InvoiceNo: invoiceNo,
			PCC:       "Var",
			BPPS:      "Yok",
			Amount:    "210,00 TL",
			Status:    "Uyuşmazlık",
			Date:      "22.03.2026",
			Note:      "Fatura PCC'de mevcut, BPPS tarafında bulunamadı.",
			Result:    "warning",
		}
	case "FD60351O9AB618":
		return InvoiceResult{
			InvoiceNo: invoiceNo,
			PCC:       "Var",
			BPPS:      "Var",
			Amount:    "PCC: 450,00 TL / BPPS: 430,00 TL",
			Status:    "Tutar Farkı",
			Date:      "21.03.2026",
			Note:      "Her iki sistemde kayıt var ancak amount alanı farklı.",
			Result:    "warning",
		}
	default:
		return InvoiceResult{
			InvoiceNo: invoiceNo,
			PCC:       "Bilinmiyor",
			BPPS:      "Bilinmiyor",
			Amount:    "-",
			Status:    "Bulunamadı",
			Date:      "-",
			Note:      "Demo veride bu fatura numarasına ait kayıt yok. Gerçek projede burada PCC/BPPS sorgusu çalışacak.",
			Result:    "info",
		}
	}
}
