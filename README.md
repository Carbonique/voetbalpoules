# About

Life is busy enough already, so why check [voetbalpoules.nl](https://voetbalpoules.nl) daily for seeing how you and your friends are doing in predicting football matches?

This application helps you by sending your and your friends predictions to a Telegram chat. 

Sit back and enjoy the show!

## Supported data

1. [The current ranking for a pool](assets/telegram-stand.png)
2. [Predictions for upcoming matches played today or tomorrow](assets/telegram-voorspellingen.png)
3. Predictions and results for matches played today or yesterday

## Usage

1. Install either the Docker container or the Go binary
2. Pass configuration as described in [Passing configuration](#passing-configuration)
3. Run the application from the CLI:
 
docker run: e.g: (assuming the .env file is used for configuration) `docker run -it --env-file=./.env ghcr.io/carbonique/voetbalpoules:latest stand`

Or 

`voetbalpoules stand`

### Passing configuration

Different possibilities exist for passing configuration:

| Variable     | Description                   | Default                  | Instructions                                                                                                                                              |
|--------------|-------------------------------|--------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| `BASE_URL`   | Voetbalpoules                 | https://voetbalpoules.nl |                                                                                                                                                           |
| `TOKEN`      | Telegram chat token           | None                     | [link](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)                                                              |
| `CHAT`       | Telegram chat id              | None                     | [link](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)                                                              |
| `COMPETITIE` | Voetbalpoules competitie name | None                     | 1. Go to https://www.voetbalpoules.nl/wedstrijd/index 2. Select the competition you want 3. `COMPETITIE` is the part in the url after `/wedstrijd/index/` |
| `POOL_ID`    | Voetbalpoules pool id         | None                     | 1. Go to https://www.voetbalpoules.nl/stand/poules 2. Select the pool you want 3. `POOL` is the integer in the url after `/stand/poules/`                 |

#### 1. Env file

Create an `.env` file containing the variables above as key-value pairs (`POOL_ID=123`). The file should reside in the same directory as the binary.

Or pass the env-file to docker: 
`docker run --env-file=./.env`

#### 2. Env vars

Environment variables take precendence over an `.env` file.
Set Environment variables in Docker on in the local shell

#### 3. Flags 

Flags take precedence over all environment variables.

Run `--help` for more information.

## Technical

### voetbalpoules-client

A voetbalpoules-client can be found in `pkg/voetbalpoules-client`. The client scrapes voetbalpoules.nl and returns the data as structs. These structs can be passed to another application or module, as in this case is done to the Telegram client. 

The voetbalpoules-client is centered around the `Client` struct. This struct is made up of Colly collector and several services for handling commmunication with different parts of voetbalpoules.nl.

### telegram

A Telegram voetbalpoules client can be found in `pkg/telegram`.
The Telegram client formats the voetbalpoules structs and sends the nicely formatted messages to a Telegram chat.

### cmd

The voetbalpoules-client and Telegram client are glued together in the `cmd` directory commands using the spf13 Cobra library.