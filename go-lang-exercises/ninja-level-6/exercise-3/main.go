/* Use the “defer” keyword to show that a deferred func runs after the func containing it exits. */

package main

import "fmt"

func createDbConnection() {
	fmt.Println("Openning db connection ...")
}

func closeDbConnection() {
	fmt.Println("DB Connection closed!!!")
}

func saveData() {
	fmt.Println("Writing some data to DB ...")
}

func main() {
	defer closeDbConnection()
	createDbConnection()
	saveData()
}
