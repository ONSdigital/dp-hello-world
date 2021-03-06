package handlers

import (
	"encoding/json"
	"github.com/ONSdigital/dp-hello-world-controller/config"
	"github.com/ONSdigital/dp-hello-world-controller/mapper"
	"github.com/ONSdigital/log.go/log"
	"net/http"
)

// ClientError is an interface that can be used to retrieve the status code if a client has errored
type ClientError interface {
	Error() string
	Code() int
}

func setStatusCode(req *http.Request, w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if err, ok := err.(ClientError); ok {
		if err.Code() == http.StatusNotFound {
			status = err.Code()
		}
	}
	log.Event(req.Context(), "setting-response-status", log.Error(err))
	w.WriteHeader(status)
}

// HelloWorld Handler
func HelloWorld(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		helloWorld(w, req, cfg)
	}
}

func helloWorld(w http.ResponseWriter, req *http.Request, cfg config.Config) {
	ctx := req.Context()
	greetingsModel := mapper.HelloModel{Greeting: "Hello", Who: "World"}
	m := mapper.HelloWorld(ctx, greetingsModel, cfg)

	b, err := json.Marshal(m)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Event(ctx, "failed to write bytes for http response", log.Error(err))
		setStatusCode(req, w, err)
		return
	}
	return
}
