package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PawapayClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

type InitiatePaymentRequest struct {
	DepositID     string `json:"depositId"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Correspondent string `json:"correspondent"`
	Payer         Payer  `json:"payer"`
	CustomerID    string `json:"customerTimestamp"`
}

type Payer struct {
	Type        string `json:"type"`
	Address     Address `json:"address"`
	DisplayName string `json:"displayName"`
}

type Address struct {
	Value string `json:"value"`
}

type InitiatePaymentResponse struct {
	DepositID string `json:"depositId"`
	Status    string `json:"status"`
	Created   string `json:"created"`
}

type PaymentCallback struct {
	DepositID     string `json:"depositId"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Correspondent string `json:"correspondent"`
	FailureReason *FailureReason `json:"failureReason,omitempty"`
}

type FailureReason struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func NewPawapayClient() *PawapayClient {
	return &PawapayClient{
		BaseURL: os.Getenv("PAWAPAY_BASE_URL"),
		APIKey:  os.Getenv("PAWAPAY_API_KEY"),
		Client:  &http.Client{},
	}
}

func (p *PawapayClient) InitiatePayment(req InitiatePaymentRequest) (*InitiatePaymentResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", p.BaseURL+"/deposits", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.APIKey)

	resp, err := p.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("payment initiation failed with status: %d", resp.StatusCode)
	}

	var result InitiatePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *PawapayClient) HandleCallback(callback PaymentCallback) error {
	// Process callback based on status
	switch callback.Status {
	case "COMPLETED":
		// Update payment status to completed
		return nil
	case "FAILED":
		// Update payment status to failed
		return nil
	default:
		// Handle other statuses (PENDING, etc.)
		return nil
	}
}
