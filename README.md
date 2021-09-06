# ape-store

- [graphql federation](https://www.apollographql.com/docs/federation/)
    - [go](https://github.com/99designs/gqlgen/tree/master/example/federation)
    - [php](https://github.com/pascaldevink/php-graphql-federation/tree/master/examples/shorthand)
- [Kafka](https://kafka.apache.org/intro):
  - [connect](https://docs.confluent.io/platform/current/connect/index.html)
  - [streams](https://kafka.apache.org/documentation/streams/)


## local development

1. install [tilt](https://docs.tilt.dev/install.html), [helm](https://helm.sh/docs/intro/install/#from-script), [helmfile](https://github.com/roboll/helmfile#installation), [helm diff](https://github.com/databus23/helm-diff#using-helm-plugin-manager--23x), and [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
2. setup [kind with local registry](https://github.com/tilt-dev/kind-local#how-to-try-it)
3. deploy dependencies `helmfile sync`
4. start environment `tilt up`
