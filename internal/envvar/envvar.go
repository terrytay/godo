package envvar

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/terrytay/godo/internal"
)

type Provider interface {
	Get(key string) (string, error)
}

type Configuration struct {
	provider Provider
}

func Load(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return internal.NewErrorf(internal.ErrorCodeUnknown, "loading env var file")
	}

	return nil
}

func New(provider Provider) *Configuration {
	return &Configuration{
		provider: provider,
	}
}

func (c *Configuration) Get(key string) (string, error) {
	res := os.Getenv(key)
	valSecret := os.Getenv(fmt.Sprintf("%s_SECURE", key))

	if valSecret != "" {
		valSecretRes, err := c.provider.Get(valSecret)
		if err != nil {
			return "", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "provider.Get")
		}

		res = valSecretRes
	}

	return res, nil
}
