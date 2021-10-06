package main

import (
	"fmt"
	"os"

	"github.com/theovassiliou/base64url"
	exitcodes "github.com/theovassiliou/go-exitcodes"
)

const testString = "A text content to be base64url encoded"
const encodedTestString = "QSB0ZXh0IGNvbnRlbnQgdG8gYmUgYmFzZTY0dXJsIGVuY29kZWQ"

func main() {
	// encode a []byte and get the encoded value
	output := base64url.Encode([]byte(testString))
	fmt.Println(output) // returns QSB0ZXh0IGNvbnRlbnQgdG8gYmUgYmFzZTY0dXJsIGVuY29kZWQ

	// decode a string
	decodedOutput, err := base64url.Decode(encodedTestString)

	if err != nil {
		fmt.Println(err)
		os.Exit(exitcodes.DATA_FORMAT_ERROR)
	}

	fmt.Println(string(decodedOutput))

	// check for symmetry
	decodedOutput, _ = base64url.Decode(base64url.Encode([]byte(testString)))
	fmt.Println(string(decodedOutput))

}
