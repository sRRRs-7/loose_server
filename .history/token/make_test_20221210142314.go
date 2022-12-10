package token

import (
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
)

func TestMaker(t *testing.T) {
	paylaod, err := NewPayload(utils.RandomString(10), 1*time.Second)
}
