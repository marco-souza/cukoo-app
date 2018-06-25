package main

import (
	"fmt"

	"github.com/zserge/webview"
)

func main() {
	team := "estufacriativa"
	// Open wikipedia in a 800x600 resizable window
	window := webview.New(webview.Settings{
		URL:       fmt.Sprintf("https://cuckoo.team/%s", team),
		Title:     "Cukoo Team - App",
		Width:     800,
		Height:    600,
		Resizable: true,
	})

	window.Dispatch(func() {
		// Setup initial userdata in localStorage before open
		window.Eval(`
			document.cookie = 'cuckooUser=' + encodeURI(JSON.stringify({
				fullname: "Marco Antônio"
				id: "eGbxrhCaaVbPMPcPABjU"
				initials: "MA"
			})) + '; heroku-session-affinity=ACyDaANoA24IAbJPbb7///8HYgAOW/liAAJCAWEBbAAAAAFtAAAABXdlYi4xahKbOriMRHVpSXQo9g6GBXmFcMAg'
			+ '; io=-acI2jCzd0ngU67aAACA'
		`)
	})

	window.Run()
}
