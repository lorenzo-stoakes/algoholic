all:
	$(MAKE) -C sort

test:
	$(MAKE) test -C sort

.PHONY: all test
