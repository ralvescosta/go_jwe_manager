package environments

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Configure(t *testing.T) {
	t.Run("should configure development env when GO_ENV is undefined", func(t *testing.T) {
		os.Setenv("GO_ENV", "")
		var envFile string
		dotEnvConfig = func(arg string) error {
			envFile = arg
			return nil
		}

		env := NewEnvironment()
		err := env.Configure()

		if err != nil {
			t.Error("dotenv need to call only dotEnvConfig function")
		}
		if envFile != ".env.development" {
			t.Error("dotenv need to call only dotEnvConfig function")
		}
	})

	t.Run("should configure env using the GO_ENV", func(t *testing.T) {
		os.Setenv("GO_ENV", "production")
		var envFile string
		dotEnvConfig = func(arg string) error {
			envFile = arg
			return nil
		}
		env := NewEnvironment()
		err := env.Configure()

		if err != nil {
			t.Error("dotenv need to call only dotEnvConfig function")
		}
		if envFile != ".env.production" {
			t.Error("dotenv need to call only dotEnvConfig function")
		}
	})
}

func Test_GO_ENV(t *testing.T) {
	t.Run("should return the config", func(t *testing.T) {
		goEnv := "development"
		os.Setenv("GO_ENV", goEnv)

		env := NewEnvironment()

		assert.Equal(t, goEnv, env.GO_ENV())
	})
}

func Test_DEV_ENV(t *testing.T) {
	t.Run("should return the config", func(t *testing.T) {
		env := NewEnvironment()

		assert.Equal(t, "development", env.DEV_ENV())
	})
}

func Test_STAGING_ENV(t *testing.T) {
	t.Run("should return the config", func(t *testing.T) {
		env := NewEnvironment()

		assert.Equal(t, "staging", env.STAGING_ENV())
	})
}

func Test_PROD_ENV(t *testing.T) {
	t.Run("should return the config", func(t *testing.T) {
		env := NewEnvironment()

		assert.Equal(t, "production", env.PROD_ENV())
	})
}
