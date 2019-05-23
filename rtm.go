// Package rtm provides primitives for Intel's Restricted Transactional Memory Operations
package rtm

// TxBegin marks the start of transaction. It will return a status code
// different to TxBeginStarted when the transaction fails.
func TxBegin() TxBeginStatus {
	s := txBegin()
	return {Status(s>>24), s&0xFF}
}

func txBegin() (status uint32)

// TxAbort aborts transaction with the provided reason, that will
// be available in the TxAbortExplicitReason field of the
// TxBeginStatus returned by TxBegin
func TxAbort(reason uint8)

// TxEnd marks the end of transaction
func TxEnd()

// InTx returns true if the processor is executing a transactional region.
func InTx() bool {
	return txTest() == 1
}

func txTest() (status uint8)

type Status uint8

const (
	// TxBeginStarted is returned by TxBegin() when transaction is started
	TxBeginStarted Status = 0
	// TxAbortExplicit bit is set if abort caused by explicit abort instruction.
	TxAbortExplicit = 1 << iota
	// TxAbortRetry bit is set if the transaction may succeed on a retry
	TxAbortRetry
	// TxAbortConflict bit is set if another logical processor triggered a
	// conflict with a memory address that was part of the transaction
	TxAbortConflict
	// TxAbortCapacity bit is set if RTM buffer overflowed
	TxAbortCapacity
	// TxAbortDebug is set if debug breakpoint triggered
	TxAbortDebug
	// TxAbortNested is set if abort occurred in a nested transaction
	TxAbortNested
)

type TxBeginStatus struct {
	Status Status
	TxAbortExplicitReason uint8
}
