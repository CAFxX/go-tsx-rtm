#include "textflag.h"

// func txBegin() (status uint32)
TEXT ·txBegin(SB),NOPTR|NOSPLIT,$0-16
    MOVL $0xffffffff, AX 
    XBEGIN fallback
fallback:
    MOVL AX, status+0(FP) // AX will be reset on abort
    RET

// func TxEnd()
TEXT ·TxEnd(SB),NOPTR|NOSPLIT,$0
    XEND
    RET

// func TxAbort(reason uint8)
TEXT ·TxAbort(SB),NOPTR|NOSPLIT,$0-16
    MOVB reason+0(FP), AL
    XABORT AL
    RET

// func txTest()
TEXT ·txTest(SB),NOPTR|NOSPLIT,$0
    XTEST
    SETNE status+0(FP)
    RET
