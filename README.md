# demoChannel

This golang demo illustrates using _goroutines_ with and without _channels.

##### Console Output

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
