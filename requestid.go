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
	"github.com/gofrs/uuid"
)

// PanicOnUUIDError tells the middleware whether to panic if generating a requestID fails
// or to eat up the error and just avoid setting a request ID altogether.
var PanicOnUUIDError bool

const (
	// Header is the value of the header in which we expect to find the request id.
	Header = "x-request-id"

	// ContextKey is the key value of gins context in which the request id can be found.
	ContextKey = "RequestId"
)

// RequestID is a helper function used for returning the
// value of the request id from a gin context.
func RequestID(c *gin.Context) string {
	return c.GetString(ContextKey)
}

// RequestIDHandler middleware enriches the gin context .
func RequestIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(Header)

		// Create request id with UUIDv4
		if requestID == "" {
			uuid, err := uuid.NewV4()

			// If there's an error generating a UUID then the randomness source
			// is depleted or something like that. We have no way of handling
			// this on our own, so we panic if the user told us to do so.
			if err != nil {
				if PanicOnUUIDError {
					panic(err)
				}
				c.Next()
				return
			}

			requestID = uuid.String()
		}

		// Expose it for use in the application
		c.Set(ContextKey, requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set(Header, requestID)
		c.Next()
	}
}
