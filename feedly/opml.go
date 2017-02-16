package feedly

// OPML : https://developer.feedly.com/v3/opml/
func (f *Feedly) OPML(fileName string) (int64, error) {
	return f.Download(opmlURL, fileName)
}
