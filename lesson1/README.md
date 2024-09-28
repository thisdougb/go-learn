# 1. Speed, systems, and scope

## Speedy farming

To harvest a field of potatoes, potato picking people need to be physically present. The smart farmer obtains technology (potato picking machines), and becomes more productive and is able to harvest more potatos per day. Manual labour is replaced by machine labour. More potatoes in a shorter time with less labour. 

In tech jobs we have an abundance of technology that helps us replace our *time spent* doing tasks, in a similar way. Our individual limiting factor is time to spend (24 hours per day). Spend it wisely, like a smart farmer.

Let's say anything ending in `.tf` or `.yaml` is a digital potato. Find your potatoes:

```
# lists top ten changed files in a git repo in the past 3 months
#   add git optional '--author <your name>' to pull out only your stats
git log --since 3.months.ago --numstat | awk '/^[0-9-]+/{print $NF}' | sort | uniq -c | sort -nr | head
```

In most companies changing a value in a configuration file costs a minimum of one hour of time spent; deciding the value, tickets, pull/push, MR, review, deployment. If your team is doing ten value changes per week, that’s more than one engineering day of time spent.

## Spend your time wisely

You'll go faster by doing less. You'll be more productive by doing less.  Both statements are true, unless you’re picking digital potatoes. 

## Begin at the beginning

For our purposes a system is *a set of things working together*. 

That's both vague and specific.  But this definition allows us to include processes and tasks carried out by people.  These are usually the biggest time-leach in an organisation by far, so they are important.

The infrastructure you manage every day is just a type of system.  You can think of the entire infrastructure, a particular cluster, or a specific Jenkins job. All are systems, differing only in scope.

Other systems could be a train station, a bus stop, a restaurant, a restaurant kitchen.

Also for our purposes scope is *an area of focus*.

Scoping allows us to move efficiently from general to specifc, which is the direction systems people should travel in. It’s just a word that we understand to mean an area of focus within a hierarchy of a system. In AWS Cloud subnets exist within in the scope of the EC2 service.

Supermarkets have scope, because frozen peas make no sense within the bakery section. You look for the freezer aisle, then look for frozen vegetables, then look for peas. You walk the scope hierarchy in your supermarket.

## The coffee cart

There is a coffee cart in our office. Let’s observe it as a system, to provide an example.

First we describe the system we’re observing in terse but relevant language. 

```
Customers join the queue.
Customers order from the menu.
Baristas take orders from customers
Baristas make drinks, and restock the cart with coffee beans, tea, and milk.
```

The words used matter, as we’ll see later. For now, just try and use words that anyone would understand the meaning of. Of course, if the system you’re observing requires specialist knowledge in the first place then use words within that context.

### Things and types of things

The difference between things and types is important to understand, when thinking about systems.

Two statements about the coffee cart system:

- observing the coffee cart we see Rita say hello to Bob and ask him for a latte.
- observing the coffee cart we see a customer join the queue and make an order with a barista. 

In the first statement we are seeing in terms of *instances* of things, in the second we are seeing in terms of *types* of things. Rita is an instance of a `Customer`, a latte is an instance of an `Order`.

In Go we use the keyword `type` to declare types of things. There are some built-in primitive types explained here: [Go Tour types](https://go.dev/tour/basics/11)

We can also build our own types composed of other types. Similar functionality is called a `class` in some other languages. So we could define the blueprint for the Customer type in Go as:

`type Customer { name string, loyaltyStamps int }`

This means every instance of `Customer`, such as Rita, would automatically have some useful attributes. In Go, whenever we are doing something with a customer we can do something like:

`fmt.Printf(“Customer %s has joined the queue\n”, customer.name)`

Upper and lower case use of `customer`  will be covered soon, so don’t worry about it for now.

*Tip: when working on systems if you find yourself dealing with an instance of a thing, then you’re almost certainly not working in a scalable way. When working at scale think and operate on types of things, not things.*

### Actions associated with things

We classify things all the time, into types. If I talk about a wheel, you know the general things a wheel can do. If I then be more specific and mention a bicycle wheel, your understanding of wheel changes but it’s still a type of wheel you think about.

So the type you choose implies its scope as another person would understand it. The attributes and actions of type `Customer` will be far more specific than those of `Person`.  For example, in our coffee cart system, notable actions for a customer may be:

```
func Customer.JoinQueue()
func Customer.CreateOrder()
```

## Fast mental model

From our terse description above we can visualise a mental model of the coffee cart business, in terms of scope. By doing this we’re jumping a few rounds (weeks) of formal software engineering process.

```
- coffeecart              // a Go module
  - business              // a Go package within module coffeecart
    - Cart                // a Go type within package business
      - func Open()       // an action for type Cart
      - func Close()      // an other action for type Cart
  - barista
    - Barista
      - func MakeOrder() 
  - customer
    - Customer
      - func JoinQueue()
      - func CreateOrder()
```

This is my, subjective, mental model of the coffee cart business. Yours will be broadly similar, but may be slightly different in the wording, the actions associated with things, etc. And that’s fine, so long as you can explain why. 

I changed the above model from `Barista.MakeCoffee()` to `Barista.MakeOrder()` after a day, because it seems to make more sense. There’s no absolutely right way to model something, just the simplest and most intuitive hierarchy you can come up with.

## Go scopes

In Go the main scope boundaries are:

```
- module
  - package
    - type
      - function
```

Most programming languages implement scope, but often differently. The transferable knowledge in this lesson is understanding how scopes make systems easier to talk about. We’ll get to why scopes makes implementation easy in later lessons.

## Homework

1. Create a terse description and a fast mental model of bus stop as a system. This shouldn’t take long, and is best done by observing a real bus stop. But don’t creep people out by staring at them.
2. Come back to your model after 24 hours, would you change anything?

