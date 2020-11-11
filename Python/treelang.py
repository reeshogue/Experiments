class TREE:
	def __init__(self):
		self.tree = {}
		self.pointer = 0, 0

	def step_fore(self):
		self.pointer[0] = not(self.pointer[0])
		self.pointer[1] += 1

	def step_back(self):
		self.pointer[0] = not(self.pointer[0])
		self.pointer[1] -= 1

	def tele_oppo(self):
		self.pointer[0] = not(self.pointer[0])

	def add_point(self):
		self.tree[self.pointer] += 1

	def sub_point(self):
		self.tree[self.pointer] -= 1

	def zero_out(self):
		self.tree[self.pointer] = 0 

class TREEInterpreter:
	def __init__(self):
		self.tree = TREE()
	def interpret(self, file):
		with open(file, 'r') as f:
			string = f.read()
		for i in string:
			if i == '0':
				self.tree.zero_out()
			elif i == '+':
				self.tree.add_point()
			elif i == '-':
				self.tree.sub_point()
			elif i == '*':
				self.tree.tele_oppo()
			elif i == 'v':
				self.tree.step_back()
			elif i == '^':
				self.tree.step_fore()
