package res

import (
	"fmt"
	"github.com/clbanning/mxj/v2"
	"os"
)

func ReadFile(path string) mxj.Map {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(xmlFile)
	mv, err := mxj.NewMapXmlReader(xmlFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return mv
}
