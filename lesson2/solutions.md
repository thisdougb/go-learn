## Homework


1. Write a new function `func sings(input Barista)` which accepts a variable of type `Barista` and prints, “barista 1 sings”
2. Create a second barista (a variable of type `Barista`) and have both baristas boogie then sing.

## Solution

### Part 1

This is a simple function, which is a copy of the `boogie()` function.

```
func sings(input Barista) {
	fmt.Printf("barista %d sings\n", input.id)
}
```

And the output:


```
barista 1 sings
```

### Part 2

I created a second barista, and refactored the variable naming to be clearer. And the `sings()` function now passes two arguments.

```
func main() {
	var barista1 Barista
	barista1.id = 1

	var barista2 Barista
	barista2.id = 2

	boogie(barista1)
	boogie(barista2)
}
```

Now we have the beginnings of a party:

```
barista 1 boogies
barista 2 boogies
```

### Part 3

The new function requires two named arguments, both of type `Barista`. In Go (and most languages) the name of each argument must be unique. This is because the named arguments are created as variables within the function (scope).

I have used `b1` and `b2` as names for the arguments. This function is very small, and these names are easy to understand in this context. The names you choose for function arguments should be clear and their meaning obvious.

```
func sings(b1 Barista, b2 Barista) {
	fmt.Printf("barista %d sings to barista %d\n", b1.id, b2.id)
}
```

The `sings()` function now takes two arguments, so the program fails to run at this point:

```
./prog.go:23:8: not enough arguments in call to sings
	have (Barista)
	want (Barista, Barista)
```

In file `./prog.go` on line `:23` we have the error `not enough arguments in call to sings`.

```
func main() {
	var barista1 Barista
	barista1.id = 1

	var barista2 Barista
	barista2.id = 2

	boogie(barista1)
	boogie(barista2)

	sings(barista1, barista2)
	sings(barista2, barista1)
}
```

When run, this gives the output:

```
barista 1 boogies
barista 2 boogies
barista 1 sings to barista 2
barista 2 sings to barista 1
```

The full program is now:

```
package main

import "fmt"

type Barista struct {
	id int
}

func boogie(input Barista) {
	fmt.Printf("barista %d boogies\n", input.id)
}

func sings(b1 Barista, b2 Barista) {
	fmt.Printf("barista %d sings to barista %d\n", b1.id, b2.id)
}

func main() {
	var barista1 Barista
	barista1.id = 1

	var barista2 Barista
	barista2.id = 2

	boogie(barista1)
	boogie(barista2)

	sings(barista1, barista2)
	sings(barista2, barista1)
}
```
