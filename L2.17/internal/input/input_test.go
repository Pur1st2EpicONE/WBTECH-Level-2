package input_test

import (
	"os"
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"

	"L2.17/internal/input"
)

func resetFlags() {
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
}

func TestParseArgs_DefaultPort(t *testing.T) {

	resetFlags()
	os.Args = []string{"telnet", "localhost"}

	data, err := input.ParseArgs()

	assert.NoError(t, err)
	assert.Equal(t, "23", data.Port)
	assert.NotEmpty(t, data.IP)
	assert.Equal(t, "localhost", data.Host)
	assert.Equal(t, 10*time.Second, data.Timeout)

}

func TestParseArgs_CustomPortAndTimeout(t *testing.T) {

	resetFlags()
	os.Args = []string{"telnet", "--timeout", "5s", "localhost", "6969"}

	data, err := input.ParseArgs()

	assert.NoError(t, err)
	assert.Equal(t, "6969", data.Port)
	assert.NotEmpty(t, data.IP)
	assert.Equal(t, "localhost", data.Host)
	assert.Equal(t, 5*time.Second, data.Timeout)

}

func TestParseArgs_InvalidPort(t *testing.T) {

	resetFlags()
	os.Args = []string{"telnet", "localhost", "65536"}

	data, err := input.ParseArgs()

	assert.Nil(t, data)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), input.ErrInvalidPort.Error())

}

func TestParseArgs_HostLookupFailure(t *testing.T) {

	resetFlags()
	os.Args = []string{"telnet", "www.abobus.invalid"}

	data, err := input.ParseArgs()
	assert.Nil(t, data)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), input.ErrHostNameLookUp.Error())

}
