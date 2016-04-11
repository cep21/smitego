package smitego

import (
	"fmt"
	"golang.org/x/net/context"
)

// Example of creating a session and making a function call.
func ExampleSession() {
	// First make a client to describe how you want to connect.  Each client returns a session
	// and primary function calls are done on the session.  Concurrent sessions are limited
	// by HiRez
	client := Client{
		DevID:   123,
		AuthKey: "AuthKey123",
	}

	// A context is how you can time out function calls
	ctx := context.Background()

	// Some functions don't require a session first and can be called on the client directly
	_ = client.Ping(ctx)

	// Most functions require a session

	session, _ := client.CreateSession(ctx)
	gods, _ := session.GetGods(ctx, English)
	fmt.Printf("Got %d gods\n", len(gods))
}
