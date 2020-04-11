## The Domain Name System and IP Addresses

### What is an IP Address?
An IP Address is simply a unique numerical identifier to your machine or a network of machines. It is how you can find other machines over the Internet to get their information you are looking for. Similarly, it is also how other machines over the Internet can identify your machine and send the correct information to you. 
An IP Address could look something like this: `172.16.254.1`.
The first number could be an identifier to your global region (US East for example). From there, the next few numbers narrow down your location based on more granular regions.

Type in IP Address into Google and it will show you your IP Address.

### What is the Domain Name System?
The Domain Name System (DNS) is basically a phone book for the internet. When going to websites to find webpages you want, you are actually asking for content from IP Addresses.
For instance, typing in `https://www.facebook.com` doesn't actually mean anything to a computer. Given just this, it has no idea where to look for the information you want.
On the other hand, it is also very annoying to have to remember specific IP Addresses for different websites. Imagine you had to remember to go to address `172.16.254.1` to access Google.

So, the DNS was created to allow people to type human-friendly website names that would then convert to IP addresses.
This is what happens when you access `https://www.facebook.com` or `https://www.google.com`. When you access these websites, the Internet browser converts `www.facebook.com` into its corresponding IP address `172.16.254.1` all under the hood.
Just like how a phone book or your Contacts app on your phone converts names to phone numbers, the DNS converts names to IP addresses.

### What is localhost?
`localhost` is basically the hostname for your own computer. It is used for accessing network services without connecting to the Internet. For example, when developing on Math Navigator, the website that you are working on will be available on `http://localhost:xxxx`. This will allow you to open your website on an Internet browser even if you do not have Internet access!
Turns out `localhost` can also be converted to the IP address `127.0.0.1` (which is used for offline network services).

### What is a port and port number?
Every machine has a large list of ports, which are communication endpoints for the Internet. Sometimes when you connect to certain web applications, a port is used to maintain a communication channel.
For instance, video games, chat applications, or streaming services require a constant communication channel and like to setup the channel with specific port numbers.
Here are a few examples:
 - the computer game League of Legends uses port numbers between 5000 to 5500
 - Discord chat uses port 6463
 - The media player app VLC uses 1234 to stream contents sometimes
 You can see more examples [here](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Registered_ports).
 
One port can only be used by one application at a time. Many applications can use a certain port number as long as they are not attempting to use that port at the same time.

### Port numbers with MathNavigator
For our purposes, when developing on MathNavigator, our server or clients will be on urls `http://localhost:808x`
The web server may be on port `8080` while a website being developed may be launched on port `8081`, for example.

What happens is the web server is launched locally at `http://localhost:8080`.
If also launching the web client, it will be available locally at `http://localhost:8081`.
From there, the web server and web client will be communicating with each other through these port numbers.

The whole point of this is to simulate how a real web server on some IP address will interact with some user's computer or phone somewhere. We are setting up this environment so we can simulate this interaction all on our own computer!
Once you are done developing, it is good practice to stop the web server and clients so those ports become available to other applications and to also free resources for your computer :)
