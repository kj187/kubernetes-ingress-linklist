package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

// CheckAndWriteHTTPError ...
func CheckAndWriteHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}

// GetFuncMap ...
func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"example": func(title string) string { return "example" },
	}
}
