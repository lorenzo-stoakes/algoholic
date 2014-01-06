assert = require('assert')
isort = require('./isort')

# A simple test which sorts a reversed set of integers and ensures they are correctly sorted
# afterward.

N = 1e4
ns = [N..1]

isort(ns)
describe 'insertion sort', ->
	it 'should sort a reverse-sorted list of integers into ascending order', ->
		for n, i in ns
			assert.equal(n, i+1)
