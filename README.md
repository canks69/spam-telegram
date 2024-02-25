# Telegram Spam Bot

Telegram Spam Bot is a simple application built using Go to send spam messages to Telegram users.

## Features
- Send spam messages to specified Telegram users.
- Set a time interval between message deliveries.
- Infinity mode to send spam messages endlessly.

## Usage
1. Make sure you have Go installed on your computer. If not, you can download it [here](https://golang.org/dl/).
2. Clone this repository to your computer.
3. Install the dependencies using the command:
    ```bash
   go mod download
    ```
4. Rename `.env-example` to `.env` and replace `<bot_token>` with the token you copied, and `<chat_id>` with the chat ID where you want to send the spam messages.
    ```
   TELEGRAM_BOT_TOKEN=<bot_token>
   TELEGRAM_CHAT_ID=<chat_id>
    ```
5. Run the application using the command:
    ```bash
   go run main.go
    ```
6. The bot is now active and ready to send spam messages to the specified users.

## Contributing
You can contribute to this project by opening pull requests or filing new issues.

## License
This project is licensed under the [MIT License](LICENSE).
