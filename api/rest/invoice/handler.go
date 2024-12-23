package invoice

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	// "github.com/bhushan217/go-rnd/middleware"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var invoice Invoice
	err = json.Unmarshal(body, &invoice)
	if err != nil {
		panic(err)
	}
	invoice.ID = uint(len(invoices) + 1)
	invoice.Timestamp = time.Now()
	invoices = append(invoices, invoice)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(int(invoice.ID))))
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value(middleware.AUTH_USER_KEY).(string)
	// log.Println(userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	invoice, exists := loadInvoices()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	invoiceOld, exists := loadInvoices()[idStr]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var invoice Invoice
	err = json.Unmarshal(body, &invoice)
	if err != nil {
		panic(err)
	}
	index := IndexOf(invoices, func(x Invoice) bool { return int(x.ID) == id })
	invoiceOld.Name = invoice.Name
	invoiceOld.Amount = invoice.Amount
	invoiceOld.Timestamp = time.Now()
	invoices[index] = invoiceOld

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invoice updated!"))
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	index := IndexOf(invoices, func(x Invoice) bool { return int(x.ID) == id })
	if index < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	invoices = append(invoices[:index], invoices[index+1:]...)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invoice deleted!"))
}

func (h *Handler) PatchByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invoice patched!"))
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
}

type Invoice struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    uint      `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

var invoices []Invoice = []Invoice{
	{ID: 1, Name: "Invoice #1", Amount: 100, Timestamp: time.Now().Add(-24 * time.Hour)},
	{ID: 2, Name: "Invoice #2", Amount: 150, Timestamp: time.Now().Add(-48 * time.Hour)},
	{ID: 3, Name: "Invoice #3", Amount: 200, Timestamp: time.Now().Add(-72 * time.Hour)},
	{ID: 4, Name: "Invoice #4", Amount: 75, Timestamp: time.Now().Add(-96 * time.Hour)},
	{ID: 5, Name: "Invoice #5", Amount: 300, Timestamp: time.Now().Add(-120 * time.Hour)},
}

func loadInvoices() map[string]Invoice {
	res := make(map[string]Invoice, len(invoices))

	for _, x := range invoices {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}

func IndexOf[T comparable](collection []T, fn func(el T) bool) int {
	for i, x := range collection {
		if fn(x) {
			return i
		}
	}
	return -1
}
