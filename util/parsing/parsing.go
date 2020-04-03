package parsing

import (
	"context"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

//Parser parses text
type Parser struct {
	//TODO - to reduce dependencies, implement OCR individually
	client *vision.ImageAnnotatorClient
	ctx    context.Context
}

//NewParser creates parser
func NewParser() *Parser {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil
	}
	return &Parser{
		client: client,
		ctx:    ctx,
	}
}

//Detect detects text from image
func (parser *Parser) Detect(filename string) []string {
	file, _ := os.Open(filename)
	var descriptions []string
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	annotations, err := parser.client.DetectTexts(parser.ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	if len(annotations) == 0 {

	} else {
		for i := 0; i < len(annotations); i++ {
			descriptions[i] = annotations[i].Description
		}
	}
	return descriptions
}

//Deconstruct deconstructs foods from image
func (parser *Parser) Deconstruct(filename string) []string {
	var foodItems []string

	return foodItems
}
