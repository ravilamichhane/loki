package thor

import (
	"context"
	"encoding/json"
	"fmt"
	"loki/common"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type value struct {
	TraceID    string
	Now        time.Time
	StatusCode int
	UserID     string
}

type ctxKey int

const key ctxKey = 1

func newContext(rw http.ResponseWriter, r *http.Request, l common.Logger) *Context {

	r = r.WithContext(context.WithValue(r.Context(), key, &value{
		TraceID: uuid.New().String(),
		Now:     time.Now(),
	}))

	return &Context{
		Logger:         l,
		ResponseWriter: rw,
		Request:        r,
	}
}
func getContext(ctx context.Context) *value {
	v, ok := ctx.Value(key).(*value)
	if !ok {
		return &value{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}
	return v
}

func GetTraceID(ctx context.Context) string {
	c := getContext(ctx)
	return c.TraceID
}

func GetTime(ctx context.Context) time.Time {
	c := getContext(ctx)
	return c.Now
}

func setStatusCode(ctx context.Context, statusCode int) {
	c := getContext(ctx)
	c.StatusCode = statusCode
}

func setUserID(ctx context.Context, userID string) {
	c := getContext(ctx)
	c.UserID = userID
}

func setTraceID(ctx context.Context, traceID string) {
	c := getContext(ctx)
	c.TraceID = traceID
}

type Context struct {
	Logger         common.Logger
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c Context) GetRequest() *http.Request {
	return c.Request
}

func (c Context) GetLogger() common.Logger {
	return c.Logger
}

func (c Context) GetWriter() http.ResponseWriter {
	return c.ResponseWriter
}

func (c Context) GetParam(name string) string {
	return mux.Vars(c.Request)[name]
}

func (c Context) GetUUIDParam(name string) (uuid.UUID, error) {
	id, err := uuid.Parse(mux.Vars(c.Request)[name])
	if err != nil {
		return uuid.UUID{}, NewTrustedError(fmt.Errorf("invalid UUID"), http.StatusBadRequest)
	}
	return id, nil
}

func (c Context) GetStatusCode() int {
	return getContext(c.Request.Context()).StatusCode
}

func (c Context) GetTraceID() string {
	return getContext(c.Request.Context()).TraceID
}

func (c Context) SetTraceID(traceID string) {
	setTraceID(c.Request.Context(), traceID)
	getContext(c.Request.Context()).TraceID = traceID
}

func (c *Context) SetUserID(userID string) {
	setUserID(c.Request.Context(), userID)
}

func (c Context) GetUserID() string {
	return c.Request.Context().Value("userID").(string)
}

func (c *Context) SetContext(ctx context.Context) {
	c.Request = c.Request.WithContext(ctx)
}

func (c Context) GetPath() string {
	return c.Request.URL.Path
}

func (c Context) GetTime() time.Time {
	return getContext(c.Request.Context()).Now
}

func (c Context) GetRemoteAddr() string {
	return c.Request.RemoteAddr
}

func (c Context) GetMethod() string {
	return c.Request.Method
}

func NewContext(rw http.ResponseWriter, r *http.Request, l common.Logger) *Context {

	r = r.WithContext(context.WithValue(r.Context(), key, &value{
		TraceID: uuid.New().String(),
		Now:     time.Now(),
	}))

	return &Context{
		Logger:         l,
		ResponseWriter: rw,
		Request:        r,
	}
}

func (c Context) GetContext() context.Context {
	return c.Request.Context()
}

func (c Context) StatusCode() int {
	return getContext(c.Request.Context()).StatusCode
}

func (c Context) SetStatusCode(statusCode int) {
	setStatusCode(c.Request.Context(), statusCode)
	// c.ResponseWriter.WriteHeader(statusCode)
}

func (c Context) Text(status int, body string) error {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.ResponseWriter.WriteHeader(status)
	_, err := c.ResponseWriter.Write([]byte(body))
	return err
}

func (c Context) SetHeader(key, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

func (c Context) JSON(status int, body interface{}) error {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(status)
	return json.NewEncoder(c.ResponseWriter).Encode(body)
}

type validator interface {
	Validate() error
}

func (ctx Context) Decode(val any) error {
	r := ctx.Request
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		if err.Error() == "EOF" {
			return NewTrustedError(fmt.Errorf("no Body Provided"), http.StatusBadRequest)
		}

		return NewTrustedError(
			fmt.Errorf("invalid JSON"),
			http.StatusBadRequest)
	}

	if v, ok := val.(validator); ok {

		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
