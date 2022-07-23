import sys, getopt
from collections import namedtuple
import scanner
import my_parser
import cgen

help_message = '''
main.py -i <inputfile> -o <outputfile>
main.py -d [-s] [-p] -i <inputfile>

options for debug mode:
-s :	run scanner (use with -d)
-p :	run parser (use with -d)
'''

def run_scanner(code, output_file):
	err, tokens = scanner.tokenize(code)
	for t in tokens:
		if t.group == "OP_PUNCTUATION" or t.group == "KEYWORD":
			output_file.write(t.value+ "\n")
		else:
			output_file.write(t.group + " " + t.value + "\n")
	if err is not None:
		output_file.write(err + "\n")
	output_file.close()


def main(argv):
	debug = False 
	run_scanner_option = False
	run_parser_option = False


	inputfile = ''
	outputfile = ''
	try:
		opts, args = getopt.getopt(argv,"dhpsi:o:",["ifile=","ofile="])
	except getopt.GetoptError:
		print(help_message)
		sys.exit(2)

	for opt, arg in opts:
		if opt == '-d':
			debug = True
		if opt == '-s':
			run_scanner_option = True
		if opt == '-p':
			run_parser_option = True
		if opt == '-h':
			print (help_message)
			sys.exit()
		elif opt in ("-i", "--ifile"):
			inputfile = arg
		elif opt in ("-o", "--ofile"):
			outputfile = arg

	code = ""
	with open(inputfile, "r") as input_file:
		code = input_file.read()
	
	if debug:
		if run_scanner_option:
			scanner.debug_main(code)
		if run_parser_option:
			my_parser.debug_main(code)
		return
	
	output_file = open(outputfile, "w")

	# phase1
	# only_scanner(code, output_file)

	# phase2
	# can_parse = my_parser.parse(code)
	# if can_parse:
	# 	output_file.write("OK")
	# else:
	# 	output_file.write("Syntax Error")

	# phase3
	mips = cgen.generate_tac(code)
	output_file.write(mips)


if __name__ == "__main__":
	main(sys.argv[1:])
