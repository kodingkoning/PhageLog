package main

// check basic properties of the start and stop values
// input: implicit parameter g, a *Gene
// return: a string explanation of error, or empty string if no error
func (g *Gene) checkSSC() string {
	// if it's forward, then the differnce between stop and start + 1 should be a multiple of three
	// if it's reverse, then the difference between start and stop + 1 should be a multiple of three
	if (g.forward && ( (g.length()%3) != 0)) || (!g.forward && (g.length())%3 != 0) {
		// error("Length must be a multiple of 3.")
		return "Error: Length must be a multiple of 3."
	}
	if g.forward && g.start > g.stop {
		// error("For a forward gene, start must be less than stop."
		return "For a forward gene, start must be less than stop."
	}
	if !g.forward && g.stop > g.start {
		// error("For a reverse gene, stop must be less than start."
		return "For a reverse gene, stop must be less than start."
	}
	return " "
}
