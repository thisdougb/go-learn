# go-learn

A Go resource for infrastructure people looking to become more productive.

We’ll build a Go app that models a simple coffee cart business. Using this scenario we can work through some of the fundamentals of Go, using enough of the language to get you started as quickly as possible on your own ideas. I’ll also sprinkle on some advanced code to show why Go is the language of choice for infrastructure coding.

With each lesson I’ll try to add a *part 2*, giving a practical example of the lesson targetting some real infrastucture that might actually be useful in the day job.

## From a certain point of view…

Many of the truths we cling to depend greatly on our own point of view. In terms of managing infrastructure, do you know where your own point of view came from? Did you join a DevOps team and become shaped by their existing processes and behaviours? Is StackOverflow your primary learning resource? Can you tell the difference betweek Kool Aid and lemonade, and does it even matter?

Through these lessons I want to give you a glimse of an alternative point of view, with a different set of truths. That of a systems programmer. You can think of this as a DevOps++ engineer, or the initial footsteps on the path to the conceptual 10x.

From this certain point of view you can become more effective at comprehending any system. This leads to quicker and more accurate troubleshooting, detecting low versus high quality decisions, feeling more confident when asserting your technical opinion.

Whether the outcome is that you can write Go code or not is somewhat less important.

## I want you to become safe, competent, and confident (in that order)

The most important thing I want you to learn is how to think differently about managing large systems. Not everything ends in coding, but everything begins by comprehending a system. The faster you can understand what you’re doing, the safer your work will be.

Working with elevated system privileges brings risk, and you can do damage at scale. Competence is required in Go itself, but more importantly in what, how, and where to add safety and quality gates. The journey from incompetent to competent isn’t very long, but is slightly longer than your first attempt to write some Go code. Giving it a fair go is the only ask here.

Whether the the final code is written in Go, Perl, Python, or Rust is less important. Once you comprehend the system problem at hand (make tf plans human-scannable, finding all public Gitlab repos, etc) you’ll be confident in choosing the right tool for the job. Whether that’s Go code, Python script, or clicking through a web UI.

## Why I use the term systems programmer

Application programmers have it easy, from a certain point of view. Consider the narrow scope of their tasks, and the quality and security gates a piece of code must pass through before it has any impact in production. It is actually quite difficult in most companies to release product-fatal application code. Java dev’s simply don’t know the fear of hovering over the return key because their piece of code could break things, a lot of things.

By contrast, systems programming usually involves elevated privileges and targeting important things at scale. So we must not only write the code that does the thing, but also create our own quality and security gates that prevents it doing something unwanted. In addition we must ensure our code operates at scale efficiently, because we mostly interact with rate-limited APIs.

Working on large distributed systems is complex. The only way to manage such systems at scale is to make computers do the work. Use your engineering nous to program the system to do the repetitive toil. That’s the systems programmer way.

## Recommended reading

In three decades of working in computing, I have seen a lot. Not least the cycles and fashions, as the things we used to do are re-discovered as the latest ‘new idea’ to fix everything. For example, containers aren’t a new idea we did them in the 1990s. Developers and Operations people working together didn’t have a name in prior decades, it’s just how work was done.

Right now I can only think of two systems programming books who’s instruction feels timeless. Although there will be plenty out there that I don’t know about. But here you go…

### ZX81 Basic Programming, by Steven Vickers at Sinclair

My first computer was a ZX80 (1980) which came in kit form that you built yourself. Then I was given a ZX81, which came complete and with a manual, 1 KB of memory, but no hard drive.

Chapter two of the ZX81 manual is titled, “Telling the computer what to do.” And this statement really still sums up how we should think about our relationship with computer systems in our day job. The manual is written to you as a person, rather than in an abstract way where the operator doesn’t exist.

The manual goes through how to read/write to the physical memory (pointers). All of the basic programming constructs (if/else, arrays, functions) are explained in familiar terms.

There is a [PDF here](https://www.historybit.it/wp-content/uploads/2018/09/ZX81_BasicProgramming.pdf).

### The C Programming Language, by Kernighan and Ritchie

In the preface they say, “This book is meant to help the reader.” And that comes through in the writing.

C has a reputation of being difficult. But given it’s a low-level langauge with only 32 keywords (C89), writing C isn’t difficult. Compiling and interacting with modern systems is where the challenges come. So don’t be put off.

However, this book is timeless because it has simple tutorials and explanations of the fundamentals of programming (if/else, arrays, functions, data structures, etc). And because there is no longer any industry pressure to learn C (as we’ve moved to higher level languages), reading the book and doing the examples is a little bit like discovering how we got to where we are. And how little things change, pointers are back in fashion!

This book is for the original version of C, so interesting simply if you want to learn good systems programming practice. More [here](https://en.wikipedia.org/wiki/The_C_Programming_Language).
