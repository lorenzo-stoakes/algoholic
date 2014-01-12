all:
	$(MAKE) -C src/sort/c

test:
	$(MAKE) test -C src/sort/c

.PHONY: all test
