Hi, this is the documentation of the slack bot using Golang.

There are three different levels on which the app can be created,
1. Interacting with the base APIs
2. Go-lang wrapper to interact with the basic Slack APIs, which gives us access to functions and which can then be used to interact with APIs
3. Since we are building a bot we, want it to have the functionality to be able to communicate with the user and give the responses back, so we'll use a socket based integration.

We'll be using the socket based package on top of the go lang slack wrapper package and then use the functions given to us by that.

We shall be using Slacker.