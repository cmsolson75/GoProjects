# Countdown Timer Project


This project is modivated by learning to understand ANSI Escape Sequences, and more advanced testing through Dependency Injection and Mocks. I also will leverage Concurency to enable multiple timers to run at the same time.


I need to improve with testing code that has a DB attached, this will help me in work and other areas.


## Stages


### MVP
NO TESTS
Just print to the screen with a countdown, no user interface eather.


### Format output
Get the same output space
Add in a minutes conversion from seconds.
This works: So hours is next; this should be similar logic


General Logic: Need to use seconds as the universal conversion: so when you get the input in 00:00:00 format you can say that 1:00 = 60 seconds
01:00 = 60 seconds.

I will need to write a parser for the input, but thats not going to be too bad.
I added in hours as well.

### Input Parser
Libraries: For now just use a raw input; so your test will be "1:00" for 1 minute, 1:00:00 for 1 hour

Out of scope: error handling for tokens, at this stage only valid inputs.

Okay I have this working.

### use ANSI to remove waterfall from the output.


## Library Notes

time library: [time](https://pkg.go.dev/time#section-documentation)
### Clock Types
stack overflow answer: [link](https://stackoverflow.com/questions/3523442/difference-between-clock-realtime-and-clock-monotonic)
Monotonic vs Wall Clock
**CLOCK_REALTIME (Wall Clock):** Machines best guess as to the current wall-clock. time-of,day time. This clock can be changed, ie jump forward and backward. This can be changed by NTP (Network time protocol)[notes](https://ubuntu.com/server/docs/about-time-synchronisation#:~:text=Network%20Time%20Protocol%20(NTP)%20is,to%20set%20its%20own%20clock.)
- **NTP: Newtork Time Protocol (NTP)** is a networking protocol for synchronising time over a network. Client requests a current time from a server, and uses it to set its own clock (WALL CLOCK). There are three tiers of NTP servers; tier one NTP are connected to atomic clocks, while tier two and tier three servers spread the load of actually handling requests across the internet. This client software is more complex that you would think. It must factor in comm delays and adjust the time in a way that does not upset all the other processes that run on the server. 

**CLOCK_MONOTONIC (Monotonic):** Represents the absolute elaped wall-clock tiem since some arbitrary, fixed point in the past. It isn't effected by changes in the system time of day clock.

General Ideas: CLOCK_REALTIME fluxuates through the day, CLOCK_MONOTONIC is a more accurate time passing representation, you cant read it though.

ANSI Escape Sequences: 
- [Colours Article](https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x329.html)
- [Curser Movement](https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html)


Great Tutorial to teach you ANSI escape sequences
- [TUTORIAL](https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html)
- To Get the style I want in this app, I should learn ANSI better.
