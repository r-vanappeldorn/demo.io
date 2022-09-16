package config

import (
	"log"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"

	"bou.ke/monkey"
)

type TestGetEnvCase struct {
	logFatalCalls int
	title         string
	defineEnv     func()
}

func TestGetEnv(t *testing.T) {
	casses := []TestGetEnvCase{
		{1, "MONGO_URI and DATABASE_NAME not defined", func() {}},
		{1, "MONGO_URI not defined", func() {}},
		{1, "DATABASE_NAME not defined", func() {}},
		{0, "MONGO_URI and DATABASE_NAME defined", func() {
			os.Setenv("MONGO_URI", "test")
			os.Setenv("DATABASE_NAME", "test")
		}},
	}

	for _, c := range casses {
		t.Run(c.title, func(t *testing.T) {
			c.defineEnv()
			called := 0
			fakeFatal := func(msg ...interface{}) {
				snaps.MatchSnapshot(t, msg[0])
				called++
			}

			patch := monkey.Patch(log.Fatal, fakeFatal)
			defer patch.Unpatch()

			NewEnv()
			assertFuncWasCalled(t, c.logFatalCalls, called)
		})
	}
}

func assertFuncWasCalled(t testing.TB, expectedCalled, called int) {
	t.Helper()
	if expectedCalled != called {
		t.Errorf("expected %d amount of calls but got %d", expectedCalled, called)
	}
}
