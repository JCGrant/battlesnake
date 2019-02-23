package bot

import "github.com/JCGrant/battlesnake/api"

var dirs = []string{"right", "down", "left", "up"}

func Move(data api.SnakeRequest) api.MoveResponse {
	dirIndex := data.Turn % 4
	move := dirs[dirIndex]
	return api.MoveResponse{
		Move: move,
	}
}

func Start(data api.SnakeRequest) api.StartResponse {
	return api.StartResponse{
		Color: "#75CEDD",
	}
}
