include $(GOROOT)/src/Make.$(GOARCH)

TARG=RPCGame.out
GOFILES=\
	RPCGameMain.go

# LinkFiles=\
#	TestMath.go\
#	TestForSleep.go

#include $(GOROOT)/src/Make.cmd or include $(GOROOT)/src/Make.pkg
include $(GOROOT)/src/Make.cmd

# 
# only link the files that declared under LinkFiles
link: 
	  $(GC) Context.go IRPCState.go RestingState.go AttackState.go MoveState.go PanickState.go
	  $(GC) RPCGameMain.go  

#$(QUOTED_GOBIN)/$(GC) $(LinkFiles)

special:
		make clean
		make link
		make	

	

