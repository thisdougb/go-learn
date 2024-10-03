## Homework

1. Create a terse description and a fast mental model of bus stop as a system. This shouldn’t take long, and is best done by observing a real bus stop. But don’t creep people out by staring at them.
2. Come back to your model after 24 hours, would you change anything?

## Solution

### Part 1

I sat down and did this in a cafe, watching a bus stop outside on the street.

```
People arrive at the bus stop and join a queue
Buses arrive at the bus stop.
People get off the bus.
People get on the bus.
Buses leave the bus stop.
```

My quick mental model would be:

```
- busstop
  - stop
    - Stop
      - id : string
      - location : string
      - timetable : list
  - bus
    - Bus
      - id : string
      - destination : string
      - route : list of stops
      - capacity : number
      - func StopIsOnRoute()
      - func ArriveAtStop()
      - func DepartStop()
  - person
    - Person
      - func JoinQueue()
      - func IsMyBus()
      - func BoardBus()
      - func ExitBus()
```


### Part 2

A day or so later I made these changes. 

- removed `Bus.destination` because this is simply the last item in Bus.route
- moved `JoinQueue()` action to the `Stop` type, because the queue feels more like a function of the bus stop
- added `LeaveQueue()` to the `Stop` type, because people also leave the queue
- added `IsMyStop()` to the `Customer` type, because people decide when to leave the bus
- as a result I added a `queue` attribute to the `Stop` type
- I added a list of passengers to the `Bus` type, because we need to know when capacity is reached

My new mental model looks like this:

```
- busstop
  - stop
    - Stop
      - id : string
      - location : string
      - timetable : list of routes
      - queue : list of people
      - func JoinQueue()
      - func LeaveQueue()
  - bus
    - Bus
      - id : string
      - route : list of stops
      - capacity : number
      - passengers : list of people
      - func StopIsOnRoute()
      - func ArriveAtStop()
      - func DepartStop()
  - person
    - Person
      - func IsMyBus()
      - func BoardBus()
      - func IsMyStop()
      - func ExitBus()
```

This model is scoped down to a bus stop. But you could easily transform this to model a full bus network.  Adding a bus station package could implement routes and timetables, and allocate them to instances of buses.

Working in scopes allows us to extend functionality easily, often without touching existing code. The Go package structure makes this particularly easy to do.

