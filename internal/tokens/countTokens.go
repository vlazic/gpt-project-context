package tokens

import (
	"fmt"

	"github.com/tiktoken-go/tokenizer"
)

func CountTokens(text string) (int, error) {
	encoding := tokenizer.Cl100kBase

	tke, err := tokenizer.Get(encoding)
	if err != nil {
		return 0, fmt.Errorf("getEncoding: %v", err)
	}

	tokenIds, _, err := tke.Encode(text)
	if err != nil {
		return 0, fmt.Errorf("error encoding to tokens: %v", err)
	}

	// return the number of tokens
	return len(tokenIds), nil
}

// slower 2x:
// "github.com/pkoukk/tiktoken-go"
// func CountTokens(text string) (int, error) {
// 	encoding := "cl100k_base"

// 	tke, err := tiktoken.GetEncoding(encoding)
// 	if err != nil {
// 		return 0, fmt.Errorf("getEncoding: %v", err)
// 	}

// 	// encode
// 	token := tke.Encode(text, nil, nil)

// 	// return the number of tokens
// 	return len(token), nil
// }
