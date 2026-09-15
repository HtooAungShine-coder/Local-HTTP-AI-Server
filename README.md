# Local-HTTP-AI-Server
Local HTTP based AI with OpenRouter in Golang programming language

# Local HTTP AI Server in Golang (Supreme Luna AI)

A sleek, lightweight, high-performance local HTTP chat server built from scratch using **Go (Golang)** [sports]. The application handles concurrent frontend form payloads, executes structured API calls to cloud LLMs via **OpenRouter**, and renders streaming text layouts back to a minimalist web interface.

## 🎮 Interface Aesthetic
* **Theme:** High-visibility terminal workspace (Solid Pitch Black & Orange minimal accent colorway)
* **Fonts:** Professional system monospace configurations optimized for long-term eye comfort.

## 🚀 Features
* Full custom standard library routing (`net/http`) without bulky external frameworks.
* Native JSON payload handling (`encoding/json`) optimized for OpenRouter Response mappings.
* Universal model protection using the `openrouter/free` model router to guarantee zero downtime from model deprecation.
* Absolute gateway proxy routing (`window.location.origin`) to prevent browser cross-origin policy (CORS) blocks.

## 📂 Project Structure
```text
├── API/
│   └── api.go       # Core structural definitions, headers, and HTTP POST Client requests
├── index.html       # Clean minimalist dark/orange UI console dashboard & response parser
├── main.go          # Route handlers, server loop execution, and terminal event log tracing
└── README.md        # Documentation layout
```

## 🛠️ Installation & Setup

### 1. Clone the repository
```bash
git clone https://github.com
cd local-http-ai-server-golang
```

### 2. Configure Your Credentials
To keep your key safe from public view, make sure your code uses `os.Getenv("OPENROUTER_API_KEY")` inside `API/api.go`.

Choose your preferred way to feed your API key into the app:

#### Option A: Running with Environment Variables (Recommended & Safe 🔒)
Do not change anything in your files. Instead, set the variable directly in your terminal console before starting the server. Run the command matching your operating system:

* **Windows PowerShell:**
  ```powershell
  $env:OPENROUTER_API_KEY="sk-or-v1-your-key-here"; go run .
  ```
* **Windows Command Prompt (CMD):**
  ```cmd
  set OPENROUTER_API_KEY=sk-or-v1-your-key-here&& go run .
  ```
* **Mac OS / Linux Terminal:**
  ```bash
  export OPENROUTER_API_KEY="sk-or-v1-your-key-here" && go run .
  ```

#### Option B: Hardcoding (For Quick Tests Only ⚠️)
Open `API/api.go` and replace the token assignment variable with your raw text token key inside the quotes. **Make sure to delete this before pushing to GitHub!**
```go
apiKey := "sk-or-v1-your-key-here"
```

### 3. Deploy inside Browser Workspace
Open your web browser tool layout and navigate to the localhost engine pipeline address:
```text
http://localhost:8080
```

## ⚙️ Backend Log Blueprint
Your server traces state actions directly to the hosting shell environment console for foolproof tracking diagnostics:
* `[FRONTEND] User input received: <message>`
* `[BACKEND SUCCESS] Response successfully fetched from OpenRouter.`
* `[BACKEND EXECUTOR FAILURE]: <error_trace_block>`

## 📜 License
MIT License. Feel free to tweak, fork, and upgrade!
