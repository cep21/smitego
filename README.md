# smitego [![Circle CI](https://circleci.com/gh/cep21/smitego.svg?style=svg)](https://circleci.com/gh/cep21/smitego)

Smitego is a go library for interacting with the Smite Hi-Rez API.
All documented public functions are implemented.

## Example

```go
package main

import (
  "fmt"
  "github.com/cep21/smitego"
  "golang.org/x/net/context"
)

func main() {
  // First make a client to describe how you want to connect.  Each client
  // returns a session and primary function calls are done on the
  // session.  Concurrent sessions are limited by HiRez
  client := smitego.Client{
    DevID:   123,
    AuthKey: "AuthKey123",
  }

  // A context is how you can time out function calls
  ctx := context.Background()

  // Some functions don't require a session first and can be called on the
  // client directly
  _ = client.Ping(ctx)

  // Most functions require a session

  session, _ := client.CreateSession(ctx)
  gods, _ := session.GetGods(ctx, English)
  fmt.Printf("Got %d gods\n", len(gods))
}
```

## Integration tests

You can test all the library functions by creating a file named info.json and
put it at the root of your project.  That file should have your
devId and authKey.  The file should be inside .gitignore and not checked into
git.  Then run with the tag integration.

File info.json
```
{
  "devId": 123,
  "authKey": "abcd"
}
```

```bash
go test -v --tags=integration .
```
