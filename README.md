# feedly
feedly cloud api wrapper for golang

## Usage

### Authentication

https://developer.feedly.com/v3/auth/

```golang:
feedly := Feedly{}

token, err := feedly.Auth()
fmt.Print(token.AccessToken)
```

### Profile

https://developer.feedly.com/v3/profile/

```golang:
profile, err := feedly.Profile()
fmt.Print(profile.Email)
```

### Markers

https://developer.feedly.com/v3/markers/

```golang:
markers, err := feedly.Markers()
fmt.Print(markers)
```

## TODO

- [x] Authentication
- [x] Categories
- [ ] Dropbox
- [x] Entries
- [ ] Evernote
- [ ] Facebook
- [x] Feeds
- [x] Markers
- [ ] Microsoft
- [x] Mixes
- [ ] OPML
- [ ] Preferences
- [x] Profile
- [x] Search
- [x] Streams
- [x] Subscriptions
- [x] Tags
- [ ] Twitter
