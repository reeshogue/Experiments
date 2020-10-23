
#Simple weird little language like BF implemented in Python.
#Just class methods are implemented, but will (maybe?) implement syntax...
#Forget attention, trees are all you need...

class TREE:
	def __init__(self):
		self.tree = {}
		self.pointer = 0, 0

	def step_fore(self):
		self.pointer[0] = not(self.pointer[0])
		self.pointer[1] += 1

	def step_back(self):
		self.pointer[1] -= 1
		self.pointer[0] = not(self.pointer[0])

	def tele_oppo(self):
		self.pointer[0] = not(self.pointer[0])

	def add_point(self):
		self.tree[self.pointer] += 1

	def sub_point(self):
		self.tree[self.pointer] -= 1

	def zero_out(self):
		self.tree[self.pointer] = 0 
