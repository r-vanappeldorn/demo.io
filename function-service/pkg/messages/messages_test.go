package messages

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestError(t *testing.T) {
	snaps.MatchSnapshot(t, Error())
}

func TestSuccess(t *testing.T) {
	snaps.MatchSnapshot(t, Succes())
}