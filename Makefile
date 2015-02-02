GOTESTOPTS ?=
check:
	@find . -iregex '.*/*_test.go' -type f -exec dirname {} \; | sort -u | while read DIR; do (cd $$DIR && go test $(GOTESTOPTS) .); done
.PHONY: check
