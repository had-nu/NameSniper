package uint

import "fmt"

func PrintBanner() {
	banner := `
	 nNNNn  nNNNn       sSSSs
     N//// ///// S/////
     N////n/////nS/////
     N////////////S////
     N////////////S///s
     N////////////S//Ss
     N///nNNNn///S/Ss     | Welcome to NameSniper
     N/// NNN ///S/Ss     | Version 1.0
     N/// NNN ///S//Ss    |
     N/// NNN ///S///Ss   | Intelligence Search Tool
     N/// NNN ///S////Ss  |
     N/// NNN ///S/////S  | Developed by hadnu
     N/// NNN ///S//////  |
     nNNn nNNn  sSSSss    | Let's hunt!
                                      -- Go 1.24
`
	fmt.Println(banner)
}