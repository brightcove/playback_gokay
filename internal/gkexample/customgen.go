package gkexample

import (
	"fmt"

	"github.com/brightcove/playback_gokay/gkgen"
)

// To run: `gokay gkexample NewCustomGenerator`
func NewCustomGenerator() *gkgen.ValidateGenerator {
	fmt.Println("Generating code with a custom validator that's the same as the default validator")
	return gkgen.NewValidateGenerator()
}
