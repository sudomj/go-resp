package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1bdev/go-resp/pkg/types"
)

func TestNewInstruction(t *testing.T) {

	instruction := types.NewInstruction('*', '4')

	assert.NotNil(t, instruction)
}

func TestNewCommand(t *testing.T) {

	command := types.NewCommand()

	assert.NotNil(t, command)
}

func TestNewToken(t *testing.T) {

	token := types.NewToken('$', 3, []byte("SET"))

	assert.NotNil(t, token)
}
