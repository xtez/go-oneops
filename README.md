# go-oneops #

go-oneops is a Go client library for accessing the OneOps API.

Currently, **go-oneops requires Go version 1.15 or greater**.

## Usage ##

```go
import "github.com/xtez/go-oneops/oneops"
```

Construct a new OneOps client, then use the various services on the client to
access different parts of the OneOps API. For example:

```go
authTransport := &oneops.BasicAuthTransport{Username: userName,
	Password: userPassword}

client, err := oneops.NewClient(oneopsURL, authTransport.Client())

// list details of authenticated user
user, _, err := client.Users.Get(ctx)
```

## License ##

This library is distributed under the BSD-style license found in the [Apache License 2.0](./LICENSE)
file.
