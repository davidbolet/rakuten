package processor

import (
	"encoding/xml"
	"fmt"
	"log"

	"cursmedia.com/rakuten/model"
)

// Process parsed the given xml represented as string
func Process(input *string) (*model.ProductResult, error) {
	var result model.ProductResult
	result = model.ProductResult{}
	log.Println(*input)
	err := xml.Unmarshal([]byte(*input), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	return &result, nil
}
