# dp-hello-world

Examples of services generated by the dp-cli tool

## Description

This repo contains examples of projects generated by the [dp-cli tool](https://github.com/ONSdigital/dp-cli).

These examples are as follows…

- [dp-hello-world-api](https://github.com/ONSdigital/dp-hello-world/tree/master/dp-hello-world-api) - A generated api service
- [dp-hello-world-controller](https://github.com/ONSdigital/dp-hello-world/tree/master/dp-hello-world-controller) A generated controller service
- [dp-hello-world-event](https://github.com/ONSdigital/dp-hello-world/tree/master/dp-hello-world-event) A generated event driven service
- [dp-hello-world-go-library](https://github.com/ONSdigital/dp-hello-world/tree/master/dp-hello-world-library-go) A generated Go library
- [dp-hello-world-js-library](https://github.com/ONSdigital/dp-hello-world/tree/master/dp-hello-world-library-js) A generated JS library

To regenerate these projects, run:

```sh
    make all
```

Otherwise for individual projects:

API:

```sh
    make generate-api
```

Controller:

```sh
    make generate-controller
```

Event-driven:

```sh
    make generate-event
```

Go library:

```sh
    make generate-library-go
```

Javascript library:

```sh
    make generate-library-js
```
