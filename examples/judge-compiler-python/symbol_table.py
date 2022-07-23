from lark.lexer import Token
from utils import SemanticError
from lark.visitors import  Interpreter, Visitor_Recursive, Visitor
from lark import Tree


class Type():
	def __init__(self, name, size=None, arr_type=None, class_ref=None, interface_ref=None):
		self.name = name
		self.size = size
		self.arr_type = arr_type
		self.class_ref = class_ref
		self.interface_ref = interface_ref;

	def are_equal(self, type2):
		if self.name != type2.name:
			return False
		if self.arr_type and not type2.arr_type or\
			not self.arr_type and type2.arr_type:
			return False

		if self.class_ref and not type2.class_ref or\
			not self.class_ref and type2.class_ref:
			return False
		
		if self.class_ref and self.class_ref.name != type2.class_ref.name:
			return False
		
		if self.arr_type:
			return self.arr_type.are_equal(type2.arr_type)
		
		return True

	
	def are_equal_with_upcast(self, type2):
		# return true if self can upcast to type2
		if self.arr_type:
			return self.are_equal(type2)

		if self.class_ref and self.class_ref.can_upcast_to(type2.class_ref):
			return True

		return self.name == type2.name


	def __str__(self) -> str:
		return f"<T-{self.name}-{self.size}-arr:{self.arr_type}-cls:{None if not self.class_ref else self.class_ref.name}>"


class Variable():
	def __init__(self, name= None, type_:Type = None, address = None, size = 0):
		self.name = name
		self.type_ = type_
		self.address = address
		self.size = size



	def __str__(self) -> str:
		return f"<V-{self.name}-{self.type_}-{self.address}-{self.size}>"
	

class Function():
	def __init__(self, name, formals=[], return_type:Type = None, prefix_label = ''):
		self.name = name
		self.return_type = return_type
		self.formals = formals	# array: variable (order is important)
		self.label = name
		self.change_name(name, prefix_label)
		
	def change_name(self, name, prefix_label=''):
		# if name != "main":
		# 	self.label = prefix_label + "func_" + name
		# else:
		# 	self.label = prefix_label + name

		self.label = prefix_label + "func_" + name


	def __str__(self) -> str:
		return f"<F-{self.name}-{self.return_type}-{[a.__str__() for a in self.formals]}>"

class Interface():
	def __init__(self, name, address, member_functions={}):
		self.name = name
		self.address = address
		self.member_functions = {}
		self.prototypes = {}
		self.set_prototypes = {member_functions}

	def get_vtable_size(self): # vtable not included
		size = len(self.member_functions)
		if self.parent:
			size += self.parent.get_vtable_size()
		return size

	def set_prototypes(self, member_functions):
		self.member_functions = member_functions
		self.prototypes = {**member_functions}



