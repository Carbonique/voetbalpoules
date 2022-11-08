# About

Life is busy enough already, so why check [voetbalpoules.nl](https://voetbalpoules.nl) daily for seeing how you and your friends are doing in predicting football matches?

This application helps you by sending your and your friends predictions to a Telegram chat. 

Sit back and enjoy the show!

## Supported data

1. [The current ranking for a pool](assets/telegram-stand.png)
2. [Predictions for upcoming matches played today or tomorrow](assets/telegram-voorspellingen.png)
3. Predictions and results for matches played today or yesterday

## Usage

1. Create an `.env` file
2. 

## Technical Architecture

## voetbalpoules-client

A voetbalpoules-client can be found in `pkg/voetbalpoules-client`. The client scrapes voetbalpoules.nl and returns the data as structs. 

## telegram

A Telegram voetbalpoules client can be found in `pkg/telegram`.
The Telegram client formats the voetbalpoules structs and sends the nicely formatted data to a Telegram chat.

## cmd

The voetbalpoules-client and Telegram client are glued together by the spf13 Cobra library.