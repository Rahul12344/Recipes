package queries

//CreateRecipeQueryForArgs Creates query for a recipe based on args provided.
func CreateRecipeQueryForArgs(args map[string]interface{}) string {
	ingredients := getIngredients(args)
	name := getName(args)
	country := getCountry(args)
	nutrition := getNutrition(args)
	if len(ingredients) != 0 {

	}
	if name != "" {

	}
	if country != "" {

	}
	if len(nutrition) != 0 {

	}
	return ""
}

func getIngredients(args map[string]interface{}) []string {
	i := args[INGREDIENTS]
	ingredients, ok := i.([]string)
	if !ok {
		return nil
	}
	return ingredients
}

func getName(args map[string]interface{}) string {
	i := args[NAME]
	name, ok := i.(string)
	if !ok {
		return ""
	}
	return name
}

func getCountry(args map[string]interface{}) string {
	i := args[COUNTRY]
	country, ok := i.(string)
	if !ok {
		return ""
	}
	return country
}

func getNutrition(args map[string]interface{}) []string {
	i := args[NUTRITION]
	nutrition, ok := i.([]string)
	if !ok {
		return nil
	}
	return nutrition
}
