# helix

A Twitch Helix API client written in Go. If you are looking for a client for Twitch's Kraken API, see [kraken](https://github.com/nicklaw5/kraken).

[![Build Status](https://travis-ci.org/nicklaw5/helix.svg?branch=master)](https://travis-ci.org/nicklaw5/helix)
[![Coverage Status](https://coveralls.io/repos/github/nicklaw5/helix/badge.svg)](https://coveralls.io/github/nicklaw5/helix)

## Package Status

This project is a work in progress. Below is a list of currently supported endpoints and features. Until a release is cut, consider this API to be unstable.

## Supported Endpoints

- [ ] Authentication
- [ ] Get Bits Leaderboard
- [x] Get Clip
- [ ] Create Clip
- [ ] Create Entitlement Grants Upload URL
- [x] Get Games
- [x] Get Top Games
- [ ] Get Game Analytics
- [x] Get Streams
- [x] Get Streams Metadata
- [x] Get Users
- [x] Get Users Follows
- [ ] Update User
- [x] Get Videos

## Usage

This is a quick example of how to get users. Note that you don't need to provide both a list of ids and logins, one or the other will suffice.

```go
client, err := helix.NewClient(&helix.Options{
    ClientID: "your-client-id",
})
if err != nil {
    // handle error
}

resp, err := client.GetUsers(&helix.UsersParams{
    IDs:    []string{"26301881", "18074328"},
    Logins: []string{"summit1g", "lirik"},
})
if err != nil {
    // handle error
}

fmt.Printf("Status code: %d\n", resp.StatusCode)
fmt.Printf("Rate limit: %d\n", resp.RateLimit.Limit)
fmt.Printf("Rate limit remaining: %d\n", resp.RateLimit.Remaining)
fmt.Printf("Rate limit reset: %d\n\n", resp.RateLimit.Reset)

for _, user := range resp.Data.Users {
    fmt.Printf("ID: %s Name: %s\n", user.ID, user.DisplayName)
}
```

Output:

```txt
Status code: 200
Rate limit: 30
Rate limit remaining: 29
Rate limit reset: 1517695315

ID: 26301881 Name: sodapoppin Display Name: sodapoppin
ID: 18074328 Name: destiny Display Name: Destiny
ID: 26490481 Name: summit1g Display Name: summit1g
ID: 23161357 Name: lirik Display Name: LIRIK
```

## Documentation & Examples

All documentation and usage examples for this package can be found [here](docs). If you are looking for the Twitch API docs, see the [Twitch Developer website](https://dev.twitch.tv/docs/api).

## Contributions

PRs are very much welcome. Where possible, please write tests for any code that is introduced by your PRs.

## License

This package is distributed under the terms of the [MIT](License) License.
