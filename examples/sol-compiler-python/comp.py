
import sys, getopt
import platform

def main(argv):

	try:
		opts, args = getopt.getopt(argv,"ha:b:o:",["inputa=","inputb=","output="])
	except getopt.GetoptError:
		print ('main.py -a <inputfile> -b <inputfile> -o <outputfile>')
		sys.exit(2)

	inputfile1 = ''
	inputfile2 = ''
	outputfile = ''
	output = []
	for opt, arg in opts:
		if opt == '-h':
			print ('test.py -a <firstfile> -b <secondfile> -o <reportfile>')
			sys.exit()
		elif opt in ("-a", "--inputa"):
			inputfile1 = arg
		elif opt in ("-b", "--inputb"):
			inputfile2 = arg
		elif opt in ("-o", "--output"):
			outputfile = arg

	count_of_unneeded_lines = 0
	if platform.system() == 'Darwin':
		count_of_unneeded_lines = 1
	elif platform.system() == 'Linux':
		count_of_unneeded_lines = 5


	input1_lines = []
	input2_lines = []
	back_str = ""
	with open(inputfile1, "r") as input_file:
		cntr = 0	
		for line in input_file:
			if cntr < count_of_unneeded_lines:
				cntr += 1
				continue
			l = line.strip()
			back_str += l
			if l != "" :
				input1_lines.append(line.strip())		
		input_file.close()
	
	with open(inputfile1, "w") as input_file:
		input_file.write(back_str)
		input_file.close()


	with open(inputfile2, "r") as input_file:
		for line in input_file:
			l = line.strip()
			if l != "" :
				input2_lines.append(line.strip())

	if len(input1_lines) != len(input2_lines):
		output = ["different line count in these files! : file {} : #{} , file {} : #{}".format(inputfile1, len(input1_lines), inputfile2, len(input2_lines)) ]
		with open(outputfile, "w") as output_file:
			for item in output:
				output_file.write("{} \n".format(item))		
		return 1

	is_ok = True
	for i in range(len(input1_lines)):
		if input1_lines[i] != input2_lines[i]:
			output.append("difference in line #{}: file: {} -----> content: {} <> file: {} -----> content: {}".format(i+1, inputfile1, input1_lines[i], inputfile2, input2_lines[i]))
			is_ok = False

	if is_ok:
		output.append("OK")

	with open(outputfile, "w") as output_file:
		for item in output:
			output_file.write("{} \n".format(item))

	if is_ok:
		return 0
	else:
		return 1



if __name__ == '__main__':
	result = main(sys.argv[1:])
	sys.exit(result)
