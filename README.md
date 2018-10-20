# Arbor

Arbor is an experimental chat protocol that models a conversation
as a tree of messages instead of an ordered list. This means that
the conversation can organically diverge into several conversations
without the messages appearing interleaved.

Arbor is unbelievably primitive right now. With time, it may develop
into something usable, but be warned that it is not currently a delightful
user experience.

## This Repo

This is `pergola`, an Arbor client that focuses on visualizing the message
tree.

Install with `go get github.com/arborchat/pergola`.

Run it with `pergola <IPv4 Address>:<Port>`.

It supports the following key bindings:
- up/down to traverse the message history linearly
- left/right to visit sibling messages of the current message
- enter to compose a reply to the current message (enter again to send)
- ctrl+c to quit
