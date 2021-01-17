CFLAGS=-std=c11 -g -static
SRCS=$(wildcard *.c)
OBJS=$(SRCS:.c=.o)

myls: $(OBJS)
	$(CC) -o myls $(OBJS) $(LDFLAGS)

clean:
	rm -f myls *.o *~ tmp*

.PHONY: clean
