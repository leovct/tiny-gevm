package main

import "testing"

func TestAdd(t *testing.T) {
	op := func(evm IEVM) error { return evm.Add() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 5} // 5 = 3+2
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestMul(t *testing.T) {
	op := func(evm IEVM) error { return evm.Mul() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 6} // 6 = 3*2
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestSub(t *testing.T) {
	op := func(evm IEVM) error { return evm.Sub() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 1} // 1 = 3-2
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestDiv(t *testing.T) {
	op := func(evm IEVM) error { return evm.Div() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 2} // 2 = 4/2
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestSDiv(t *testing.T) {
	op := func(evm IEVM) error { return evm.SDiv() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 2} // 2 = 4/2
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.Mod() }
	initialStack := []uint64{1, 5, 12}
	expectedStack := []uint64{1, 2} // 2 = 12%5
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestSMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.SMod() }
	initialStack := []uint64{1, 5, 12}
	expectedStack := []uint64{1, 2} // 2 = 12%5
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestAddMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.AddMod() }
	initialStack := []uint64{1, 7, 2, 15}
	expectedStack := []uint64{1, 3} // 3 = (15+2)%7
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestMulMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.MulMod() }
	initialStack := []uint64{1, 7, 2, 15}
	expectedStack := []uint64{1, 2} // 2 = (15*2)%7
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestExp(t *testing.T) {
	op := func(evm IEVM) error { return evm.Exp() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 16} // 16 = 2**4
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

func TestSignExtend(t *testing.T) {
	op := func(evm IEVM) error { return evm.SignExtend() }
	initialStack := []uint64{1, 0, 4}
	expectedStack := []uint64{1, 4}
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}
