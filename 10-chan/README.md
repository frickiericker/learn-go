# Single producer, multiple receivers

Receiver registers itself to Broadcaster via a registration channel:

    Receiver
    Receiver --> Broadcaster
    ...

Broadcaster receives messages from Producer via a channel, and distributes the
received messages to registered Receivers:

                             --> Receiver
    Producer --> Broadcaster --> Receiver
                             --> ...

This looks similar to the star network topology. Other topologies such as tree
and line would be possible; [linked channel][1] is a line network I guess.

[1]: https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/
