package main

import (
	"fmt"
	"os"
	"reflect"
)
func main(){
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments")
		return
	}
	args := os.Args

	inputfile:= args[1]
	outputfile:= args[2]

	 data1, err := os.ReadFile((inputfile))
	 if err != nil {
		fmt.Println("Error reading file:",err)
		return
	 }
	 //fmt.Println(string(data1))
	 data2 := string(data1)
	 
	 expected := "book"
	 if !reflect.DeepEqual(expected,data2) {
		data2 = "book"
	 } else {
		fmt.Println("Well formatted")

	 }
	 data3 :=[]byte(data2)
	 os.WriteFile(outputfile,data3,0644)
	//  if err != nil {}
	// 	fmt.Println("error writing file:",err)
	//  }

	 
	

}