class Class():
	def __init__(self, name, address, member_data= {}, member_functions={}, parent=None):
		self.name = name
		self.address = address
		self.member_data = {}
		self.member_functions = {}
		self.fields = {}
		self.access_modes = {}
		self.parent = parent
		self.interfaces = {}
		self.set_fields(member_data, member_functions)

	def get_access_mode(self, name):
		if name in self.access_modes:
			return self.access_modes[name]
		
		if self.parent:
			return self.parent.get_access_mode(name)
	
	def get_object_size(self): # vtable not included
		size = len(self.member_data)
		if self.parent:
			size += self.parent.get_object_size()
		return size


	def get_vtable_size(self): # vtable not included
		size = len(self.member_functions)
		if self.parent:
			size += self.parent.get_vtable_size()
		return size
			

	def can_upcast_to(self, class2):
		# check if self can upcat to class2
		if self.name == class2.name:
			return True
		
		if not self.parent:
			return False
		
		return self.parent.can_upcast_to(class2)


	def get_func_index_offset(self):
		if not self.parent:
			return 0
		return self.parent.get_func_index_offset() + len(self.parent.member_functions)
	
	def get_var_index_offset(self):
		if not self.parent:
			return 0
		return self.parent.get_var_index_offset() + len(self.parent.member_data)
	
	def set_fields(self, member_data, member_functions):
		self.member_data = member_data
		self.member_functions = member_functions
		
		if member_data.keys() & member_functions.keys():
			raise SemanticError(f"class '{self.name}' members must have distinct names")
		
		self.fields = {**member_data, **member_functions}


	def get_func_and_index(self, name, error=True, tree=None):
		if self.parent:
			f, i = self.parent.get_func_and_index(name, error=False)
			if f:
				return (f, i)
		
		if name in self.member_functions:
			index = list(self.member_functions.keys()).index(name)
			return (self.member_functions[name], self.get_func_index_offset() + index)
		
		if error:
			raise SemanticError(f'Function {name} not found in class {self.name}', tree=tree)
		return (None, None)

	def get_var_and_index(self, name, error=True, tree=None):
		if self.parent:
			f, i = self.parent.get_var_and_index(name, error=False)
			if f:
				return (f,i)
		
		if name in self.member_data:
			index = list(self.member_data.keys()).index(name)
			return (self.member_data[name], self.get_var_index_offset() + index)
		
		if error:
			raise SemanticError(f'Variable {name} not found in class {self.name}', tree=tree)
		return (None, None)

	def __str__(self) -> str:
		return f"<C: {self.name}\
			\n\tdata: {[v for v in self.member_data]}\
			\n\tmethods: {[v for v in self.member_functions]}>"
			# \n\tdata: {[v.__str__() for v in self.member_data.values()]}\
			# \n\tmethods: {[v.__str__() for v in self.member_functions.values()]}>"

interface_stack = []
class_stack = []
currect_access_mode = None


last_prime = 1
def get_next_prime():
	global last_prime


	ok = False
	while not ok:
		ok = True
		last_prime += 1
		for i in range(2, last_prime):
			if last_prime % i == 0:
				ok = False
	
	return last_prime




class SymbolTable():
	symbol_tables = []

	def __init__(self, parent=None):
		# self.classes = {}		# dict {name: Class}
		self.variables = {}     # dict {name: Variable}
		self.functions = {}     # dict {name: Function}
		self.types = {}			# dict {name: Type}
		self.prototypes = {}
		self.parent = parent
		SymbolTable.symbol_tables.append(self)


	def find_var(self, name, tree=None, error=True, depth_one=False):
		if name in self.variables:
			return self.variables[name]
		if self.parent and not depth_one:
			return self.parent.find_var(name, tree, error)

		if error:
			raise SemanticError(f'Variable {name} not found in this scope', tree=tree)
		return None

	def find_func(self, name, tree=None, error=True, depth_one=False):
		if name in self.functions:
			return self.functions[name]
		if self.parent and not depth_one:
			return self.parent.find_func(name, tree, error)

		if error:
			raise SemanticError(f'Function {name} not found in this scope', tree=tree)
		return None

	def find_type(self, name, tree=None, error=True, depth_one=False):
		if name in self.types:
			return self.types[name]
		if self.parent and not depth_one:
			return self.parent.find_type(name, tree, error)

		if error:
			raise SemanticError(f'Type {name} not found in this scope', tree=tree)
		return None


	def add_var(self, var:Variable, tree=None):
		if self.find_var(var.name, error=False, depth_one=True):
			raise SemanticError('Variable already exist in scope', tree=tree)
		
		self.variables[var.name] = var

	def add_func(self, func:Function, tree=None):
		if self.find_func(func.name, error=False, depth_one=True):
			raise SemanticError('Function already exist in scope', tree=tree)

		self.functions[func.name] = func

	def add_type(self, type_:Type, tree=None):
		if self.find_type(type_.name, error=False, depth_one=True):
			raise SemanticError('Type already  exist in scope', tree=tree)

		self.types[type_.name] = type_



	def get_index(self):
		return SymbolTable.symbol_tables.index(self)

	def __str__(self) -> str:
		return f"SYMBOLYABLE: {self.get_index()} \
			 PARENT: {self.parent.get_index() if self.parent else -1}\
				 \n\tVARIABLES: {[v.__str__() for v in self.variables.values()]}\
				 \n\tFUNCTIONS: {[f.__str__() for f in self.functions.values()]}"



