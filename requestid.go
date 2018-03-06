//
// ginrequestid
//
// Set an UUID4 string as Request ID into response headers ("X-Request-Id") and
// expose that value as "RequestId" in order to use it inside the application for logging
// or propagation to other systems.
//
package ginrequestid

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		request_id := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if request_id == "" {
            uuid4, _ := uuid.NewV4()
            requestID = uuid4.String()
		}

		// Expose it for use in the application
		c.Set("RequestId", request_id)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", request_id)
		c.Next()
	}
}
