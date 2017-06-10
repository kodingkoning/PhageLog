/* gui.go provides the GUI for the phageLog project using the WALK library
 * Author: Elizabeth Koning
 * Written in 2017
 */
package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//"strings"
	//"fmt"
	"strconv"
	"os"
)

func main() {
	var inTE, outTE, inOther, inBlast, inFunction, inFS, inCP, errorOut *walk.TextEdit
	var inGeneNum, inStart, inStop, inSD, inSDRank, inEndLast, inLengthRank, inST, inTotalGenes, inAltGeneMark, inAltGlimmer *walk.NumberEdit
	var lastLabel *walk.Label
	var inReverse, inGlimmer, inGeneMark *walk.CheckBox
	var gene Gene
	lastLabelText := "End of 0"

	MainWindow{
		Title: "Phage Log",
		MinSize: Size{300, 500},
		Layout: VBox{ },
		Children: []Widget {
			VSplitter{ Children: []Widget{
				Composite{
					Layout: VBox{},
					Children: []Widget {
						Composite { Layout: HBox{}, Children: []Widget{
							Label{Text: "Phage"},
							TextEdit{AssignTo: &inTE},
							Label{Text: "Gene"},
							NumberEdit{AssignTo: &inGeneNum,
								OnValueChanged: func() {
									lastLabelText = "End of " + strconv.FormatFloat(inGeneNum.Value()-1, 'f', -1, 64)
									lastLabel.SetText(lastLabelText)
							}, },
							Label{Text: "of"},
							NumberEdit{AssignTo: &inTotalGenes},
						}, },
						Composite{ Layout: HBox{}, Children: []Widget{
							Label{Text: "Start"},
							NumberEdit{AssignTo: &inStart},
							Label{Text: "Stop"},
							NumberEdit{AssignTo: &inStop},
							CheckBox{Text: "reverse", AssignTo: &inReverse},
						}, },
						Composite{ Layout: HBox{}, Children: []Widget{
							Label{Text: "SD"},
							NumberEdit{AssignTo: &inSD, Decimals: 3},
							Label{Text: "Rank"},
							NumberEdit{AssignTo: &inSDRank},
							Label{Text: "Length rank"},
							NumberEdit{AssignTo: &inLengthRank},
						}, },
						Composite { Layout: HBox{}, Children: []Widget{
							Label{AssignTo: &lastLabel, Text: lastLabelText},
							NumberEdit{AssignTo: &inEndLast},
							Label{Text: "ST Start"},
							NumberEdit{AssignTo: &inST},
						}, },
					}, },
					TabWidget{ ContentMarginsZero: true, Pages: []TabPage{
						TabPage{Title: "Blast", Content: TextEdit{AssignTo: &inBlast}, },
						TabPage{Title: "Functon", Content: TextEdit{AssignTo: &inFunction}, },
						TabPage{Title: "FSource", Content: TextEdit{AssignTo: &inFS}, },
						TabPage{Title: "CP", Content: TextEdit{AssignTo: &inCP}, },
						TabPage{Title: "SCS", Content: Composite { Layout: HBox{}, Children: []Widget {
							HSplitter { Children: []Widget {VSplitter { Children: []Widget {
								CheckBox{Text: "Glimmer", AssignTo: &inGlimmer}, 
								CheckBox{Text: "GeneMark", AssignTo: &inGeneMark},
							}, }, VSplitter { Children: []Widget {
							Label{ Text: "Other GeneMark" },
							NumberEdit{AssignTo: &inAltGeneMark, },
							Label{ Text: "Other Glimmer" },
							NumberEdit{AssignTo: &inAltGlimmer, },
						}, }, }, },
						}, }, },
						TabPage{Title: "Notes", Content: TextEdit{AssignTo: &inOther}, },
					}, },
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			TextEdit {AssignTo: &errorOut, MaxSize: Size{250, 20}, },
			PushButton{
				Text: "Record",
				OnClicked: func() {
					// get all raw values
					gene.name = inTE.Text()
					gene.num = int(inGeneNum.Value())
					gene.start = int(inStart.Value())
					gene.stop = int(inStop.Value())
					gene.sdScore = inSD.Value()
					gene.sdRank = int(inSDRank.Value())
					gene.lastEnd = int(inEndLast.Value())
					gene.lengthRank = int(inLengthRank.Value())
					if inReverse.CheckState() == 1 { // if is checked
						gene.forward = false
					} else {
						gene.forward = true
					}
					gene.stStart = int(inST.Value())
					gene.notes = inOther.Text()
					gene.blast = inBlast.Text()
					gene.function = inFunction.Text()
					gene.functionSource = inFS.Text()
					gene.codingPot = inCP.Text()
					gene.glimmer = inGlimmer.CheckState() == 1
					gene.geneMark = inGeneMark.CheckState() == 1
					gene.altGlimmer = int(inAltGlimmer.Value())
					gene.altGeneMark = int(inAltGeneMark.Value())

					// check start and stop values
					sscCheck := gene.checkSSC()
					if sscCheck != "" {
						errorOut.SetText(sscCheck)
					}

						// get output
						setGeneOutput( outTE, gene )

						// copy the results to the clipboard
						walk.Clipboard().SetText( outTE.Text() )

						// save the info to a file
						f, _ := os.Create( gene.name + strconv.Itoa(gene.num) + ".txt" )
						f.WriteString( gene.getOutput() )
						f.Close()

						/*
						// save info to text file
						c, err := os.Open( gene.name + ".txt")
						if err != nil { // create the file and write header to it
							c, _ = os.Create( gene.name + ".csv")
							c.WriteString("name, index, SSC, CP, SD, SCS, Gap, BLAST, LO, ST, F, FS\n")
						}
						_, err = c.WriteString( gene.getCSV() )
						if err != nil {
							outTE.SetText("CSV error")
						}
						c.Close()
						*/
				}, },
				PushButton {Text: "Clear", OnClicked: func() {
					// Not cleared: name, out of/, 

					// numbers
					if gene.stop > gene.start {
						inEndLast.SetValue(float64(gene.stop))
					} else {
						inEndLast.SetValue(float64(gene.start)) }
					inStart.SetValue(0)
					inStop.SetValue(0)
					inGeneNum.SetValue(inGeneNum.Value()+1) // changed, not cleared
					inST.SetValue(0)
					inSD.SetValue(0)
					inSDRank.SetValue(0)
					inLengthRank.SetValue(0)
					inAltGeneMark.SetValue(0)
					inAltGlimmer.SetValue(0)

					// text
					inCP.SetText("")
					inOther.SetText("")
					inFunction.SetText("")
					inGlimmer.SetCheckState(0)
					inGeneMark.SetCheckState(0)
					inCP.SetText("")
					inBlast.SetText("")
					inFS.SetText("")
					inReverse.SetCheckState(0)
					outTE.SetText("")
					errorOut.SetText("")
				}, },
			// end of button row
		},
	}.Run()
}

func setGeneOutput(outTE *walk.TextEdit, g Gene) {
	outTE.SetText( g.title() + "\n" )
	outTE.AppendText( g.SSC() + "\n" )
	outTE.AppendText( g.CP() + "\n" )
	outTE.AppendText( g.SD() + "\n" )
	outTE.AppendText( g.SCS() + "\n" )
	outTE.AppendText( g.gap() + "\n" )
	outTE.AppendText( g.BLAST() + "\n" )
	outTE.AppendText( g.LO() + "\n" )
	outTE.AppendText( g.ST() + "\n" )
	outTE.AppendText( g.F() + "\n" )
	outTE.AppendText( g.FS() + "\n" )
	outTE.AppendText( g.other() + "\n" )
}
