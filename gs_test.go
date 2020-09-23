package pubsub

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/igm/sockjs-go/sockjs"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

type wst struct {
	conn sockjs.Session
	client *Client
	state TransportState
}

func (t wst) Write(d []byte) (int, error) {
	fmt.Println("writing!!!")
	return len(d), t.conn.Send(string(d))
}

func (t wst) State() TransportState {
	return t.state
}

func (t wst) Close() error {
	return t.conn.Close(0, "")
}

func TestPubsub(t *testing.T) {
	pubsub := New(Config{
		Shards:           8,
		PubQueueConfig: PubQueueConfig{
			BufferSize: 400,
			Writers: 16,
			ReadersPerWriter: 4,
		},
		ShardConfig: ShardConfig{
			ClientBufferSize: 250,
		},
		TopicBuckets: 16,
	})

	go pubsub.Start(context.Background())

	h := sockjs.NewHandler("/pubsub", sockjs.DefaultOptions, func(s sockjs.Session) {
		t := wst{
			conn:   s,
			state:  OpenTransport,
		}

		client, err := pubsub.Connect(ConnectOptions{
			Transport:    t,
			ID: NewClientID(uuid.New().String(), uuid.New().String()),
		})

		if err != nil {
			fmt.Println(err)
			t.Close()
			return
		}

		pubsub.Subscribe(SubscribeOptions{
			Topics:       []string{"chats/global"},
			Clients:      []string{string(client.ID())},
		})

		for {
			data, err := s.Recv()
			if err != nil {
				pubsub.InactivateClient(client)
				return
			}
			pubsub.Publish(PublishOptions{
				Topics: []string{"chats/global"},
				Data:   []byte(data),
			})
		}
	})

	http.HandleFunc("/pubsub/", func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		h.ServeHTTP(w, r)
	})
	http.ListenAndServe("localhost:3000", nil)
}
