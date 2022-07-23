from lark import Tree

class SemanticError(Exception):
	def __init__(self, message="", line=None, col=None, tree:Tree=None):
		self.message = message
		
		self.line = 0
		self.col = 0

		if tree:
			self.line = tree.meta.line 
			self.col =tree.meta.column

		if line:
			self.line = line
		if col:
			self.col = col

	def __str__(self) -> str:
		return f"l{self.line}-c{self.col}:: {self.message}"
