package pubsub

import (
	"pubsub/publication"
	"time"
)

type SubscribeOptions struct {
	Topics []string
	Clients []string
	Users []string
	EventOptions
}

func (opts SubscribeOptions) defaultEventOptions() {
	if opts.Time == 0 {
		opts.Time = time.Now().UnixNano()
	}
	if opts.Event == "" {
		opts.Event = SubscribeEvent
	}
}

type UnsubscribeOptions struct {
	Topics []string
	Clients []string
	Users []string
	All bool
	EventOptions
}

func (opts UnsubscribeOptions) defaultEventOptions() {
	if opts.Time == 0 {
		opts.Time = time.Now().UnixNano()
	}
	if opts.Event == "" {
		opts.Event = UnsubscribeEvent
	}
}

type PublishOptions struct {
	Topics []string
	Clients []string
	Users []string
	Data []byte
	EventOptions
}

type PublicationOptions struct {
	Topics 		[]string
	Clients 	[]string
	Users 		[]string
	Publication publication.Publication
	EventOptions
}

type ConnectOptions struct {
	Transport Transport
	ID		  ClientID
	EventOptions
}

type DisconnectOptions struct {
	Clients []string
	Users	[]string
	All		bool
	EventOptions
}


func (opts PublishOptions) defaultEventOptions() {
	if opts.Time == 0 {
		opts.Time = time.Now().UnixNano()
	}
	if opts.Event == "" {
		opts.Event = PublishEvent
	}
}

type EventOptions struct {
	Time int64
	Event string
}