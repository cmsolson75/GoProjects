# Weight Converter Project

## Notes

I am reviewing the bufio library
- I think this stands for Buffered IO

Reading Source: [stackoverflow answer](https://stackoverflow.com/questions/1450551/buffered-vs-unbuffered-io#1450563)
What is buffered vs unbuffered IO.
A buffered IO operation is excepting all of a input bytestream without writing each one to disk, if we made it unbuffered this would slow down the processing but if the system dumps core then you would have the output writen to the disk. On the other hand with a buffer we get faster performance and is ideal for most system operations.

Next Reading Source: [stackoverflow buffio](https://www.reddit.com/r/golang/comments/15r962b/when_should_you_consider_using_bufio/)
When the data is very small using bufio won't be useful, but its good when you don't want to keep everything in memmory (ISNT A BUFFER HELD IN MEMORY)

Readings: 
[bufio docs](https://pkg.go.dev/bufio#section-documentation)

