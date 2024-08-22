package main

import (
	"io"
	"net/http"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://images.rawpixel.com/image_png_800/cHJpdmF0ZS9sci9pbWFnZXMvd2Vic2l0ZS8yMDIyLTA1L2pvYjcyNC0xODctcC5wbmc.png")
	if err != nil {
		http.Error(w, "Unable to fetch image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", imageHandler)
	http.ListenAndServe(":8080", nil)
}
