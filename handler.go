package hatter

import (
	"net/http"

	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var (
	_ http.Handler = new(Handler)
)

type Handler struct {
	handler func(Context) (Response, error)
	logger  logrus.FieldLogger
}

func NewHandler(handleFunc func(Context) (Response, error), options ...func(*Handler)) *Handler {
	var handler = &Handler{
		handler: handleFunc,
	}
	for _, option := range options {
		option(handler)
	}
	if handler.logger == nil {
		var logger = logrus.StandardLogger()
		logger.Formatter = &logrus.TextFormatter{
			FullTimestamp: true,
		}
		logger.SetLevel(logrus.DebugLevel)
		handler.logger = logger
	}
	return handler
}

type Context struct {
	Request Request
	Logger  logrus.FieldLogger
}

func WithLogger(logger logrus.FieldLogger) func(handler *Handler) {
	return func(handler *Handler) {
		handler.logger = logger
	}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var aliceRequest Request
	if err := json.NewDecoder(request.Body).Decode(&aliceRequest); err != nil {
		handler.logger.WithError(err).Errorf("unable to decode Alice request")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := handler.handler(Context{
		Request: aliceRequest,
		Logger:  handler.logger,
	})
	if err != nil {
		handler.logger.WithError(err).Errorf("unable to process Alice request")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Version = aliceRequest.Version
	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		handler.logger.WithError(err).Errorf("unable to encode response")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("server", "The Mad Hatter")
}
