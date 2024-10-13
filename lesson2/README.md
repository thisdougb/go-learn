## What a model is for

If you try and find out why a software model is useful you’ll find opinions from many points of view. For an infrastructure person, a software model is typically used as a tool to understand how a particular system scales. By contrast, application developers tend to model business processes such as buying and selling widgets.

We have this idea to start a coffee cart business, so we’ll use a software model to find out how many baristas to hire and how many coffees it could sell (a scaling problem). The real value of the model is being able to run many simulations, and try out two, three, four, baristas. Or simulate the effect of a customer order per minute, and then a two orders per minute.

Now, you could do this on paper. Coding is not rocket science (even when it is). But it’s important to understand that the benefit of computers, to systems programmers, is principally around speed and consistency. Doing long calculations on paper is slow, and because you’re human your calculations will very likely contain mistakes. This is why telling computers what to do, and then making them do it at scale, is the smart thing to do.

But first we need to learn some Go basics.

## Plan

In this lesson we'll take advantage of [scope](https://github.com/thisdougb/go-learn/tree/main/lesson1) to focus only on the barista. The barista is a component of the coffee cart system from lesson 1. To make things simpler we’ll use the [Go Playground](https://go.dev/play/) to run code.

In our little simulator we get to decide what a barista is and what they can do. It’s always better to know what you intend to do, before you begin. So let’s write out what our barista will be, in pseudocode:

```
type Barista struct {
  id int    // each barista will be identified by a number
}

// makeOrder simulates the time taken to make the order
func makeOrder() {
  order = pop(orderQueue)
  pause(order.TimeToMakeThis)
}

// acceptOrder accepts a valid order from a customer, and puts it in a queue
func acceptOrder(order) {
  validateOrder(order)
  push(orderQueue, order)
}

// restock replenishes the cart stock from the store room
func restock() {
  decrementStoreStock(sizeOfCartStock)
  pause(timeToRestock)
}
```

Pseudocode is halfway between natural language and computer language. It can’t run, but it does let you clearly set out your thoughts and what the final code will do. It’s one layer deeper towards writing code, still easy to think and talk about but starting to give us an idea of the amount of work involved.

