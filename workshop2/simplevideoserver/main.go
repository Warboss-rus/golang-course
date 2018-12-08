package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/list", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `[{
			"id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
			"name": "Black Retrospetive Woman",
			"duration": 15,
			"thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"
		}]`)
	})
	http.HandleFunc("/api/v1/video/d290f1ee-6c54-4b01-90e6-d701748f0851", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{
	    	"id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    		"name": "Black Retrospetive Woman",
    		"duration": 15,
    		"thumbnail":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
    		"url":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4" 
		}`)
	})
	http.ListenAndServe(":8000", nil)
}
