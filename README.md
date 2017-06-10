# phagegenerecord
PhageLog assists in annotating the genomes of bacteriophages, specifically those found through the SEA-PHAGES program. It is written in Go, using the Walk library for the GUI and works in the Windows OS.

First used at Calvin College in 2017, we encourage other schools to use it as they find useful. We would be glad to hear your thoughts about what you appreciate about the program and what additional features would be useful.

How to use the program:
Those looking for the compiled version should download phageLog.exe and phageLog.exe.manifest. If you place both documents in the same folder and then open phageLog.exe, phageLog and a blank terminal window will open. Do not close the terminal window, as it will also close the phageLog window.
Once opened, it is designed to fit to one side of DNA Master, so they can be viewed side by side. From there, fill out the fields from DNA Master and other programs you are using. A few notable fields are:
- "Rank" refers to the its SD score rank. If the chosen start has the best SD score, enter "1", etc.
- "Length Rank" refers to its length's rank. If this is the longest, enter "1".
- "End of x" is the end of the previous gene. This is used to calculate the overlap or gap. The number is updated as the gene number is changed.
- Text entered in Blast, Function, FSource (Function Source), and Notes goes directly to the output, prepended with the appropriate label.
- In SCS (Start Choice Source): If both Glimmer and GeneMark are checked, it will output "Glimmer and GeneMark". If one is checked, it will output "Called by [program], [other program] suggested start at x." unless the other start would be 0, in which case it outputs "Called by [program], not [other program]. If neither box is checked, it will output as "No source."
- For CP (Coding Potential), if "all" is entered (with no extra newline or space characters), it lists the result as "ORF includes all coding potential shown on the GeneMark CP-output." Otherwise, it presents the result as the text entered in the box.
- When "Record" is clicked, PhageLog generates the output and generates any errors. The output is displayed in the large text box, and errors are shown in the lower text field. A text file named with the phage name and gene number is saved in the same location as PhageLog. Also, the results are copied to the clipboard, so they can be pasted to another document.
- "Clear" empties most fields, leaving the phage name and incrementing the gene number. End of the last is updated to the Stop field.

We plan to develop a tool that reformats the annotations to be pastable in DNA Master's "Documentation" tab. This should be released in time for 2018 genome annotations.
