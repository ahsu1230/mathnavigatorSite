## Introduction to REGEX

What is REGEX? Regular expressions, or REGEX, is a sequence of characters that define a search pattern. You can use REGEXs to identify a string as a URL, an email address, or a phone number, among others. In this product, you will most likely be using REGEXs to identify phone numbers, emails, names, and occasionally dates. 

[REGEX Cheatsheet](https://www.rexegg.com/regex-quickstart.html)  
[regexp Documentation](https://golang.org/pkg/regexp/)

## How to use

You will most likely be using REGEXs in Repo Tests and in Domain Tests. In the domains folder, there is a file called [regex.go](./../constellations/orion/src/domains/regex.go). In here, you will see a list of many REGEXs that are used in the domain files. Taking a look at a domain that uses a regex comparison, say [account.go](./../constellations/orion/src/domains/account.go), you will find a line that includes `regexp.MatchString(REGEX_EMAIL, primaryEmail)`. This will return a boolean if the string `primaryEmail` matches `REGEX_EMAIL`.  

Let's do a quick walkthrough of the regex.
`REGEX_EMAIL` is set to `^[^ ]+@[^ ]+$`  

The initial `^` means beginning of string so we start at that point instead of possibly finding a match somewhere else.  
The `[]` means a character set. Any character that is in this character set will match it.  
The `^` inside the `[]` means NOT. This means that any character NOT inside the character set will match.  
The ` ` after the `^` is a simple space.  
The `+` after the `]` means repetitive (one or more).

Putting this part together `^[^ ]+` we have something that says, from the beginning of the string, one or more characters in a row that are not spaces.

The second part of the regex is `@[^ ]+$`
The `@` is a simple `@` symbol. Be careful when using symbols and backslashes -- if you don't have all of them memorized, you may wind up accidentally using an escape sequence or something else.
We've already done the `[^ ]+`, but as a refresher, it means a sequence of one or more non-space characters.
The `$` means end of string.

Putting the second part together, we have something that says, starting with an `@` sign, we have one or more non-space characters until the end of the string.

Putting both parts together, we have, from the beginning of the string, a sequence of non-space characters until a `@` sign, and then more non-space characters until the end of the string.