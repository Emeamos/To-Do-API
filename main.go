package main

type todo struct{
	ID					string	`json:"id"`
	Item				string	`json:"title"`
	Completed			bool	`json:"complete"`
}

var todos = []todo{
	{ID:"1", Item:"read movie", Completed: false},
	{ID:"2", Item:"play game", Completed: false},
	{ID:"3", Item:"record movie", Completed: false},

}