/* import.go provides the methods to import phage's genes
 * Author: Elizabeth Koning
 * Written in 2017
 */

package main

import(
  "stringconv"
)

func readNextGene(type input) {
  string result = ""
  while( !input.eof() ) {
    string str = input.getline()
    if( str == "" ) break
    result += str
  }

  //turn this into an array, split at \n chars

  //First line: get two numbers from it, and see if the beginning is CDS Complement or CDS
    //throw error if neither

  //Second line: \t/gene="#"
    //check for right start, then try to turn section between "" into an int

  //Third line: \t/product="gp#"
    //check for start, turn #into int, make sure 

  //Fourth line: \t/locus tag="NAME_#"

  //Fifth line-end: \t\t/note="note contents"
    //Specifically: call strength, SSC, CP, SD, SCS, Gap, BLAST, LO, ST, F, FS, ST, Notes
}
