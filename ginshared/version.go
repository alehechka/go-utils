package ginshared

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VersionHandler provides a simple handler that will return the provided version in the accepted format
func VersionHandler(version string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accept := ctx.GetHeader("Accept")
		verObj := gin.H{"version": version}

		if strings.Contains(accept, "application/json") {
			ctx.JSON(http.StatusOK, verObj)
			return
		}

		if strings.Contains(accept, "application/xml") {
			ctx.XML(http.StatusOK, verObj)
			return
		}

		ctx.Data(http.StatusOK, "text/plain", []byte(version))
	}
}
