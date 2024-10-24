package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Envelope map[string]any

func (e Envelope) String() string {
	js, err := json.Marshal(e)
	if err != nil {
		log.Println("error:", err)
		return ""
	}
	return string(js)
}

func WriteJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
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

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		log.Println("error: body must only have a single JSON value")
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("error: body must only have a single JSON value")
	}
	return nil
}
