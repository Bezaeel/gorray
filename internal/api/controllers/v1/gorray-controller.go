package v1

import (
	"encoding/csv"
	"errors"
	"net/http"

	"github.com/bezaeel/gorray/internal/pkg/repository"
	"github.com/bezaeel/gorray/pkg/communication"
	"github.com/gin-gonic/gin"
)
func Multiply(c *gin.Context) {
	s := repository.GetGorrayRepository();
	upload, err := c.FormFile("file")
	if err != nil {
		communication.NewError(c, http.StatusBadRequest, errors.New("expected file of name 'file'"))
	}
	file, err :=  upload.Open()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusBadRequest, 
			errors.New("error reading file"),
		)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusInternalServerError, 
			errors.New("error: unable to process request at this time"),
		)
	}
	result := s.Multiply(records)
	c.JSON(http.StatusOK, result);
}

func Sum(c *gin.Context) {
	s := repository.GetGorrayRepository();
	upload, err := c.FormFile("file")
	if err != nil {
		communication.NewError(c, http.StatusBadRequest, errors.New("expected file of name 'file'"))
	}
	file, err :=  upload.Open()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusBadRequest, 
			errors.New("error reading file"),
		)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusInternalServerError, 
			errors.New("error: unable to process request at this time"),
		)
	}
	result := s.Sum(records)
	c.JSON(http.StatusOK, result);
}

func Flatten(c *gin.Context) {
	s := repository.GetGorrayRepository();
	upload, err := c.FormFile("file")
	if err != nil {
		communication.NewError(c, http.StatusBadRequest, errors.New("expected file of name 'file'"))
	}
	file, err :=  upload.Open()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusBadRequest, 
			errors.New("error reading file"),
		)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusInternalServerError, 
			errors.New("error: unable to process request at this time"),
		)
	}
	result := s.Flatten(records)
	c.JSON(http.StatusOK, result);
}

func Transpose(c *gin.Context) {
	s := repository.GetGorrayRepository();
	upload, err := c.FormFile("file")
	if err != nil {
		communication.NewError(c, http.StatusBadRequest, errors.New("expected file of name 'file'"))
	}
	file, err :=  upload.Open()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusBadRequest, 
			errors.New("error reading file"),
		)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusInternalServerError, 
			errors.New("error: unable to process request at this time"),
		)
	}
	result := s.Transpose(records)
	c.String(http.StatusOK, result);
}

func Echo(c *gin.Context) {
	s := repository.GetGorrayRepository();
	upload, err := c.FormFile("file")
	if err != nil {
		communication.NewError(c, http.StatusBadRequest, errors.New("expected file of name 'file'"))
	}
	file, err :=  upload.Open()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusBadRequest, 
			errors.New("error reading file"),
		)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if  err != nil {
		communication.NewError(
			c, 
			http.StatusInternalServerError, 
			errors.New("error: invalid file format"),
		)
	}
	result := s.Echo(records)
	c.String(http.StatusOK, result);
}