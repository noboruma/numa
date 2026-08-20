#include "textflag.h"

// long vdsoGetCPU(unsigned *, unsigned *, void *)

TEXT ·getCPUFast(SB),NOSPLIT,$32-16
	RDTSCP
	MOVL    CX, AX
	SHRL    $12, AX
	ANDL    $4095, CX

	MOVQ    CX, cpu+0(FP)
	MOVQ    AX, node+8(FP)
	RET
