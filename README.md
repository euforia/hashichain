# hashichain

A toolchain for the hashi stack.

## Features

### Nomad

- Write nomad job spec in HCL
- Convert docker-compose file/s to a nomad job spec

### Vault

- Copy all key-value pairs under a prefix
- Copy keys under a prefix
- Retrieve all key-value pairs under a prefix

## Compose to Nomad

There are a couple of things to be aware of when converting compose files to a nomad
job spec:

- All services are put in a single task group (except datastores).
- Each identified database is put in its own task group.  This is to ensure the database is not re-deployed when other other components of the application are updated / re-deployed.
