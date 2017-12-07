# Abyss

This is an attempt to learn go, and get some admin experience running kafka.

There will be two major components, the producer and the consumer.

The producer should listen to messages from a single chat client (slack, keybase, discord...) and play them into a kafka topic in near real time.

The consumer should replay messages from all kafka topics into a single chat client from all of the others.

This *should* allow us to bridge the gap between worlds.
