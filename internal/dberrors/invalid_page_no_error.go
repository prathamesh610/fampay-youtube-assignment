package dberrors

import "fmt"

type InvalidPageNumber struct {
	PageNumber int `json:"page-number"`
	ValidTill  int `json:"valid-pages-till"`
}

func (e *InvalidPageNumber) Error() string {
	return fmt.Sprintf("unable to find page-no: %d with valid pages till: %s", e.PageNumber, e.ValidTill)
}