class ParentVisitor(Visitor):
	def __default__(self, tree):
		for subtree in tree.children:
			if isinstance(subtree, Tree):
				assert not hasattr(subtree, 'parent')
				subtree.parent = tree




### stack contains last type:Type visited, remember to pop from stack
# also remember to push into stack :)
stack = [] 

def IncDataPointer(size):
	cur = SymbolTableVisitor.data_pointer
	SymbolTableVisitor.data_pointer += size
	return cur

class SymbolTableVisitor(Interpreter):
	data_pointer = 0
	"""
	Each Node set it's children SymbolTables. 
	If defining a method make sure to set all children symbol tables
	and visit children. default gives every child (non token) parent
	symbol table
	"""


	def __default__(self, tree):
		for subtree in tree.children:
			if isinstance(subtree, Tree):
				subtree.symbol_table = tree.symbol_table
		
		self.visit_children(tree)
	

	def type(self, tree):
		type_ = tree.children[0].value
		stack.append(Type(type_))


	def function_decl(self, tree):
		# stack frame
		#			------------------- 		
		# 			| 	argument n    |			\
		# 			| 		...		  |				=> caller
		# 			| 	argument 1    |			/
		#			------------------- 
		#  $fp -> 	| saved registers |			\
		#  $fp - 4	| 		...		  |			 \
		#			-------------------				=> callee
		# 			| 	local vars	  |			 /
		# 			| 		...		  |			/
		#  $sp ->	| 		...		  |			
		#  $sp - 4	-------------------

		# access arguments with $fp + 4, $fp + 8, ...


		# check if function is a member function
		function_class = None
		if len(class_stack) > 0:
			function_class = class_stack[-1]

		# access
		global currect_access_mode
		access_mode = currect_access_mode
		if currect_access_mode:
			currect_access_mode = None
		

		# type 
		type_ = Type("void") # void

		if isinstance(tree.children[0], Tree):
			tree.children[0].symbol_table = tree.symbol_table
			self.visit(tree.children[0])
			type_ = stack.pop()

		func_name = tree.children[1].value

		# set formal scope and visit formals
		formals_symbol_table = SymbolTable(parent=tree.symbol_table)
		tree.children[2].symbol_table = formals_symbol_table

		# TODO 
		# not sure what to do here and what types do formals need to be 
		# now they are list of types:Type (but without size)
		sp_initial = len(stack)
		self.visit(tree.children[2])
		formals = []
		
		while len(stack) > sp_initial:
			f = stack.pop()
			formals.append(f)
		
		formals = formals[::-1]

		# Add this to formals and symbol table
		if function_class:
			this = Variable(
				name="this",
				type_=Type(
					name=function_class.name,
					class_ref=function_class
				),
				address= IncDataPointer(4)
				)
			formals = [this, *formals]
			formals_symbol_table.add_var(this)


			# add access_mode
			function_class.access_modes[func_name] = access_mode

		

		# set body scope
		body_symbol_table = SymbolTable(parent=formals_symbol_table)
		tree.children[3].symbol_table = body_symbol_table
		self.visit(tree.children[3])

		# change function label in mips code to not get confused with other functions with same name
		prefix_label = ''
		if function_class:
			prefix_label = "class_" + function_class.name + "_"

		tree.symbol_table.add_func(Function(
				name = func_name,
				return_type = type_,
				formals = formals,
				prefix_label=prefix_label
		),tree)

	

	def variable(self, tree):
		
		# check if variable is a member data
		variable_class = None
		if len(class_stack) > 0:
			variable_class = class_stack[-1]
		
		# access
		global currect_access_mode
		access_mode = currect_access_mode
		if currect_access_mode:
			currect_access_mode = None
		

		tree.children[0].symbol_table = tree.symbol_table
		self.visit(tree.children[0])
		type_ = stack.pop()

		var_name = tree.children[1].value
		

		if variable_class:
			# add access_mode
			variable_class.access_modes[var_name] = access_mode


		var = Variable(
				name=var_name,
				type_=type_,
				address= IncDataPointer(4),
				)

		tree.symbol_table.add_var(var, tree)
		
		# We need var later (e.g. in formals of funtions)
		stack.append(var)


	def array_type(self, tree):
		tree.children[0].symbol_table = tree.symbol_table
		self.visit(tree.children[0])
		mem_type = stack.pop()
		stack.append(Type("array",arr_type = mem_type))


	def if_stmt(self, tree):
		
		# expr
		tree.children[1].symbol_table = tree.symbol_table
		self.visit(tree.children[1])

		# body
		body_symbol_table = SymbolTable(parent=tree.symbol_table)
		tree.children[2].symbol_table = body_symbol_table
		self.visit(tree.children[2])

		# else
		if len(tree.children) > 3:
			else_symbol_table = SymbolTable(parent=tree.symbol_table)
			tree.children[4].symbol_table = else_symbol_table
			self.visit(tree.children[4])


	def while_stmt(self, tree):
		# expr
		tree.children[1].symbol_table = tree.symbol_table
		self.visit(tree.children[1])

		# body
		body_symbol_table = SymbolTable(parent=tree.symbol_table)
		tree.children[2].symbol_table = body_symbol_table
		self.visit(tree.children[2])


	def for_stmt(self, tree):
		# for types: (number is child number)
		#	for (2;4;6) 8
		#	for (2;4;) 7
		#	for (;3;5) 7
		# 	for (;3;) 6

		childs = []
		for subtree in tree.children:
			if isinstance(subtree, Tree):
				childs.append(subtree.data)
			else:
				childs.append(subtree.value)
		
		expr1_num = None
		expr2_num = None
		expr3_num = None
		body_num = None

		# type 1
		if len(childs) == 9:
			expr1_num = 2
			expr2_num = 4
			expr3_num = 6
			body_num = 8

		# type 2
		elif len(childs) == 8 and childs[3] == ';':
			expr1_num = 2
			expr2_num = 4
			body_num = 7

		# type 3
		elif len(childs) == 8 and childs[2] == ';':
			expr2_num = 3
			expr3_num = 5
			body_num = 7

		# type 4
		elif len(childs) == 7:
			expr2_num = 3
			body_num = 6

		
		# expr
		expr_symbol_table = SymbolTable(parent=tree.symbol_table)
		
		if expr1_num:
			tree.children[expr1_num].symbol_table = expr_symbol_table
			self.visit(tree.children[expr1_num])
		
		tree.children[expr2_num].symbol_table = expr_symbol_table
		self.visit(tree.children[expr2_num])
			
		if expr3_num:
			tree.children[expr3_num].symbol_table = expr_symbol_table
			self.visit(tree.children[expr3_num])
			
		# body

		body_symbol_table = SymbolTable(parent=expr_symbol_table)
		tree.children[body_num].symbol_table = body_symbol_table
		self.visit(tree.children[body_num])

	def interface_decl(self, tree):
		# INTERFACE IDENT "{" prototype* "}"
		interface_name = tree.children[1].value
		interface = Interface(
			name=interface_name,
			address=IncDataPointer(4),
		)
		type_ = Type(
			name=interface_name,
			interface_ref = interface,
			size=4
		)
		tree.symbol_table.add_type(type_)

		interface_symbol_table = SymbolTable(parent=tree.symbol_table)

		interface_stack.append(interface)
		for subtree in tree.children:
			if isinstance(subtree, Tree) and subtree.data == 'prototype':
				subtree.symbol_table = interface_symbol_table
				initial_stack_len = len(stack)
				self.visit(subtree)
				while initial_stack_len < len(stack):
					stack.pop()

		interface_stack.pop()
		interface.set_prototypes(
			member_functions = interface_symbol_table.prototypes
		)




	def class_decl(self, tree):
		# CLASS IDENT (EXTENDS IDENT)? (IMPLEMENTS IDENT ("," IDENT)*)?  "{" field* "}"
		

		# TODO extends
		# TODO implements
		# TODO access modes

		# name
		class_name = tree.children[1].value

		parent_name = None

		if len(tree.children) > 2 and isinstance(tree.children[2], Token) and tree.children[2].value == 'extends':
			parent_name = tree.children[3].value


		class_ = Class(
			name= class_name,
			address= IncDataPointer(4),	# this memory will be used for vtable
			parent=parent_name
		)

		type_ = Type(
			name=class_name,
			class_ref=class_,
			size=4
		)

		tree.symbol_table.add_type(type_)
		
		# fields
		class_symbol_table = SymbolTable(parent=tree.symbol_table)

		class_stack.append(class_)
		
		for subtree in tree.children:
			if isinstance(subtree, Tree) and subtree.data == 'field':
				subtree.symbol_table = class_symbol_table
				initial_stack_len = len(stack)
				self.visit(subtree)
				while initial_stack_len < len(stack):
					stack.pop()

		class_stack.pop()

		class_.set_fields(
			member_data=class_symbol_table.variables,
			member_functions=class_symbol_table.functions
		)

		
	def field(self, tree):
		# TODO access mode
		tree.children[0].symbol_table = tree.symbol_table
		tree.children[1].symbol_table = tree.symbol_table

		access_mode = self.visit(tree.children[0])
		
		global currect_access_mode
		currect_access_mode = access_mode

		self.visit(tree.children[1])
		

	def access_mode(self, tree):
		if tree.children:
			return tree.children[0].value
		return 'public'

	

	def statement_block(self, tree):
		new_symbol_table = SymbolTable(parent=tree.symbol_table)
		for subtree in tree.children:
			if isinstance(subtree, Tree):
				subtree.symbol_table = new_symbol_table
				self.visit(subtree)



