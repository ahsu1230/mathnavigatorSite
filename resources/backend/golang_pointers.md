# Golang Pointers

All variables and values used by a program are stored somewhere in memory, and a pointer is simply a memory address.

## How do we use them?

The syntax for declaring a pointer of type T is *T. For example:

`var p *int`

The variable `p` is a pointer to an integer value.

`x := 5`
`p = &x`

The following code assigns `p` to the memory address of the variable `x`. The & operator simply returns the memory address of the variable it is used on.

`fmt.Println(*p) // prints 5`

The following code demonstrates the use of the * operator, which returns the value stored at the address of the pointer it is used on. This is known as **dereferencing**.

## Where do we use them in this project?

When we serialize or deserialize objects with `Marshal` and `Unmarshal`, we pass pointers of the objects as arguments. More information on serialization and deserialization can be found [here](json.md).
