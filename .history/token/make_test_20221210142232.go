package token

import (
	"testing"

	"github.com/sRRRs-7/loose_style.git/utils"
)

func TestMaker(t *testing.T) {
	NewPayload(utils.RandomString())
}
