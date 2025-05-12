# AMQP

There are practice notes how to send and receive messages via AMQP.

## Setup

Start up ActiveMQ by executing `docker run --name activeMQ -e ARTEMIS_USER=secret -e ARTEMIS_PASSWORD=secret -v path_to_local:/var/lib/artemis-instance --network my_network -p 61616:61616 -p 8161:8161 -d apache/activemq-artemis` command[^1].

## How to Send or Receive Messages

Checkout [provider](./provider/) and [consumer](./consumer/) directories.

## How to Enable SSL on ActiveMQ

In `broker.xml` configuration file of ActiveMQ, append the value `sslEnabled=true;keyStorePath=PATH_TO_KEY_STORE;keyStoreType=YOUR_KEY_STORE_TYPE`  to the `acceptor` node[^2].

Note:

- The file specified by the `keyStorePath` property **MUST** contain both the SSL key and certificate.
- If the file is in `JKS` format, you don't need to explicitly set the `keyStorePath` property.

[^1]: [Official Document-Docker](https://activemq.apache.org/components/artemis/documentation/latest/docker.html#docker)
[^2]: [Official Document-Configuring the Transport](https://activemq.apache.org/components/artemis/documentation/latest/configuring-transports.html#configuring-netty-ssl)
