// Package rtm provides primitives for Intel's Restricted Transactional Memory Operations
package rtm

// TxBegin marks the start of transaction. It will return a status code
// different to TxBeginStarted when the transaction fails.
func TxBegin() TxBeginStatus {
	s := txBegin()
	return TxBeginStatus(s)
}

func txBegin() (status uint32)

// TxAbort aborts transaction with the specified reason, that will
// be available by calling AbortExplicit on the TxBeginStatus returned
// by TxBegin.
func TxAbort(reason uint8)

// TxEnd commits the transaction.
func TxEnd()

// InTx returns true if the processor is currently executing a transactional region.
func InTx() bool {
	return txTest() == 1
}

func txTest() (status uint8)

// TxBeginStatus encapsulates the statuses returned by TxBegin
type TxBeginStatus uint32
// Started is true if the transaction is started
func (s TxBeginStatus) Started() bool { return s == 0xFFFFFFFF }
// AbortExplicit is true if abort caused by explicit abort instruction
func (s TxBeginStatus) AbortExplicit() (bool, uint8) { return s & (1<<24) != 0, uint8(s & 0xFF) }
// AbortRetry is true if the transaction may succeed on a retry
func (s TxBeginStatus) AbortRetry() bool { return s & (1<<25) != 0 }
// AbortConflict is true if another logical processor triggered a
// conflict with a memory address that was part of the transaction
func (s TxBeginStatus) AbortConflict() bool { return s & (1<<26) != 0 }
// AbortCapacity is true if RTM buffer overflowed
func (s TxBeginStatus) AbortCapacity() bool { return s & (1<<27) != 0 }
// AbortDebug is true if debug breakpoint triggered
func (s TxBeginStatus) AbortDebug() bool { return s & (1<<28) != 0 }
// AbortNested is true if abort occurred in a nested transaction
func (s TxBeginStatus) AbortNested() bool { return s & (1<<29) != 0 }
