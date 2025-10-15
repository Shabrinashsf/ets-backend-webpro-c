package middleware

import (
	"net/http"

	"github.com/Shabrinashsf/ets-backend-webpro-c/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/response"
	"github.com/gin-gonic/gin"
)

func OnlyAllow(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.MustGet("role").(string)

		for _, role := range roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}

		response.BuildResponseFailed(dto.MESSAGE_FAILED_TOKEN_NOT_VALID, dto.ErrRoleNotAllowed.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusForbidden, nil)
	}
}
