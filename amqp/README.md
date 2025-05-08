# AMQP

There are practice notes how to send and receive messages via AMQP.

## Setup

Start up ActiveMQ by executing `docker run --name activeMQ -e ARTEMIS_USER=secret -e ARTEMIS_PASSWORD=secret -v path_to_local:/var/lib/artemis-instance --network my_network -p 61616:61616 -p 8161:8161 -d apache/activemq-artemis` command[^1].

[^1]: [Official Document-Docker](https://activemq.apache.org/components/artemis/documentation/latest/docker.html#docker)

## How to Send or Receive Messages

Checkout [provider](./provider/) and [consumer](./consumer/) directories.
