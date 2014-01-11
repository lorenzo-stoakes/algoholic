# Insertion Sort, O(n^2) worst-case. Interestingly, O(n) in the case of an already-sorted
# input.
#
# We assign the function to module.exports so consumers can simply write isort = require('isort').
module.exports = (ns) ->
	# The algorithm divides the slice into sorted and unsorted elements. The sorted portion
	# starts trivially with ns[0...1] and we iterate through the remaining unsorted elements
	# one-by-one positioning each into its correct position in the sorted portion.
	#
	# Loop Invariant: ns[0...i] is sorted.
	for key, i in ns[1...]
		# We loop over each of the sorted elements which belong to the 'right' of the key
		# and move them up a place overwriting their neighbour.
		ns[j+1] = ns[j] for j in [i+1..0] when key < ns[j]
		# Finally we have the correct index to place our key value.
		ns[j+1] = key
		# Now the loop invariant is established for ns[0:i+1] and thus we can increment i and
		# loop again.

	# We don't want coffeescript to generate code accumulating a return array.
	return
