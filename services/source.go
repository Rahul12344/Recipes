package services

import "github.com/Rahul12344/Recipes/models"

//CrowdSourcer finds crowd sourced items
type CrowdSourcer interface {
	Find(radius, lat, lon float64, item interface{}) []models.Item
	Add(items ...interface{})
	Remove(items ...interface{})
}

//SourcerService sources
type SourcerService struct {
	crowdSourcer CrowdSourcer
}

//NewSourcerService constructs new sourver service
func NewSourcerService(cs CrowdSourcer) *SourcerService {
	return &SourcerService{
		crowdSourcer: cs,
	}
}

//AddItems adds items
func (cs *SourcerService) AddItems(items ...models.Item) {
	cs.crowdSourcer.Add(items)
}

//RemoveItems removes items
func (cs *SourcerService) RemoveItems(items ...models.Item) {
	cs.crowdSourcer.Remove(items)
}

//FindBestLocation find location of closest ingredients at best price
func (cs *SourcerService) FindBestLocation(radius, lat, lon float64, ingredients ...string) ([]string, map[string][]models.Item) {
	var items map[string][]models.Item
	for _, ingredient := range ingredients {
		items[ingredient] = cs.crowdSourcer.Find(radius, lat, lon, ingredient)
	}
	return ingredients, items
}
