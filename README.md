# demoChannel

This golang demo illustrates using _goroutines_ with and without _channels.

### Console Output

The demo consists of a main function that calls two goroutines.

##### countdown function
The countdown funtion displays to the console at intervals. When called as a goroutine, the call returns immediately
while the implementation carries out a five-second countdown, displaying at each elapsed second. This function
would operate identically without being called as a goroutine, except that control would not return until the five
seconds elapsed.

##### echo function
The main passes a channel to the echo function. After one second, echo sends its string argument through the channel. 
The main blocks on a receive from the channel, displaying the received value when the goroutine carriews out its
send. 

```

GOROOT=/usr/lib/go-1.10 #gosetup
GOPATH=/home/parallels/go #gosetup
/usr/lib/go-1.10/bin/go build -o /tmp/___go_build_main_go /home/parallels/go/src/demoChannel/main.go #gosetup
/tmp/___go_build_main_go #gosetup
(msg from main: the goroutine countdown is probably executing)

5 ...
4 ...
3 ...
2 ...
1 ...
(msg from countdown; I'm finishing)

Hello, World, via six-second-delayed goroutine channel!

Process finished with exit code 0

```
