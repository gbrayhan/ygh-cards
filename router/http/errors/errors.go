package errors

import (
  "net/http"

  "github.com/gin-gonic/gin"

  domain "github.com/gbrayhan/academy-go-q12021/domain/errors"
)

type messagesResponse struct {
  Messages []string `json:"messages"`
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
  // Execute request handlers and then handle any errors
  c.Next()

  errs := c.Errors

  if len(errs) > 0 {
    err, ok := errs[0].Err.(*domain.AppError)
    if ok {
      resp := messagesResponse{Messages: []string{err.Error()}}
      switch err.Type {
      case domain.NotFound:
        c.JSON(http.StatusNotFound, resp)
        return
      case domain.ValidationError:
        c.JSON(http.StatusBadRequest, resp)
        return
      case domain.ResourceAlreadyExists:
        c.JSON(http.StatusConflict, resp)
        return
      case domain.NotAuthenticated:
        c.JSON(http.StatusUnauthorized, resp)
        return
      case domain.NotAuthorized:
        c.JSON(http.StatusForbidden, resp)
        return
      case domain.RepositoryError:
        c.JSON(http.StatusInternalServerError, messagesResponse{Messages: []string{"We are working to improve the flow of this request."}})
        return
      default:
        c.JSON(http.StatusInternalServerError, messagesResponse{Messages: []string{"We are working to improve the flow of this request."}})
        return
      }
    }

    // Error is not AppError, return a generic internal server error
    c.JSON(http.StatusInternalServerError, "Internal Server Errror")
    return
  }
}
