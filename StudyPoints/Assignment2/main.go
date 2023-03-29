package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

const crypt1 = "d fdw kdg orwv ri ixq sodblqj zlwk d judb edoo"
const crypt2 = "qr rqh nqrzv li fhdvdu dfwxdoob xvhg wkh nlqg ri hqfubswlrq qrz fdoohg fhdvdu hqfubswlrq. dw wkdw wlph rqob ihz frxog uhdg. wkhuhiruh wkhuh zhuh qr qhhg iru frpsoha dojrulwkpv."
const crypt3 = `d whaw zlwk pruh wkdq rqh olqh.
wkh whaw frqwdlqv frppd,
ixoo vwrs dqg hafodpdwlrq pdun !`

func main() {

	// The three most common letters for all the texts are d, o, w
	// if we assume 'e' is the most common letter, then the key would be 1, since 'd' is 1 to the left of 'e'
	// if we assume 'e' is the second most common letter, then the key would be -10, since 'o' is 10 to the right of 'e'
	// if we assume 'e' is the thrid most common letter, then the key would be -18 or 8, since 'w' is either 18 to the right of 'e' or 8 to the left of 'e' with loop around
	// if we assume 't' is the most commmon letter, then the key would be 16, since 'd' is 16 to the left of 't'

	// if we assume 'a' is the most common letter, then the key would be -3, since 'd' is 3 to the right of 'a'

	encryptionKey := -3 // all letters are moved 3 to the right

	// crypt 1
	fmt.Println("Crypt 1:")
	crypt1Stats := CountNumberOfEachRune(crypt1) // See func.
	crypt1Keys := KeysSortedByValue(crypt1Stats) // See func.
	fmt.Printf("Most common letters in decreasing order: %s\n", string(crypt1Keys))
	decrypted1 := ShiftRunes(crypt1, encryptionKey) // See func.
	fmt.Println(string(decrypted1))

	// crypt 2
	fmt.Println("\nCrypt 2:")
	crypt2Stats := CountNumberOfEachRune(crypt2)
	crypt2Keys := KeysSortedByValue(crypt2Stats)
	fmt.Printf("Most common letters in decreasing order: %s\n", string(crypt2Keys))
	decrypted2 := ShiftRunes(crypt2, encryptionKey)
	fmt.Println(string(decrypted2))

	// crypt 3
	fmt.Println("\nCrypt 3:")
	crypt3Stats := CountNumberOfEachRune(crypt3)
	crypt3Keys := KeysSortedByValue(crypt3Stats)
	fmt.Printf("Most common letters in decreasing order: %s\n", string(crypt3Keys))
	decrypted3 := ShiftRunes(crypt3, encryptionKey)
	fmt.Println(string(decrypted3))
}

// Creates a map where each key is the letter (rune in golang) and each value is the number of times a letter appears in the text
func CountNumberOfEachRune(s string) map[rune]int {
	stats := make(map[rune]int)
	for _, letter := range crypt1 {
		if letter < 'a' || letter > 'z' { // ignore all non small letters
			continue
		}

		stats[letter] += 1
	}
	return stats
}

// Sort all keys in decending order based their value and return a ordred array with the keys
func KeysSortedByValue(cryptMap map[rune]int) []rune {
	keys := maps.Keys(cryptMap)
	quickSort(keys, 0, len(keys)-1, cryptMap)
	return keys
}

func quickSort(a []rune, left, right int, cryptMap map[rune]int) {
	if left >= right {
		return
	}
	i, j := left, right
	pivot := cryptMap[a[(left+right)/2]]

	for i <= j {
		for cryptMap[a[i]] > pivot {
			i++
		}
		for cryptMap[a[j]] < pivot {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	quickSort(a, left, j, cryptMap)
	quickSort(a, i, right, cryptMap)
}

// loop through
func ShiftRunes(s string, seed int) string {
	a := []rune(s)
	for i := range a {
		if a[i] < 'a' || a[i] > 'z' { // ignore all non small letters
			continue
		}
		v := int(a[i]) + seed
		if v < 'a' { // loop back to the end of the alphabet if value becomes less than 'a'
			v += 26
		} else if v > 'z' { // loop back to the start of the alphabet if value becomes greater than 'z'
			v -= 26
		}

		a[i] = rune(v)
	}

	return string(a)
}
