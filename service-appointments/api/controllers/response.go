package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shaahinm/calendar/config"
)

type JsonResponseWriter struct {
	writer *http.ResponseWriter
}

type StandardContainer struct {
	Data    any `json:"data"`
	Code    int `json:"code"`
	Version int `json:"version"`
}

func NewJsonResponseWriter(writer *http.ResponseWriter) *JsonResponseWriter {
	return &JsonResponseWriter{
		writer: writer,
	}
}

var version = config.Envs.ResponseVersion

func Ok(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusOK)
	jsonResponseWriter.addPayload(payload, http.StatusOK)
}

func Created(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusCreated)
	jsonResponseWriter.addPayload(payload, http.StatusCreated)
}

func Accepted(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusAccepted)
	jsonResponseWriter.addPayload(payload, http.StatusAccepted)
}

func UnprocessableContent(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusUnprocessableEntity)
	jsonResponseWriter.addPayload(payload, http.StatusUnprocessableEntity)
}

func Unauthorized(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusUnauthorized)
	jsonResponseWriter.addPayload(payload, http.StatusUnauthorized)
}

func Bad(writer *http.ResponseWriter, payload any) {
	jsonResponseWriter := NewJsonResponseWriter(writer)
	jsonResponseWriter.addHeader()
	jsonResponseWriter.addStatus(http.StatusBadRequest)
	jsonResponseWriter.addPayload(payload, http.StatusBadRequest)
}

func (r *JsonResponseWriter) addHeader() {
	writer := *r.writer
	writer.Header().Set("Content-type", "application/json")
}

func (r *JsonResponseWriter) addStatus(status int) {
	writer := *r.writer
	writer.WriteHeader(status)
}

func (r *JsonResponseWriter) addPayload(data any, code int) {
	writer := *r.writer
	standardContainer := StandardContainer{
		Data:    data,
		Code:    code,
		Version: version,
	}
	serializedData, _ := json.Marshal(&standardContainer)
	writer.Write(serializedData)
}
