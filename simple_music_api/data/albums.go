package data

import "gin-form/simple_music_api/models"

var Albums = []models.Album{
	{
		ID:     "1",
		Title:  "Taste of you",
		Artist: "Rezz",
		Price:  1.99,
	},
	{
		ID:     "2",
		Title:  "Go",
		Artist: "Google",
		Price:  9999,
	},
	{
		ID:     "3",
		Title:  "C#",
		Artist: "Microsoft",
		Price:  -1,
	},
}
