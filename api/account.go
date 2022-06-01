package api

import (
	"net/http"

	db "github.com/NaokiYazawa/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
} // アカウント作成時は残高が 0 であるため、Balance は除いた

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // 正常なリクエストかどうかを検証
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // errorResponse でオブジェクトへ変換
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// HandlerFunc が Context 入力を持つ関数として宣言されている
