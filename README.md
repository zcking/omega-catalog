# Omega Catalog

This project is an inspiration I took from Databricks' open sourced [Unity Catalog](https://github.com/unitycatalog/unitycatalog), now part of the Linux Foundation AI & Data projects. 

Originally I just wanted to contribute to that amazing project, and I still intend to. However, I quickly realized there are significant portions that were *not* shared in the open source release (at least not yet), such as support for other cloud storages, generation of temporary storage access credentials, concept of external locations, etc. And in general I felt the tech stack could be improved for maintainability, performance, and extensibility. 

**DISCLAIMER:** Again, please don't take this scoffing at the unity catalog project, I absolutely love the Databricks products, and I think most people don't realize how difficult it is for a company to take something proprietary, deeply-integrated with their other products, and turn around and make it OSS--it probably has more coupling than you'd imagine. Even if this repo just sits on the shelf of my personal GitHub and gets a few stones thrown my way, I'm just happy to explore the idea of it and hopefully others out there will too.

Therefore, I this project is my own version of a Unity Catalog, and built with the following tech stack:

- [Go](https://go.dev/) - programming language of choice
- [gRPC](https://grpc.io) - modern open source high performance Remote Procedure Call (RPC) framework. This allows you to connect and call the APIs from almost any programming language you prefer, rather than integrate REST APIs. HTTP/2 is nice too.
- [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - gRPC to JSON proxy generator, for those cases where REST just won't go away. Completely generative so maintenance is very easy.
- [buf](https://buf.build/docs/introduction) - Protocol buffers build tool, for generating all the protobuf bindings.
- [postgres](https://www.postgresql.org/) - an actual database to persist all the metadata about catalogs, schemas, functions, etc. and will scale to handle more concurrent requests + horizontal deployment scaling.

---

## Getting Started: Docker

To build and run the Docker image, just run:  

```shell
docker compose up --build
```

You should see an output like the following:  

```
docker run --rm -it -p 8080:8080 -p 8081:8081 omega-catalog
{"level":"info","ts":1718525361.8108,"caller":"server/main.go:80","msg":"gRPC server listening on 0.0.0.0:8081"}
2024/06/16 03:09:21 gRPC Gateway listening on 0.0.0.0:8080
```

You can call the REST API to create a catalog like so:  

```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/catalogs' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "demo",
  "comment": "demo catalog",
  "properties": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}'
```

And list catalogs with:  

```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/catalogs/fake' \
  -H 'accept: application/json'
```

## Changing Protobuf

You can change the protobuf at [proto/omega_catalog/v1/omega_catalog.proto](./proto/omega_catalog/v1/omega_catalog.proto). Then use `make generate` to generate all new stubs, which are written to the [gen/](./gen/) directory.
