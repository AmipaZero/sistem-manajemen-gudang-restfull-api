package helper
import (
	"net/http"
	"sistem-manajemen-gudang/model/web"
	"github.com/gin-gonic/gin"
)

//  success response
func Success(ctx *gin.Context, code int, data interface{}) {
	if data == nil {
		data = gin.H{} // default kosong
	}

	response := web.WebResponse{
		Code:   code,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(code, response)
}

//  bad request
func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD_REQUEST",
		Data:   gin.H{"message": message},
	})
}

//  unauthorized
func Unauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   gin.H{"message": message},
	})
}

//  not found
func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "NOT_FOUND",
		Data:   gin.H{"message": message},
	})
}

//  internal server error
func InternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "ERROR",
		Data:   gin.H{"message": message},
	})
}
