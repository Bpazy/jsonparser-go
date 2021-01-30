package jsonparser

import (
	"fmt"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tokenizer := NewTokenizer(`{"min_position":7,"has_more_items":false,"items_html":"Bike","new_latent_count":0,"data":{"length":26,"text":"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."},"numericalArray":[23,28,23,27,31],"StringArray":["Nitrogen","Nitrogen","Carbon","Nitrogen"],"multipleTypesArray":3,"objArray":[{"class":"upper","age":3},{"class":"upper","age":0},{"class":"middle","age":7},{"class":"upper","age":2},{"class":"middle","age":0}]}`)
	tokenizer.Tokenize()
	for _, token := range tokenizer.Tokens {
		fmt.Printf("%s\n", token.Value)
	}
}
