# Cordless

Cordless is supposed to be a custom [Discord](https://discordapp.com) client
that aims to have a low memory footprint and be aimed at powerusers.

This project was mainly inspired by [Southclaws](https://github.com/Southclaws)
[Cordless](https://github.com/Southclaws/cordless), which he sadly didn't
develop any further.

## How to use it

### Installation

First you have to grab the code via:

```shell
go get github.com/Bios-Marcel/cordless
```

In order to execute this command
[you need to have go installed](https://golang.org/doc/install).

In order to execute the application, simply run the executable, which lies at
`$GOPATH/bin/cordless`. In order to be able to run this from your commandline,
`$GOPATH/bin` has to be in your `PATH` variable.

### Configuration

Cordless doesn't need anything besides your Discord token. In order to retrieve
your token, simply follow the steps in the graphic below:

![Steps to retrieve discord token - by ripcord](https://cancel.fm/ripcord/static/app_misc/discord_token_howto_en-US.png)

After retrieving the token, you have to insert it into the `Token` field of
your configuration file. In order to find the location of the configuration
file, simply run cordless via your terminal and check the output, as it will
tell you, where its configuration file lies.