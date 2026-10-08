package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"walletAPI/internal/application/transaction"
	"walletAPI/internal/application/wallet"
)

type WalletHandler struct {
	createWallet  *wallet.CreateWallet
	depositMoney  *wallet.DepositMoneyRequest
	withdrawMoney *wallet.WithdrawMoney
	transferMoney *transaction.TransferMoney
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
	transferMoney *transaction.TransferMoney) *WalletHandler {
	return &WalletHandler{
		createWallet:  createWallet,
		depositMoney:  depositMoney,
		withdrawMoney: withdrawMoney,
		transferMoney: transferMoney,
	}
}

// CreateWallet godoc
// @Summary Cria uma nova carteira
// @Description Cria uma nova carteira associada a um usuário
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body CreateWalletRequest true "Dados da carteira"
// @Success 201 {object} wallet.Wallet
// @Failure 400 {object} map[string]string
// @Router /wallets [post]
func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var req CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	wallet, err := h.createWallet.Execute(req.UserID)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())

		return
	}

	json.NewEncoder(w).Encode(wallet)
}

// DepositMoney godoc
// @Summary Realiza um depósito
// @Description Adiciona saldo à carteira informada
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id path int64 true "ID da carteira"
// @Param request body DepositMoneyRequest true "Dados do depósito"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wallets/{id}/deposit [post]
func (h *WalletHandler) DepositMoney(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	walletID, err := strconv.ParseInt(id, 10, 64)

	var req DepositMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.depositMoney.Execute(walletID, req.Amount)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// WithdrawMoney godoc
// @Summary Realiza um saque
// @Description Retira saldo da carteira informada
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id path int64 true "ID da carteira"
// @Param request body WithdrawMoneyRequest true "Dados do saque"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wallets/{id}/withdraw [post]
func (h *WalletHandler) WithdrawMoney(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	walletID, err := strconv.ParseInt(id, 10, 64)

	var req DepositMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.withdrawMoney.Execute(walletID, req.Amount)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TransferMoney godoc
// @Summary Realiza uma transferência
// @Description Transfere saldo de uma carteira para outra
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body TransferMoneyRequest true "Dados da transferência"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wallets/transfer [post]
func (h *WalletHandler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	var req TransferMoneyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.transferMoney.Execute(
		req.FromWalletID,
		req.ToWalletID,
		req.Amount,
	)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
