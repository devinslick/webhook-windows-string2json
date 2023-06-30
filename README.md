# webhook-windows-string2json
A quick tool to send a json payload to a webhook

# Webhook Tool

The **Webhook Tool** is a command-line application written in Go that sends a JSON payload as a POST request to a webhook destination URL. It accepts key-value pairs as input and converts them into JSON format.

## Usage

To use the **Webhook Tool**, follow the instructions below:

1. **Download** the compiled executable for your operating system from the [Releases](link-to-releases-page) page.

2. **Run** the tool from the command line with the following parameters:

    ```
    webhook.exe -url [webhook_url] -string2json [key-value pairs]
    ```

    Replace `[webhook_url]` with the destination URL of the webhook. Replace `[key-value pairs]` with an unlimited number of key-value pairs to be included in the JSON payload. Each key-value pair should be separated by a space.

    **Example:**

    ```
    webhook.exe -url https://example.com/webhookaddress -string2json key1 string1 key2 number2 key3 %RANDOMNUMBER%
    ```

3. **Provide** the necessary details for the parameters:

    - `-url` is the destination URL of the webhook.
    - `-string2json` is the key-value pairs to be converted into JSON format.

4. **Review** the output for success or error messages.

## Dependencies

The **Webhook Tool** has the following dependencies:

- Go programming language (version 1.16 or later)
- Standard Go libraries: `encoding/json`, `flag`, `fmt`, `log`, `net/http`, and `strings`

## Building from Source

If you prefer to build the **Webhook Tool** from source, follow the steps below:

1. **Clone** the repository:

    ```
    git clone [repository_url]
    ```

2. **Navigate** to the cloned repository:

    ```
    cd webhook-tool
    ```

3. **Build** the application:

    ```
    go build -ldflags "-H=windowsgui" -o webhook.exe webhook.go
    ```

4. **Follow** the usage instructions mentioned above to run the tool.

## Contributing

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, please submit a pull request.

## License

This project is licensed under the [GNU General Public License v3.0](link-to-license-file). Please see the LICENSE file for more details.
