#include "textflag.h"

// func TxBegin() (status uint32)
TEXT ·TxBegin(SB),NOPTR|NOSPLIT,$0
    MOVL $0xffffffff, AX
    XBEGIN fallback // BYTE $0xc7; BYTE $0xf8; LONG $0
fallback:
    MOVL AX, status+0(FP) // AX will be reset on abort
    RET

// func TxEnd()
TEXT ·TxEnd(SB),NOPTR|NOSPLIT,$0
    XEND // BYTE $0x0f; BYTE $0x01; BYTE $0xd5
    RET

// func TxAbort() - this will return always $0xf0 on abort
TEXT ·TxAbort(SB),NOPTR|NOSPLIT,$0
    XABORT $0xf0 // BYTE $0xc6; BYTE $0xf8; BYTE $0x01;
    RET

// func TxTest() (status uint8)
TEXT ·TxTest(SB),NOPTR|NOSPLIT,$0
    XTEST // BYTE $0x0f; BYTE $0x01; BYTE $0xd6
    SETNE status+0(FP)
    RET

// func TxTestAndEnd() (status uint8)
TEXT ·TxTestAndEnd(SB),NOPTR|NOSPLIT,$0-1
    XTEST
    JE notx
    XEND
    MOVB $1, status+0(FP)
    RET
notx:
    MOVB $0, status+0(FP)
    RET
