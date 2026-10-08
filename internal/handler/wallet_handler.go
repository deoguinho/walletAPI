package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"walletAPI/internal/application/wallet"
)

type WalletHandler struct {
	createWallet  *wallet.CreateWallet
	depositMoney  *wallet.DepositMoneyRequest
	withdrawMoney *wallet.WithdrawMoney
	transferMoney *wallet.TransferMoney
}

type CreateWalletRequest struct {
	UserID int64 `json:"user_id"`
}

type DepositMoneyRequest struct {
	Amount int64 `json:"amount"`
}

type WithdrawMoneyRequest struct {
	Amount int64 `json:"amount"`
}

type TransferMoneyRequest struct {
	FromWalletID int64 `json:"from_wallet_id"`
	ToWalletID   int64 `json:"to_wallet_id"`
	Amount       int64 `json:"amount"`
}

func NewWalletHandler(
	createWallet *wallet.CreateWallet,
	depositMoney *wallet.DepositMoneyRequest,
	withdrawMoney *wallet.WithdrawMoney,
	transferMoney *wallet.TransferMoney) *WalletHandler {
	return &WalletHandler{
		createWallet:  createWallet,
		depositMoney:  depositMoney,
		withdrawMoney: withdrawMoney,
		transferMoney: transferMoney,
	}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var req CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wallet, err := h.createWallet.Execute(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}

func (h *WalletHandler) DepositMoney(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	walletID, err := strconv.ParseInt(id, 10, 64)

	var req DepositMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.depositMoney.Execute(walletID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WalletHandler) WithdrawMoney(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	walletID, err := strconv.ParseInt(id, 10, 64)

	var req DepositMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.withdrawMoney.Execute(walletID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WalletHandler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var req TransferMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.transferMoney.Execute(
		req.FromWalletID,
		req.ToWalletID,
		req.Amount,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
