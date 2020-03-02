# Mocking

The "mocking" is a common technique for unit tests. Remember how we mentioned multiple layers in the webserver and how unit tests are meant to ONLY test one layer at a time?
However, sometimes this is hard because one layer's logic might depend on another layer's function.

For instance, let's say we have a controller which calls a method from a service.
In controller:
```
func someControllerMethod() {
    ...
    result := service.doSomething1()
    if (result) {
        ...
    }
    ...
}
```

It would be hard to unit test the function `someControllerMethod()` because it is calling service's `doSomething()` method which could have a lot of other logic. We want to keep a controller's unit tests focused ONLY on the controller. We can solve our problem by "mocking" the service.

## Mocking in Golang (Interfaces)
Remember interfaces from school or from other languages? Golang has the same.
An interface is a "class" that you can't instantiate and has no code. All it has are method signatures.
A class implementing an interface must fill in code for all the interface methods.

With interfaces, we can define one interface and have multiple implementations of that interface. For instance,

We can define an Interface called `serviceInterface` which has 2 methods.
```
type serviceInterface interface {
    doSomething1()
    doSomething2()
}
```

We can now define two structs that implement this interface. One implementation `realService` fills in the code for the 2 methods and does what we expect it do (normal code). Another implementation `fakeService` fills in the code for the 2 methods as well, but it has no real logic. It can simply return a hard-coded value.
Here's an example: In this example, we are creating `fakeService` which implements the interface methods from `serviceInterface`. Notice how these methods have almost no code and don't do much.
```
type fakeService struct {
    ...    
}

func (fs *fakeService) doSomething1() {
    return 10
}

func (fs *fakeService) doSomething2() {
    return "bacon"
}
```

When it comes to unit testing the controller, we don't really care much about the service's logic. In fact, maybe the service has a whole bunch of other dependencies that could really make this a lot harder. Instead, we are pretending what the service will do. We are pretending that after service does all of its processing, it will just return 10 or "bacon". This is the technique known as "mocking".

Let's go back to our controller. `service` in this code is the `fakeService` with the mocked methods. Now we can easily unit test controller because `service` methods are much simpler and return expected results.
```
var service = fakeService
// var service = realService // NOT USED FOR TESTING
...
func someControllerMethod() {
    ...
    result := service.doSomething1() // result becomes 10
    if (result) {
        ...
    } else {
        result2 := service.doSomething2() // result2 becomes "bacon"
    }
    ...
}
```

But if we want to use our code to actually function (like back in the real application), we can simply use the other `service`. Notice, how the function for `someControllerMethod()` stays exactly the same, we've simply switched out the services.
```
// var service = fakeService // NOT USED FOR PRODUCTION
var service = realService
...
func someControllerMethod() {
    ...
    result := service.doSomething1()                // is a real value
    if (result) {
        ...
    } else {
        result2 := service.doSomething2()           // is also a real value
    }
    ...
}
```

When setting up our application, we simply default to `realService` when working for real and for unit tests, use `fakeService`. We use this mocking interface technique for most of the different layers in the codebase.