At some point in the future we’ll be able to give AI some pseudocode and it’ll generate a working executable program. Increasing your skill in writing pseudocode will help you in the future, here’s the [wiki page](https://en.wikipedia.org/wiki/Pseudocode) to get started.

I’ve introduced a `queue` and two methods,  `push` and `pop`. A fifo queue operates on the basis of first in, first out. This is exactly how the legendary British social queue system works in every bus stop, retail shop, _insert any arbitrary reason British people join queues_. The person in the queue who joined first, is the first out (into the bus, served in the shop, etc). The `push` function adds something to the end of a queue, and `pop` gets and removes the first thing in the queue. More on queues [here](https://www.macs.hw.ac.uk/~hwloidl/Courses/F28DA/lectures/l05.pdf).

### Now we’re ready to Go

Open up the Go [playground](https://go.dev/play/), and you’ll see this:

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

Click the run button, and the lower section will print some output. The entry point in Go is the `func main()` in  `package main`. This is where a Go program starts executing when you click run.

#### A little about types

Go is statically typed, which means it won’t run if there are type errors. Type errors could be when you try to use a type that doesn’t exist, or try to call a function with the wrong argument types. There are some built-in types, like `int` (a number), `string` (a string of characters). And then there are types you compose yourself, like our `Barista`.

By contrast, Python is dynamically typed. Which means a variable can hold a string, then you can then assign it an integer. A python program will run even when type errors are present. Python was originally a language for quickly building prototypes, where dynamic typing is desirable. Go was built to operate on systems at scale, so static typing is more appropriate.

In the Go playground replace the `main()` function with this:

```
func main() {
	var b Barista
	fmt.Println(b)
}
```

Here were are declaring a variable `b` to be of type `Barista`, and then trying to print it out using the `fmt` package. If you want to find out how a Go package or function works, you should look it up. Here is the [fmt](https://pkg.go.dev/fmt) package.

When we try and run our new program, we get an error. 

```
./prog.go:8:8: undefined: Barista

Go build failed.
```

It’s important when learning a language to be able to read error messages, because beginners improve by understanding where they went wrong. Fortunately in Go the errors are quite easy to read.

The error above says the error is in file`./prog.go`, which is particular to the playground. The next number `:8` tells you which line the error occured on, in that file. Usually this is enough to spot what the problem is. You can ignore the second number, and the last part of the error is the description.

So we can read this as, in `./prog.go` on line `:8` we have an `undefined: Barista` error.

Of course, this is because we haven’t yet defined our Barista type. So let’s do that now. After the import statement declare the Barista type like this:

```
type Barista struct {
  id int
}
```

In Go a custom type is declared with the keyword `type` followed by the name of the type, and then a definition of the type we want, here it's a `struct { }`. Sort of like declaring a blueprint. Our barista is a `struct` with an `id` atribute which itself is of type `int` (a number). A `struct` is the type we use to declare a collection of fields. In programmer-speak we can say our barista type is composed of an `id` field of type `int`.

Because we’re using the Go playground we are declaring our barista type in the main package. But it’s more common in Go to declare a type in the package that it’s principally used in. We’ll cover more on that later. The uppercase B signifies that this type is exported (usable) outside this package, and again we’ll cover that later. We are juggling chickens and eggs here, but all will become clear with a little practice.

Now when we run the program it doesn’t error. It doesn’t do anything yet, but we fixed the missing type:

```
{0}

Program exited.
```

#### It’s Alive!

Remember in lesson 1 when we covered _things_ and _types of things_? Our code so far only declares the `Barista` as the _type of thing_. We haven’t really created any baristas yet.

There are a few ways to create _things_ in Go. Start a new playground sessions in another tab and replace `main()` with this:

```
func main() {
	var x int
	x = 1
	fmt.Println(x)

	y := 1
	fmt.Println(y)
}
```

Here `x` is declared the long way, and `y` the short way. The short way is most commonly used, but we’re going to continue with the long way until it really sinks in. Declare a variable with an explicit type, then set its value. Variables always have a defined type, or the program will not compile.

Back to our barista program, and replace the `main()` function with this:

```
func main() {
	var b Barista
	b.id = 1
	fmt.Printf("%+v\n", b)
}
```

We’re declaring the variable `b` to be an instance of type `Barista`, and then setting the field `id` of this instance `b` to 1.  Ensure the previous sentence really makes sense to you. Look above and describe what we did with `x` in the same plain language style. When you are trying to find errors in your code, being able to talk yourself through your code will help a great deal.

When we now run the program and print the variable `b`  we get `{id:1}`. Look up the [fmt](https://pkg.go.dev/fmt) package and read about the `fmt.Printf()`  function, and find out what `%+v` is doing. Being able to find answers in the documentation is the fastest way to learn/fix things through understanding, rather than copy/paste from StackOverflow.

We now have a barista! Declaring and using a composite type is a significant step in learning Go. Well done. Have a grin, because you’re no longer a complete novice in Go.

#### This barista boogies

Let’s animate our baristas with a boogie function. A _boogie function_ is not the name of some clever computer science algorithm. It is in fact, [this](https://en.wikipedia.org/wiki/Boogie).

Functions in Go are usually quite short and simple, with one catch that we’ll come to. Update your playground code with this:

```
func boogie() {
	fmt.Println("boogie")
}

func main() {
	var b Barista
	b.id = 1
	boogie()
}
```

We have declared a function, with the same broad style as `main()`. A `func` with the name `boogie` which takes no arguments `()` and returns nothing. The code for the function is a single statement inside `{}`.

Inside our `main()` function we call the new function by name followed with `()`, and Go will simply run the `boogie` function code at that point. All simple so far.

To improve this a little we can pass our created barista (the instance, not the type) into this new function as an argument. But first we must ensure the boogie function is expecting an argument of the correct type.

Replace the boogie function with this:

```
func boogie(input Barista) {
	fmt.Printf("barista %d boogies\n", input.id)
}
```

Now inside the `()` we have one argument, which we have named `input` and is of type `Barista`. The `input` is a variable, whose scope is contained within the function. You can create other variables with the name `input` elsehwere in the code and they are completely different to this `input` in this function.

But when you try and run the program, you should get this error:

```
./prog.go:17:2: not enough arguments in call to boogie
	have ()
	want (Barista)
```

Hopefully you can read this in plain language, and the cause will be apparent. But let’s go through it anyway. In file `prog.go` on line `:17` the error description is `not enough arguments in call to boogie`. Further information is given, for this type of error, `have()` and `want (Barista)`.

In Go functions must be called using the correct number and type of arguments. This error tells us that the first (and only in this case) argument is of type `Barista`. So we must call the function with a variable of that type. Update line 17 and call the boogie function using our variable `b` as the argument:

```
	boogie(b)
```

and voila:

```
barista 1 boogies
```

### Knowledged gained so far

We have covered a lot of the things you’ll need to understand to make learning Go a happy experience. Here’s a quick recap:

- how scope works, and Go specific scopes (lesson 1)*
- making a plan before you start writing code*
- writing pseudocode to set out your thinking clearly*
- how to run simple programs in the Go playground
- how to declare composite types in Go (new types made of structs)
- how to declare and call functions
- how to pass arguments to functions
- how to read error messages so we understand what Go is telling us*
- how to lookup the Go package documentation*

If any of the above hasn’t really sunk in yet, don’t worry. Go off and do something else, and come back and re-read it later. 

My wild guess is that the above list covers maybe 20% of learning Go for day to day use. This is a big step forwards. 

I have put an `*` next to the items in the list that are the really valuable things to understand. Approaching a programming task in the right frame of mind is the single biggest power-up you can make to your work. With a little experience and practice these things become second nature rather than explicit steps. And the pay-back over time is huge, as you’ll ultimately be doing more by doing less.

### Homework

1. Write a new function `func sings(input Barista)` which accepts a variable of type `Barista` and prints, “barista 1 sings”
2. Create a second barista (a variable of type `Barista`) and have both baristas boogie then sing.

#### Something more advanced

Modify the sing function to accept two arguments, `func sing(first Barista, second Barista)`.  Call that function from `main()` and make the first barista sing to the second barista. The have the second barista return the gesture.

You should have printed output similar to:

```
barista 1 sings to barista 2
barista 2 sings to barista 1
```
