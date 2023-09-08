package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/simple-bank/db/sqlc"
	"github.com/simple-bank/util"
	"net/http"
	"time"
)

type createUserReq struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserRes struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PasswordCreateAt time.Time `json:"password_create_at"`
	CreatedAt        time.Time `json:"created_at"`
}

func newUserResponse(user db.User) UserRes {
	return UserRes{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordCreateAt: user.PasswordCreateAt,
		CreatedAt:        user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
