# WebSocket Group Chat Application

This is a WebSocket-based group chat application built using **Golang (Gin framework)** and **Gorilla WebSocket** library. The application allows multiple users to join specific groups and communicate in real-time.

---

## Features

- Real-time messaging using WebSocket.
- Group-based communication.
- Broadcast messages to all members in a group.
- Easy-to-use UI with an input box and chat display.

---

## Prerequisites

1. **Go** (version 1.16 or higher) installed on your system.
2. **Node.js** and **npm** (if you want to customize the frontend).
3. **Web browser** to access the application.

---

## Installation and Setup

### Step 1: Clone the Repository

```bash
git clone <repository_url>
cd <repository_folder>
```

### Step 2: Install Dependencies

Install the required Go modules:

```bash
go mod tidy
```

Ensure the `github.com/gin-gonic/gin` and `github.com/gorilla/websocket` packages are installed.

---

### Step 3: Project Structure

Ensure the following files exist in your project directory:

```
.
├── main.go          # Backend code for WebSocket server
├──static
|      ├── index.html       # Frontend file for the chat interface
├──controlllers
|      ├── webSocket.go       # Frontend file for the chat interface
├── go.mod           # Go module file
├── go.sum           # Dependency file
```

---

### Step 4: Run the Application

1. Start the WebSocket server:

   ```bash
   go run main.go
   ```

2. Open your web browser and navigate to:

   ```
   http://localhost:8080/
   ```

---

## How to Use

1. Open the application in your browser.
2. Enter your **username (****`userID`****)** and **group name (****`groupID`****)**.
3. Click the **"Join Group"** button to connect to the chat server.
4. Open a new tab or window to add more users.
   - Use a different username but the same group name to chat in the same group.
   - Use a different group name to create and join a new group.
5. Type messages in the input box and click **"Send"** to broadcast messages.

---


## Customization

### Frontend Changes

Modify the `index.html` file to customize the UI.

### Backend Changes

- Update `main.go` to implement additional features like message history or user authentication.
- Use persistent storage to save messages if needed.

---

## Troubleshooting

1. **WebSocket Connection Fails:**

   - Ensure the server is running on `http://localhost:8080`.
   - Check the browser console for errors.

2. **Messages Not Broadcasting:**

   - Verify that all users are connected to the same group.
   - Check the server logs for errors.

---

