
package middlewares
import (

	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)
/*
Check permission Middleware, roles, stuff
*/
func CheckPermission(role string) gin.HandlerFunc {
	fmt.Println("Hello", role)
    if role != "admin" {
		return func (c *gin.Context) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"message": "Unauthorized",
			})
		}
	}

	return func (c *gin.Context) {
       c.Set("TestVar", role)
	   c.Next()
	}
}