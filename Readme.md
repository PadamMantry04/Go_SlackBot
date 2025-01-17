Hi, this is the documentation of the slack bot using Golang, Gemini and Slack

There are three different levels on which the app can be created,
1. Interacting with the base APIs
2. Go-lang wrapper to interact with the basic Slack APIs, which gives us access to functions and which can then be used to interact with APIs
3. Since we are building a bot we, want it to have the functionality to be able to communicate with the user and give the responses back, so we'll use a socket based integration.

We'll be using the socket based package on top of the go lang slack wrapper package and then use the functions given to us by that.

We shall be using Slacker.

Moreover, for working on the responses, we shall be using Gemini AI API.

We are using a go wrapper for the Gemini API.

Additionally we shall store the chats in a interactive history struct to maintain a memory.

The code flow is as follows, the main.go is the driving program for our bot. It calls the Logger Initializer, LoadEnv, as well as sets up a gemini chat session. It also starts a new slack bot instance and starts listening actively for new requests to respond to.

The code has been decentralized or modularized for ease of understanding. 

Thanks.

