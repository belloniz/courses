package subjects

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string)  {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string)  {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string)  {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func compareStrings(str1,str2 string, )  {
	if str1 == str2 {
		fmt.Println("String 1:", str2, "and String 2:",str2,"are equal.")
		return
	}
	fmt.Println("Strings are not equal.")
}

func mutate(s []rune) string  {
	s[0] = 'a'
	return string(s)
}

func PrintStrings(print string)  {
	if print == "yes" {
		fmt.Println("Accessing individual bytes of a string , Accessing individual characters of a string and Runes")
		name := "Hello World"
		fmt.Println("String:", name)
		printBytes(name)
		fmt.Println()
		printChars(name)
		fmt.Println("\n")
		name = "Señor"
		fmt.Printf("String: %s\n", name)
		printBytes(name)
		fmt.Println()
		printChars(name)
		fmt.Println()

		fmt.Println("Accessing individual runes using for range loop")
		charsAndBytePosition(name)
		fmt.Println()

		fmt.Println("Creating a string from a slice of bytes")
		byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
		str := string(byteSlice)
		fmt.Println(str)
		fmt.Println()
		byteSlice = []byte{67, 97, 102, 195, 169}
		str = string(byteSlice)
		fmt.Println(str)
		fmt.Println()

		fmt.Println("Creating a string from a slice of runes")
		runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
		str = string(runeSlice)
		fmt.Println(str)
		fmt.Println(len(name))
		fmt.Println()

		fmt.Println("String length\n")
		word1 := "Senõr"
		fmt.Printf("String: %s\n", word1)
		fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
		fmt.Printf("Number of byte: %d\n", len(word1))
		fmt.Println()

		word2 := "Pets"
		fmt.Printf("String: %s\n", word2)
		fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
		fmt.Printf("Number of bytes: %d\n", len(word2))
		fmt.Println()

		fmt.Println("String comparasion")
		string1 := "Go"
		string2 := "Go"
		compareStrings(string1, string2)
		fmt.Println()

		string3 := "Hello"
		string4 := "World"
		compareStrings(string3, string4)
		fmt.Println()

		fmt.Println("String concatenation")
		string5 := "Go "
		string6 := "is awesome"
		result := string5 + string6
		fmt.Println(result)
		fmt.Println()

		string7 := "Go"
		string8 := "is cool"
		result2 := fmt.Sprintf("%s %s", string7, string8)
		fmt.Println(result2)
		fmt.Println()

		fmt.Println("Strings are immutable, we have to use Runes")
		h := "hello"
		fmt.Println(mutate([]rune(h)))
	} else {
		// Do nothing
	}
}
