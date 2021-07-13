package main

import (
	"net/http"
	"time"
)

// An Event contains the widget ID, a body of data,
// and an optional target (only "dashboard" for now).
type Event struct {
	ID     string
	Body   map[string]interface{}
	Target string
}

func NewEvent(id string, data map[string]interface{}, target string) *Event {
	data["id"] = id
	data["updatedAt"] = int32(time.Now().Unix())
	return &Event{
		ID:     id,
		Body:   data,
		Target: target,
	}
}

// Dashing struct definition.
type Dashing struct {
	started bool
	Broker  *Broker
	Worker  *Worker
	Server  *Server
	Router  http.Handler
}

// ServeHTTP implements the HTTP Handler.
func (d *Dashing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !d.started {
		panic("dashing.Start() has not been called")
	}
	d.Router.ServeHTTP(w, r)
}

// Start actives the broker and workers.
func (d *Dashing) Start() *Dashing {
	if !d.started {
		if d.Router == nil {
			d.Router = d.Server.NewRouter()
		}
		d.Broker.Start()
		d.Worker.Start()
		d.started = true
	}
	return d
}

// NewDashing sets up the event broker, workers and webservice.
func NewDashing(root string, port string, token string) *Dashing {
	broker := NewBroker()
	worker := NewWorker(broker)
	server := NewServer(broker)

	server.webroot = root
	worker.webroot = root
	worker.url = "http://127.0.0.1:" + port
	worker.token = token

	server.dev = debugmode

	return &Dashing{
		started: false,
		Broker:  broker,
		Worker:  worker,
		Server:  server,
	}
}
