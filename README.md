# A mini-wallet with FALCON-512 signatures

This is a highly experimental external signer with quantum safe signatures.
It is build on liboqs-go. It uses NIST round 2 standartisation level for FALCON.
We will update to the NIST Round 3 standard as soon as it is made available.

> :warning: **Do not use this in production**

Setting up liboqs-go can be difficult. You can run this code with the official
liboqs-go docker container.

```bash
docker pull openquantumsafe/go
```

Mount your current directory into the container and use it as an environment:

```bash
 docker run --rm -it --workdir=/app -v ${PWD}:/app openquantumsafe/go /bin/bash
```

After that you can run the code with:

```bash
 go run main.go
```
