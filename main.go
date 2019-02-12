/*
	Practice 3141 written in Go for performance reasons.
*/

package main

import (
	"flag"
	"fmt"
	"math"
)

/*
	Where execution of the program begins.
*/
func main() {

	epsPtr := flag.Float64("eps", 0.01, "The residual value of epsilon to compare against.")
	verbosePtr := flag.Bool("verbose", false, "Prints extra execution info to the terminal.")

	flag.Parse()

	/*
		In each loop we create a list (piList) and fill it with the last two values
		from the summation. Instead of re-running the summation, we will have it return
		the last two values instead of only the last.
	*/

	/*
		Declaring variables in Go is a bit different than in Python. Python runs in an
		interpreter, so it's able to determine what data types to use for your variables
		based on what you assign to those variables. Go, on the other hand, is a complied
		language, so it needs to know what will go in a variable as soon as its created.
		This takes more forethought from the programmer, but the advantage is that both you
		and the computer always know what to expect to find in a variable. This means fewer
		errors on the programmer's part and much faster execution on the program's part.
	*/
	var eps float64 = 1.0
	N := 2
	var piList [2]float64

	fmt.Println("Working...")
	/*
		Emulate the loop in the Python version.
		Go doesn't have while loops, so we adapt its robust for-loop structure.
	*/
	for eps > *epsPtr {
		/*
			Just like in Python, we will continue to loop until we get a value of
			epsilon that is less than the residual we are looking for.
		*/
		N++

		var kStart int
		if N > 3 {
			kStart = N
		} else {
			kStart = 1
		}

		piList = piSummation(N, kStart, piList[1])
		eps = epsilon(piList)

		/*
			This chunk of code is just used for showing the values as they are computed.
			It doesn't do anything unless the '-verbose' flag is passed at run time.
		*/
		if *verbosePtr {
			if N%100 == 0 { // if N is divisible by 100
				fmt.Printf("\nCurrent values at term %d\n", N)
				fmt.Printf("pi(%d) = %.12f\n", N, piList[1])
				fmt.Printf("epsilon = %.12f\n", eps)
			}
		}

	}
	fmt.Println("Done!")

	piVal := truncatePi(piList[1], *epsPtr)

	fmt.Printf("\nValue of epsilon to compare to: %.12f\n", *epsPtr)
	fmt.Printf("Number (N) of terms required to reach that: %d\n", N)
	fmt.Printf("Value of pi(%d): %.12f\n", N, piVal)
}

/*
	This function represents the summation in Practice_3141. Pass in the value of
	N and the computed value of the piSummation(N-1) and piSummation(N) will be returned
	in that order as a float64 array.
*/
func piSummation(N int, kStart int, resultStart float64) [2]float64 {
	/*
		NOTE: To save a significant amount of time, the results from the previous run
		are included and used as the starting point. That way we don't have
		to recompute 99% of the summation each time.

		To do this, we would begin from our previous result and set k to the previous
		value of N.
	*/

	var result [2]float64 // Result will be an array, so we have fewer vaiables
	result[1] = resultStart

	for k := kStart; k <= N; k++ {
		result[0] = result[1]
		kFloat := float64(k)
		result[1] += math.Pow(-1.0, float64(kFloat+1.0)) / (float64(0.5)*kFloat - float64(0.25))
		/*
			That last line is a bit funky because we have to force Go to use high-precision
			floating-point numbers if you are using static numbers (i.e 0.25).
		*/
	}

	return result
}

/*
	This function represents the epsilon equation. The value returned is the
    residual represented as a floating-point number with 64-bit precision.
*/
func epsilon(piList [2]float64) float64 {
	return math.Abs(piList[1]-piList[0]) / piList[0]
}

/*
	Cleans up the value of pi so that it's easy to determine the last valid digit in
	the value.
*/
func truncatePi(piVal float64, residual float64) float64 {
	// Determine how many decimal places are being used by residual
	count := 0
	for residual < 1.0 {
		count++
		residual *= 10
	}
	decimalPlace := math.Pow10(count)

	cleanPi := math.Round(piVal*decimalPlace) / decimalPlace

	return cleanPi
}
