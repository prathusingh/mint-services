# Mint

- <b>It is a platform that enables retail investors to trade tech stocks with higher conviction</b>
- <b>The platform will support other domain stocks and options in the future. </b>
- <b> A customer can browse photographers and can interact with their work. </b>

#### Tech Stack

- Web Client uses React, GraphQl, Relay, Webpack
- Backend uses Go microservices
- Uses GRPC and CQRS driven Kafka based intercommunication mechanism
- Uses Postgres SQL and GCP for data storage

##### Microservices

- gateway - endpoint for web client to interact with backend. communicates with user service using grpc.
- user - handles authentication, authorization and user selection of tech stocks
- sourcer - procures all the data relevant to user selections.
- analytics - the engine responsible for prividing insights to user
- chat - enables chatting capabilities with other user(s)
