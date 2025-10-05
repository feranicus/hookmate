Of course\! Here is a great README.md for your `hookmate` project. It's friendly, informative, and has plenty of emojis to make it welcoming.

Just create a new file named `README.md` in the root of your `hookmate` folder and paste the content below into it.

-----

# HookMate 🎣

[](https://www.google.com/search?q=https://goreportcard.com/report/github.com/feranicus/hookmate)
[](https://opensource.org/licenses/Apache-2.0)

**HookMate** is your new best friend for developer automation\! It's a lightweight, blazing-fast, and self-hosted API service written in Go. It acts as a central hub to receive webhooks and trigger actions across different services like Slack, Discord, GitHub, and more.

Think of it as your personal, code-driven "Zapier for Developers" that you control completely. No vendor lock-in, no complex UIs—just simple, powerful automation.

-----

## 🤔 What is HookMate?

In modern development, everything is event-driven. A `git push`, a new Docker image, a failed security scan—these are all events. HookMate listens for these events (via webhooks) and allows you to define simple actions in response.

It's designed to be:

  * **Minimalist:** No heavy frameworks or unnecessary features.
  * **Fast:** Built with Go for high performance and low resource usage.
  * **Developer-First:** Configuration is simple YAML, and the API is straightforward.
  * **Self-Hosted:** Run it anywhere—on your laptop, a server, or a Kubernetes cluster using Docker. You own your data and your workflows.

-----

## ✨ Core Features

  * **Webhook Gateway:** A single, stable entry point for all your incoming webhooks.
  * **Simple YAML Configuration:** Define your server settings and integration secrets in one easy-to-read file. ⚙️
  * **Blazing Fast:** Written in pure Go for maximum performance. ⚡️
  * **Container Ready:** Comes with a `Dockerfile` for easy deployment anywhere. 🐳
  * **Extensible:** Easily add new services (Discord, Telegram, etc.) to notify. 🔗
  * **Health Checks:** Includes a `/healthz` endpoint for monitoring. ❤️

-----

## 🚀 Quick Start

Get HookMate up and running in less than 5 minutes\!

### Prerequisites

Before you start, make sure you have the following installed:

  * [Go](https://go.dev/doc/install) (Version 1.22 or later)
  * A Slack Incoming Webhook URL. You can create one [here](https://api.slack.com/messaging/webhooks).

### 1\. Clone the Repository

Open your terminal and clone the project to your local machine.

```sh
git clone https://github.com/feranicus/hookmate.git
cd hookmate
```

### 2\. Configure Your Webhook

The application is configured using a `config.yaml` file.

1.  Create a copy of the example configuration:

    ```sh
    cp config.example.yaml config.yaml
    ```

    *(Note: You'll need to create `config.example.yaml` if it doesn't exist yet, or just create `config.yaml` directly)*

2.  Open `config.yaml` and paste your Slack Webhook URL into it:

    ```yaml
    server:
      port: "8080"

    integrations:
      slack:
        webhook_url: "YOUR_SUPER_SECRET_SLACK_WEBHOOK_URL_HERE"
    ```

### 3\. Run the Application

Now, let's bring the server to life\! The included `Makefile` makes this super easy.

```sh
# This will download dependencies and start the server
make run
```

You should see a log message confirming the server has started on port 8080.

### 4\. Send a Test Webhook\!

Success\! Your HookMate instance is running. Open a **new terminal** and send a test JSON payload using `curl`.

```sh
curl -X POST -H "Content-Type: application/json" \
-d '{"user":"feranicus", "message":"HookMate is alive! 🎉"}' \
http://localhost:8080/webhook/slack
```

Check your Slack channel... you should see a new message from HookMate\!

-----

## 🔌 API Endpoints

The MVP of HookMate is simple and focused.

  * `GET /healthz`
      * **Description:** A health check endpoint. Returns a `200 OK` with `{"status": "ok"}` if the service is running. Useful for load balancers and monitoring.
  * `POST /webhook/slack`
      * **Description:** The main endpoint for receiving a generic JSON payload to be forwarded to Slack.
      * **Body:** Any valid JSON object.
      * **Response:** Returns a `202 Accepted` immediately while it processes the notification in the background.

-----

## 🙌 Contributing

Contributions are welcome and appreciated\! Whether it's fixing a bug, adding a new integration, or improving the documentation, we'd love your help.

Please feel free to open an issue or submit a pull request.

-----

## 📜 License

This project is licensed under the **Apache 2.0 License**. See the `LICENSE` file for details.
