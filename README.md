# feedly-go
feedly cloud api wrapper

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
- [ ] Entries
- [ ] Evernote
- [ ] Facebook
- [ ] Feeds
- [x] Markers
- [ ] Microsoft
- [ ] Mixes
- [ ] OPML
- [ ] Preferences
- [x] Profile
- [ ] Search
- [ ] Streams
- [x] Subscriptions
- [ ] Tags
- [ ] Twitter
