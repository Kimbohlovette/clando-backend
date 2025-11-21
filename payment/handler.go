package payment

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	client *PawapayClient
}

func NewHandler() *Handler {
	return &Handler{
		client: NewPawapayClient(),
	}
}

func (h *Handler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var req InitiatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.InitiatePayment(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) HandlePaymentCallback(w http.ResponseWriter, r *http.Request) {
	var callback PaymentCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.client.HandleCallback(callback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
