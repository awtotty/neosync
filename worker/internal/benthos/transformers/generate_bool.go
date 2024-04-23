package transformers

import (
	"math/rand"

	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_bool", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
		return func() (any, error) {
			return generateRandomBool(), nil
		}, nil
	})
	if err != nil {
		panic(err)
	}
}

// Generates a random bool value and returns it as a bool type.
func generateRandomBool() bool {
	//nolint:gosec
	randInt := rand.Intn(2)
	return randInt == 1
}
