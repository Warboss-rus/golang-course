package handlers

import "errors"

var videos = []Video{
	{
		"d290f1ee-6c54-4b01-90e6-d701748f0851",
		"Black Retrospetive Woman",
		15,
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	},
	{
		"sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
		"Go Rally TEASER-HD",
		41,
		"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
		"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4",
	},
	{
		"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
		"Танцор",
		92,
		"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
		"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4",
	},
}

func findVideoById(id string) (Video, error) {
	for _, v := range videos {
		if v.Id == id {
			return v, nil
		}
	}
	return Video{}, errors.New("Invalid video requested. Id=" + id)
}
