package ui

import "fmt"

func PrintBanner() {
	banner := `
      ___          ___     
     /\__\        /\  \    
    /::|  |      /::\  \    | Welcome to NameSniper
   /:|:|  |     /:/\ \  \   | Version 1.0
  /:/|:|  |__  _\:\~\ \  \ 
 /:/ |:| /\__\/\ \:\ \ \__\ | Intelligence Search Tool
 \/__|:|/:/  /\:\ \:\ \/__/ | Developed by hadnu
     |:/:/  /  \:\ \:\__\   
     |::/  /    \:\/:/  /   | Let's hunt!
     /:/  /      \::/  /   
     \/__/        \/__/    

    -- using Go 1.24.0                                      
`
	fmt.Println(banner)
}

