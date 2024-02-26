package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})

	n.Handle("generate", func(msg maelstrom.Message) error {
		body := map[string]any{
			"type": "generate_ok",
			"id":   uuid.New().String(),
		}

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