class TypeVisitor(Interpreter):
	def __default__(self, tree):
		self.visit_children(tree)

	def variable(self, tree):
		global variable_inits_code

		type_ = self.visit(tree.children[0])
		var_name = tree.children[1].value
		variable = tree.symbol_table.find_var(var_name, tree=tree)

		variable.type_ = type_

	def type(self, tree):
		type_name = tree.children[0].value
		type_ = tree.symbol_table.find_type(type_name, tree=tree)

		return type_

	def array_type(self, tree):
		arr_type = self.visit(tree.children[0])
		type_ = Type(name="array", arr_type=arr_type)

		return type_


	def class_decl(self, tree):

		class_name = tree.children[1].value

		class_ = tree.symbol_table.find_type(class_name).class_ref
		
		if class_.parent:
			parent_class = tree.symbol_table.find_type(class_.parent).class_ref
			if not parent_class:
				raise SemanticError("Can only extend from classes")
			
			class_.parent = parent_class
		
		for child in tree.children:
			if isinstance(child, Tree):
				self.visit(child)
		

	def function_decl(self, tree):


		# type
		
		type_ = Type("void")
		if isinstance(tree.children[0], Tree):
			type_ = self.visit(tree.children[0])
		

		# name
		func_name = tree.children[1].value
		function = tree.symbol_table.find_func(func_name, tree=tree)

		function.return_type = type_

				
		for child in tree.children:
			if isinstance(child, Tree):
				self.visit(child)
		

	