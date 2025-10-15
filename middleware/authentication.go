package middleware

import (
	"net/http"
	"strings"

	"github.com/Shabrinashsf/ets-backend-webpro-c/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/response"
	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		role, err := jwtService.GetRoleByToken(authHeader)
		if err != nil {
			response := response.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("user_id", userId)
		ctx.Set("role", role)
		ctx.Next()
	}
}
