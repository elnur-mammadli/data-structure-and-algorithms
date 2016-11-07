import unittest
def bubble_sort(array, order_type):
	if type(array) is list:
		for i in range(len(array)):
			for j in range(i+1, len(array)):
				if order_type == 'ASC':
					if array[j] < array[i]:
						array[j], array[i] = array[i], array[j]
				elif order_type == 'DESC':
					if array[j] > array[i]:
						array[j], array[i] = array[i], array[j]
	return array

class BubbleSortTest(unittest.TestCase):
	def test_sorted(self):
		self.assertEqual(bubble_sort([5,4,2,12,32], 'ASC'), [2, 4, 5, 12, 32])
		self.assertEqual(bubble_sort([5,4,2,12,32], 'DESC'), [32, 12, 5, 4, 2])

if __name__ == "__main__":
	unittest.main()