package main
import(
	"fmt"
	"os"
)
func commandExit(_ *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}