isort = require('./isort')

# A simple test which sorts a reversed set of integers and ensures they are correctly sorted
# afterward.

N = 1e4
ns = [N..1]

isort(ns)

for n, i in ns when n != i+1
	throw new Error("Index #{i} not sorted, got #{n} expected #{i+1}.")
