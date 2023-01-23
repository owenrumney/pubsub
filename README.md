# pubsub

## What is it?

It's a simple pubsub implementation in Go for interacting with the [GCP PubSub Emulator](https://cloud.google.com/sdk/gcloud/reference/beta/emulators/pubsub).

The standard `gcloud pubsub` command [doesn't support using the emulator](https://cloud.google.com/pubsub/docs/emulator?authuser=2#using_the_emulator).

## How do I use it?

### Install

The simplest and only way to install at the moment is to use `go install`... other ways will be added shortly.

```bash
go install github.com/owenrumney/pubsub/cmd/pubsub@latest
```

### Operation?

To get the top level help, use

```bash
pubsub --help

pubsub help
Usage:
  pubsub [command]

Available Commands:
  help          Help about any command
  subscriptions List create and delete subscriptions
  topics        List create and delete topics

Flags:
  -d, --debug               Enable debug logging
  -h, --help                help for pubsub
  -p, --project string      GCP project ID (default "test-project")
  -t, --topic-name string   Topic name

Use "pubsub [command] --help" for more information about a command.
```

All commands require a project ID to be set using `-p` or `--project`. If this isn't done, the default `test-project` is used.

#### Topics

To get help on topics, use

```bash
pubsub topics --help

List create and delete topics

Usage:
  pubsub topics [command]

Available Commands:
  create      Create a topic
  delete      Delete a topic
  list        List topics

Flags:
  -h, --help   help for topics

Global Flags:
  -d, --debug               Enable debug logging
  -p, --project string      GCP project ID (default "test-project")
  -t, --topic-name string   Topic name

Use "pubsub topics [command] --help" for more information about a command.

```

##### Listing Topics

```bash
pubsub topics -p my-test-project list
```

##### Creating Topics

```bash
pubsub topics create -p my-test-project -t my-topic
```

##### Deleting Topics

```bash
pubsub topics delete -p my-test-project -t my-topic
```

#### Subscriptions

To get help on subscriptions, use

```bash
pubsub subscriptions --help

List create and delete subscriptions

Usage:
  pubsub subscriptions [command]

Available Commands:
  create      Create a subscription
  delete      Delete a subscription
  list        List subscriptions

Flags:
  -h, --help   help for subscriptions

Global Flags:
  -d, --debug               Enable debug logging
  -p, --project string      GCP project ID (default "test-project")
  -t, --topic-name string   Topic name

Use "pubsub subscriptions [command] --help" for more information about a command.
```

##### Listing Subscriptions

```bash
pubsub subscriptions -p my-test-project list
```

##### Creating Subscriptions

```bash
pubsub subscriptions create -p my-test-project -t my-topic -s my-subscription
```

When creating a subscription, you can optionally create push endpoint using the `-e` or `--push-endpoint` flag.

##### Deleting Subscriptions

```bash
pubsub subscriptions delete -p my-test-project -s my-subscription
```
