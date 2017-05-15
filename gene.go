package main

import (
	"strconv"
)

type Gene struct {
	// Strings: name, codingPot, blast, function, functionSource, notes
	// int: num, start, stop, sdRank, altGeneMark, altGlimer, lastEnd, lengthRank, stStart
	// bool: forward, geneMark, glimmer
	// float: sdScore
	name string
	num int
	start int
	stop int
	forward bool
	codingPot string
	sdScore float64
	sdRank int
	blast string
	geneMark bool
	glimmer bool
	altGeneMark int
	altGlimmer int
	function string
	functionSource string
	notes string
	lastEnd int
	lengthRank int
	stStart int
}

func (g *Gene) getOutput() string {
	result := g.title() + "\n"
	result += g.SSC() + "\n"
	result += g.CP() + "\n"
	result += g.SD() + "\n"
	result += g.SCS() + "\n"
	result += g.gap() + "\n"
	result += g.BLAST() + "\n"
	result += g.LO() + "\n"
	result += g.ST() + "\n"
	result += g.F() + "\n"
	result += g.FS() + "\n"
	result += g.other() + "\n"
	return result
}

func (g *Gene) getCSV() string {
	return g.name +"," + strconv.Itoa(g.num) + "," + g.SSC() + "," + g.CP() + "," + g.SD() + "," + g.SCS() + ","+ g.gap() + "," + g.BLAST() + "," + g.LO() + "," + g.ST() + "," + g.F() + "," + g.FS() + "\n"
}

// generates the title of the output
func (g *Gene) title() string {
	return "Phage " + g.name + ", Gene #" + strconv.Itoa(g.num) + ":";
}

// generates the start and stop (SSC) description
func (g *Gene) SSC() string {
	result := "SSC: Start: " + strconv.Itoa(g.start) + " Stop: " + strconv.Itoa(g.stop);
	if g.forward {
		result += " (forward).";
	} else {
		result += " (reverse).";
	}
	return result;
}

// adds label to output describing CP coverage
func (g *Gene) CP() string {
	if g.codingPot == "all" {
		g.codingPot = "ORF includes all coding potential shown in GeneMark-CP output."
	}
	return "CP: " + g.codingPot
}

// generate the output describing the SD score and its rank
func (g *Gene) SD() string {
	var rank string
	if g.sdRank == 1 {
		rank = ""
	} else {
		rank = findRank(g.sdRank) + " "
	}
	return "SD: " + strconv.FormatFloat(g.sdScore, 'g', -1, 64) + "; " + rank + "best score."
}

// calculate length of gene
func (g * Gene) length() int {
	if g.stop-g.start > 0 {
		return g.stop-g.start+1
	} else {
		return g.start-g.stop+1
	}
}

// generate the output describing the length and rank of the length
func (g *Gene) LO() string {
	//result := "LO: ";
	//if g.stop-g.start > 0 {
	//	result += strconv.Itoa(g.stop-g.start+1) + " bp; "
	//} else {
	//	result += strconv.Itoa(g.start-g.stop+1) + " bp; "
	//}
	//if (g.lengthRank != 1) {
	//	result += findRank(g.lengthRank) + " "
	//}
	//result += "longest possible ORF."
	//return result
	result := "LO: " + strconv.Itoa( g.length() ) + " bp; "
	if (g.lengthRank != 1) {
		result += findRank(g.lengthRank) + " "
	}
	result += "longest possible ORF."
	return result
}

// generate the output describing Starterator results
func (g *Gene) ST() string {
	if g.stStart == 0 {
		return "ST: Starterator not available."
	} else if g.start == g.stStart {
		return "ST: Starterator agrees with start at bp " + strconv.Itoa(g.stStart)
	} else {
		return "ST: Starterator suggested start at bp " + strconv.Itoa(g.stStart)
	}
}

// generate the ouptut describing the gap or overlap
func (g *Gene) gap() string {
	result := "Gap: ";

	// if first gene, there is no previous gene
	if g.num == 1 {
		result += "No previous gene for first gene."
	} else {
		// if forward, take difference between endLast and start,
		if g.start < g.stop {
			if g.lastEnd >= g.start { // overlap, add one for including both ends
				result += strconv.Itoa( g.lastEnd - g.start + 1 ) + " overlap"
			} else { // gap, subtract one to not include ends
				result += strconv.Itoa( g.start - g.lastEnd - 1 ) + " gap"
			}
		} else {
		// if reverse, take the difference between endLast and start,
			if g.lastEnd >= g.stop { // overlap
				result += strconv.Itoa( g.lastEnd - g.stop + 1 ) + " overlap"
			} else { // gap
				result += strconv.Itoa( g.stop - g.lastEnd - 1 ) + " gap"
			}
		}
		result += " with previous gene."
	}
	return result
}

// generate the output describing the BLAST results
func (g *Gene) BLAST() string {
	// TODO make this more elegant, and different possibilities
	return "BLAST: " + g.blast
}

// generates function description with a label for output
func (g *Gene) F() string {
	// TODO make more elegant, more checks
	return "F: " + g.function
}

// generates function source description with label for output
func (g *Gene) FS() string {
	if g.functionSource == "all" {
		g.functionSource = "Consulted BLAST, Phamerator, and HHpred."
	}
	return "FS: " + g.functionSource
}

// generates the description of the source of the start choice
func (g *Gene) SCS() string {
	if g.glimmer {
		if g.geneMark {
			return "SCS: Glimmer and GeneMark."
		} else if g.altGeneMark == 0 {
			return "SCS: Called by Glimmer, not GeneMark."
		} else {
			return "SCS: Called by Glimmer, GeneMark suggested start at " + strconv.Itoa(g.altGeneMark)
		}
	} else {
		if g.geneMark {
			if g.altGlimmer == 0 {
				return "SCS: Called by GeneMark, not Glimmer."
			} else {
				return "SCS: Called by GeneMark, Glimmer suggested start at " + strconv.Itoa(g.altGlimmer) + "."
			}
		} else {
			return "SCS: No source."
		}
	}
}

// generates notes with the "Other notes: " label for output
func (g *Gene) other() string {
	return "Other notes: " + g.notes;
}
