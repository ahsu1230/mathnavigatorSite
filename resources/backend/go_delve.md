# Debugging with Go Delve

https://github.com/go-delve/delve

Delve is a Go debugger tool. It is currently mostly used as a Command-Line debugger - though fancy GUIs may be available: <https://github.com/go-delve/delve/blob/master/Documentation/EditorIntegration.md>

## Starting instructions

Install Delve by following the instructions found here: <https://github.com/go-delve/delve/tree/master/Documentation/installation>. Double check that it is installed correctly by running:

```
dlv version
```

You should see information about the Delve Debugger's Version and Build Id. You can also try 

*There are many ways to use Delve. I will only go through methodologies that will be useful for debugging in orion.* Please use `dlv help` to get a list of all commands that will be useful to you when it comes to debugging individual files / programs.

## The Basics

With Delve, you can do the following things:
- Set breakpoints on function names or file/line numbers
- Run a program and stop at a breakpoint
- Observe values of variables when stopping at a breakpoint
- Clear existing breakpoints
- Re-run process / test

The first thing you have to do with Delve, is to compile a program / test file using Delve using `dlv debug` or `dlv test` (more examples) later. Once you compile using Delve, you'll be in the `dlv` shell like so:

```
$ dlv test program_repo_test.go
Type 'help' for list of commands.
(dlv) 
```

## Breakpoints

Once you're in the Delve CLI (you should see `(dlv)` at the beginning of the statement), you can set breakpoints like this:

```
break <FILE_NAME>.go:<LINE_NUMBER>
break program_repo_test.go:15
```

When you create a breakpoint, a number (id) is assigned to that breakpoint. To delete a breakpoint, clear it by using its assigned id.

```
clear 1
```

Use this to display all break points.

```
breakpoints
```

Use this to clear all break points.

```
clearall
```

## Running a Test with Delve

When in the Delve CLI, run a test using `continue`. The program will run until it hits a breakpoint. When the breakpoint is hit, you may evaluate local variables. Once you're done, use `continue` to run until the next breakpoint or to the end of the program.

When the program is finished, you can `restart` it and prepare to run the test again. Or make a code change and use `rebuild` to build the program again (with new code) and restart the process. Here's an example of the workflow.

```
(dlv) break program_repo_test.go:86         // creates a breakpoint
(dlv) continue                              // starts test

...breakpoint hit...

(dlv) args                                  // prints the arguments of current function
(dlv) locals                                // prints the current local variables
(dlv) print program1.ProgramId              // prints the value of an expression
(dlv) whatis program1                       // prints the type of an expression

(dlv) next                                  // steps to next line in code (can evaluate expressions again)
(dlv) continue                              // done checking stuff, move on to next breakpoint!

...program finishes...

(dlv) restart                               // restart from the top point
(dlv) continue                              // start test again
```

When you are finished with Delve, simply type in `exit` and you should go back to your regular CLI.


## Debugging the `main` package

If you have a file with a `main` package, you can run:

```
dlv debug ______.go 
```

## Debugging a repo_test.go

Example:

```
dlv test ./src/repos/program_repo_test.go
```

## Debugging a controller_test.go

```
dlv test ./src/controllers/program_controller_test.go
```

## Debugging an integration test

```
dlv test ./src/tests_integration
```