package asciirev

import (
	"fmt"
)

func FS(inputStr, banner string) {
	data, err := FileReader(banner)
	// TODO not use Fatal, os.Exit
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err.Error())
		return
	}
	mapData := AsciiMapper(data)
	output := AsciiPrinter(mapData, inputStr)
	fmt.Println(output)
}
