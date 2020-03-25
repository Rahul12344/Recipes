package main

import recipes "github.com/Rahul12344/Recipes"

func main() {
	config := recipes.Conf()
	recipes.ExecuteServer(config)
}
