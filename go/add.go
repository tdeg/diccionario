package diccionario

import (
	"log"
	"net/http"

	"github.com/for-the-kidz/diccionario/wordlist"
	"github.com/gin-gonic/gin"
)

// AddRequest represents the request body for the /add endpoint.
type AddRequest struct {
	// Word is the word to add to the word list.
	Word string `json:"word" binding:"required"`
}

// Add a new word to the word list.
func (s *Server) Add(c *gin.Context) {
	var req AddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// TODO: we need to make this error cleaner
		// right now it looks something like this:
		// "error": "Key: 'AddRequest.Word' Error:Field validation for 'Word' failed on the 'required' tag"
		// add debug logging
		c.JSON(http.StatusBadRequest, ApiError{Err: err, Desc: "unable to parse request"})
		return
	}

	log.Println("request:", req)
	log.Println("retrived word:", req.Word)

	wl, err := s.w.GetWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{Err: err, Desc: "unable to retrieve word list"})
		return
	}

	exists, err := wordlist.WordExists(req.Word, wl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{Err: err, Desc: "unable to check if word exists"})
		return
	}

	// TODO: we need to fix the MarshalJSON for the ApiError to handle nil errors or use a different response
	if exists {
		c.JSON(http.StatusConflict, ApiError{Err: nil, Desc: "word already exists"})
		return
	}

	c.Status(http.StatusCreated)

	// implement your logic here
}
