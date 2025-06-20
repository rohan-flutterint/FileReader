def combine_files(file_names, output_file):
	try:
		with open(output_file,'w') as outfile:
			for filename in file_names:
				with open(filename, 'r') as infile:
					outfile.write(infile.read())
					outfile.write('\n')
		print("sucessfully wrote files")
	except FileNotFoundError:
		print("files not found")
	except Exception as e:
		print(f"error occurred : {e}")



file_names = ["combined_file_3.txt","combined_file_2.txt"]

output_file = "output.txt"
combine_files(file_names,output_file)