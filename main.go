package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Prime number checker")
	l := widget.NewLabel("Input the number to check the prime number")
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		widget.NewVBox(
			l, e,
			widget.NewButton("click me !", func() {
				n, _ := strconv.Atoi(e.Text)
				arr := PrimeFactor(n)
				primeFactor := ""
				for i := 0; i < len(arr); i++ {
					primeFactor += strconv.Itoa(arr[i]) + " "
				}
				if IsPrime(n) {
					l.SetText(strconv.Itoa(n) + " is prime number")
				} else {
					l.SetText(strconv.Itoa(n) + " is not prime number    |" + strconv.Itoa(n) + " = " + primeFactor)
				}
			}),
		),
	)

	w.ShowAndRun()
}

func IsPrime(n int) bool {
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactor(x int) []int {
	arr := []int{}
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			arr = append(arr, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	arr = append(arr, x)
	return arr
}
