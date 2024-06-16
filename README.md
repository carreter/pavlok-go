# Pavlok CLI and API wrapper written in Go
This package provides a simple CLI and Go API wrapper for the [Pavlok](https://shop.pavlok.com/) smartwatch's API.

## CLI Installation
Install Go and run the following:
```shell
go install github.com/carreter/pavlok-go/cmd/pavlok-cli@latest
```

Also ensure that `$GOPATH/bin` is in your `PATH`.

Now you're all set to start zapping away!

## Usage:
Get an API key by signing in to the [Pavlok API docs](https://pavlok.readme.io/reference/intro/authentication)
with the same email as the Pavlok account your device is registered to.

To send a stimulus, run `pavlok-cli [flags] <stimulus type>`.

Allowable stimulus types:
- `zap`: Zaps the user.
- `beep`: Emits a beep.
- `vibe`: Vibrates the device.

Flags:
- `--apikey` (string, required): The API key.
- `--reason` (string): The reason for the stimulus.
- `--value` (int): The value (strength) of the stimulus.
