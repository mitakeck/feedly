package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser : 指定 URL をブラウザで開く処理
func OpenBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		return fmt.Errorf("Unable to open on %s", runtime.GOOS)
	}
	return nil
}
