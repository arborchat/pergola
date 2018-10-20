package clientio

import (
	"io"
	"log"

	messages "github.com/arborchat/arbor-go"
)

// HandleConn reads from the provided connection and writes new messages to the msgs
// channel as they come in.
func HandleNewMessages(conn io.ReadWriteCloser, msgs chan<- *messages.ChatMessage, welcomes chan<- *messages.ProtocolMessage) {
	readMessages := messages.MakeMessageReader(conn)
	defer close(msgs)
	for fromServer := range readMessages {
		switch fromServer.Type {
		case messages.WelcomeType:
			welcomes <- fromServer
			close(welcomes)
			welcomes = nil
		case messages.NewMessageType:
			// add the new message
			msgs <- fromServer.ChatMessage
		default:
			log.Println("Unknown message type: ", fromServer.String)
			continue
		}
	}
}

// HandleRequests reads from the requestedIds and outbound channels and sends messages
// to the server. Any message id received on the requestedIds channel will be queried
// and any message received on the outbound channel will be sent as a new message
func HandleRequests(conn io.ReadWriteCloser, requestedIds <-chan string, outbund <-chan *messages.ChatMessage) {
	toServer := messages.MakeMessageWriter(conn)
	for {
		select {
		case queryId := <-requestedIds:
			a := &messages.ProtocolMessage{
				Type: messages.QueryType,
				ChatMessage: &messages.ChatMessage{
					UUID: queryId,
				},
			}
			toServer <- a
		case newMesg := <-outbund:
			a := &messages.ProtocolMessage{
				Type:        messages.NewMessageType,
				ChatMessage: newMesg,
			}
			toServer <- a
		}
	}
}
