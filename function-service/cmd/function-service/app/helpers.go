package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (app *Application) WriteJSON(w http.ResponseWriter, data interface{}, status int, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func DecodeJSON(body io.Reader, input interface{}) error {
	if err := json.NewDecoder(body).Decode(input); err != nil {
		return err
	}

	v := validator.New()
	v.RegisterValidation("number", func(fl validator.FieldLevel) bool {
		fmt.Printf("%+v", fl.Field())

		return true
	})
	if err := v.Struct(input); err != nil {
		return err
	}

	return nil
}
