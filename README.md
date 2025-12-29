# 2048 Game Project

Implementation of the classic 2048 game using **Go (WebAssembly)** for core game logic and **React** for the UI.

## Project Structure

- `gameplay/`: Contains the core game logic written in Go.
  - `logic.go`: Implements the game rules, board updates, and tile merging logic.
  - `model.go`: Defines the data structures (`GameBoard`, `Direction`) and helper methods.
  - `main_wasm.go`: The entry point for compiling to WebAssembly. Exposes functions (`newGame`, `move`, `getState`) to the JavaScript environment.
  - `main_cli.go`: Initial CLI implementation (if available) or basic Go entry point.
- `ui/`: Contains the React frontend application.
  - `src/`: React source code (Components, styles).
  - `public/`: Static assets (`index.html`, `wasm_exec.js`, etc.).
  - `rsbuild.config.js`: Configuration for the Rsbuild bundler.
  - `package.json`: NPM dependencies and scripts.

## Tech Stack

- **Core Logic**: Go (Golang) compiled to WebAssembly (WASM).
- **Frontend**: React 19.
- **Bundler**: Rsbuild (Rust-based web bundler).
- **Styling**: CSS / Tailwind (if applicable from context, but checks show standard css usage currently).
- **Linting/Formatting**: Biome.

## Development & Build

### Prerequisites

- Go 1.21+
- Node.js & Yarn (or npm)

### 1. Build the Go WASM

You need to compile the Go code into a `.wasm` file that the browser can load.

```bash
cd gameplay
GOOS=js GOARCH=wasm go build -o ../ui/public/main.wasm .
```

> **Note**: This command compiles the Go code targeting the `js` OS and `wasm` architecture and outputs the `main.wasm` file directly into the `ui/public` directory so it can be served by the dev server.

### 2. Run the Frontend

Navigate to the `ui` directory to install dependencies and start the development server.

```bash
cd ui
yarn install
yarn dev
```

The application will be available at `http://localhost:3000`.

### 3. Deploy (Production Build)

To create a production-ready build of the frontend:

```bash
cd ui
yarn build
```

The output will be in `ui/dist`. You can deploy this directory to any static file hosting service (Vercel, Netlify, GitHub Pages, S3, etc.). Ensure `main.wasm` is correctly served with the MIME type `application/wasm`.

## How it Works

1.  **Initialization**: When the React app loads, it fetches the `main.wasm` file and instantiates the Go WebAssembly runtime using `wasm_exec.js`.
2.  **Interaction**: User inputs (arrow keys, swipes) are captured by React.
3.  **Bridge**: React calls the Go functions (e.g., `global.move('up')`) exposed via the `syscall/js` package in `main_wasm.go`.
4.  **State Update**: The Go logic calculates the new board state and returns it as a JSON string.
5.  **Rendering**: React parses the new state and updates the UI accordingly.
