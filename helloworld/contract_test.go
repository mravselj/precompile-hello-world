// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package helloworld

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = vmerrs.ErrOutOfGas
	_ = big.NewInt
	_ = common.Big0
	_ = require.New
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	tests = map[string]testutils.PrecompileTest{
		"calling sayHello from NoRole should succeed": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output string // CUSTOM CODE FOR AN OUTPUT
				output = ""       // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackSayHelloOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: SayHelloGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling sayHello from Enabled should succeed": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output string // CUSTOM CODE FOR AN OUTPUT
				output = ""       // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackSayHelloOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: SayHelloGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling sayHello from Admin should succeed": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output string // CUSTOM CODE FOR AN OUTPUT
				output = ""       // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackSayHelloOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: SayHelloGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"insufficient gas for sayHello should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SayHelloGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"calling setGreeting from NoRole should fail": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput string
				testInput = ""
				input, err := PackSetGreeting(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedErr: ErrCannotSetGreeting.Error(),
		},
		"calling setGreeting from Enabled should succeed": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput string
				testInput = ""
				input, err := PackSetGreeting(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {
				// this function does not return an output, leave this one as is
				packedOutput := []byte{}
				return packedOutput
			}(),
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling setGreeting from Admin should succeed": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput string
				testInput = ""
				input, err := PackSetGreeting(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {
				// this function does not return an output, leave this one as is
				packedOutput := []byte{}
				return packedOutput
			}(),
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"readOnly setGreeting should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput string
				testInput = ""
				input, err := PackSetGreeting(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"insufficient gas for setGreeting should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput string
				testInput = ""
				input, err := PackSetGreeting(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetGreetingGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
	}
)

// TestHelloWorldRun tests the Run function of the precompile contract.
func TestHelloWorldRun(t *testing.T) {
	// Run tests with allowlist tests.
	// This adds allowlist tests to your custom tests
	// and runs them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.RunPrecompileWithAllowListTests(t, Module, state.NewTestStateDB, tests)
}

func BenchmarkHelloWorld(b *testing.B) {
	// Benchmark tests with allowlist tests.
	// This adds allowlist tests to your custom tests
	// and benchmarks them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.BenchPrecompileWithAllowList(b, Module, state.NewTestStateDB, tests)
}
