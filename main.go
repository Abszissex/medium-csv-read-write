package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main () {

	/**
		// 'input.csv' format. The order of the columns is not important
		// Just make sure that the column name is correct

		Firstname 	| Lastname 	| Country
		---------------------------------
		<value>		| <value>	| <value>
		<value>		| <value>	| <value>
	 */

	// Read persons from 'input.csv'
	persons := readCSVFile("./input.csv")

	// Modify persons a bit and write into new array.
	var modifiedPersons []Person
	for _,person := range persons {
		person.Country = "AnotherCountry"
		modifiedPersons = append(modifiedPersons, person)
	}

	// Write modified persons into 'output.csv'
	writeCSVFile(modifiedPersons, "./output.csv")
}

func writeCSVFile(persons []Person, outputPath string) {

	// Define header row
	headerRow := []string{
		"Firstname", "Lastname", "Country",
	}

	// Data array to write to CSV
	data := [][]string{
		headerRow,
	}

	// Add persons to output data
	for _, person := range persons {
		data = append(data, []string{
			// Make sure the property order here matches
			// the one from 'headerRow' !!!
			person.Firstname,
			person.Lastname,
			person.Country,
		})
	}

	// Create file
	file, err := os.Create(outputPath)
	checkError("Cannot create file", err)
	defer file.Close()

	// Create writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write rows into file
	for _, value := range data {
		err = writer.Write(value)
		checkError("Cannot write to file", err)
	}
}



func readCSVFile(filePath string) (persons []Person) {
	isFirstRow := true
	headerMap := make(map[string]int)

	// Load a csv file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		// Read row
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		checkError("Some other error occurred", err)

		// Handle first row case
		if isFirstRow {
			isFirstRow = false

			// Add mapping: Column/property name --> record index
			for i,v := range record {
				headerMap[v] = i
			}

			// Skip next code
			continue
		}

		// Create new person and add to persons array
		persons = append(persons, Person{
			Firstname:   record[headerMap["Firstname"]],
			Lastname:   record[headerMap["Lastname"]],
			Country: record[headerMap["Country"]],
		})
	}
	return
}

func checkError(message string, err error) {
	// Error Logging
	if err != nil {
		log.Fatal(message, err)
	}
}


type Person struct {
	Firstname string
	Lastname string
	Country string
}