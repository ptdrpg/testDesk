package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	// "github.com/webview/webview_go"
	"github.com/LightningDev1/webview"
)

// func main() {
// 	// Déterminer le chemin du fichier index.html
// 	dir:= filepath.Join("build/client")
// 	go func () {
// 		fs := http.FileServer(http.Dir(dir))
// 		http.Handle("/", fs)
// 		err := http.ListenAndServe(":3000", nil)
// 		if err != nil {
// 			fmt.Println("erreur", err)
// 		}
// 	}()
// 	// indexPath := "file://" + filepath.Join(dir, "index.html")

// 	// Lancer WebView
// 	w := webview.New(true)
// 	defer w.Destroy()
// 	w.SetTitle("Application Desktop")
// 	w.SetSize(1024, 768, webview.HintNone)
// 	w.Navigate("http://localhost:3000")
// 	w.Run()
// }

func main() {
	dir:= filepath.Join("build/client")
	go func () {
		fs := http.FileServer(http.Dir(dir))
		http.Handle("/", fs)
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			fmt.Println("erreur", err)
		}
	}()

	w := webview.New(1024, 768, "test desktop", false)
	defer w.Destroy()
	w.Navigate("http://localhost:3000")
	w.Run()
}
