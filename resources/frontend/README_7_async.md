# Asynchronous Programming

All of the programming you've been doing has probably been, what we call, Synchronous coding.

```
print("First");
print("Second");
print("Third");
```

As you might've expected, the output of these 3 statements occur one after the other. We first print "First", then "Second", and finally "Third". Usually, we are confident that the order of this output would never change.
In other words, each statement is expected to wait until the previous statement has completely finished.

But sometimes, particular coding statements can take a long time. For example, syncing your data with the cloud or retrieving content from a website are operations which would take a few hundred milliseconds. To a computer, that's a VERY long time!

But thankfully, computers are also great at multi-tasking! Your computer contains many mini-workers (called **processes**) that all can perform jobs/tasks at the same time. For example, one worker is syncing your Gmail, another is playing music from Spotify and another is processing what you type with your keyboard. Doing all of these complex tasks on one process would be very hard and inefficient for one worker, so it is best to split the tasks across several workers.

**Asynchronous programming** allows a developer to dictate which tasks to send off to other workers (the more technical term here is **threads**) and which tasks to continue working on with the current worker.

Imagine if we use an asynchronous call on the first print statement.
```
doAsync(print("First"));    // asynchronous call
print("Second");
print("Third");
print("Fourth");
```
With this, we can ask another worker to execute `print("First")`. This example is fairly trivial, but imagine, instead, the statement was a much more complex function.
This is great because another worker will execute the first task while our current worker is unblocked and can continue executing the following statements.

However, this can get tricky because we don't know when `print("First")` will finish.
It could finish instantly before "Second", finish late after "Fourth" or anywhere in between.

*I highly recommend watching the following video for more code examples of asynchronous programming.*

## Video

Here's a video that summarizes this topic fairly well and provides Javascript code to highlight asynchronous programming.
https://www.youtube.com/watch?v=Kpn2ajSa92c

# More on Promises
... coming soon ...

# Orion

For the case of `orion`, most of the asynchronous parts occur with React component lifecycles (componentDidMount(), render(), etc.) and with HTTP requests. Look at `api.js` and `axios`.
