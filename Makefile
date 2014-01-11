all:
	$(MAKE) -C src/sort

test:
	$(MAKE) test -C src/sort

.PHONY: all test
