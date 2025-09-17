## Project Overview

This project is a monolithic web application that combines a Go backend with a Vue.js frontend. The entire application is designed to be compiled into a single executable file.

**Backend:**
- **Language:** Go (version 1.23.4)
- **Functionality:** Serves a simple REST API at `/api/hello` and, more importantly, serves the compiled frontend assets.
- **Key Feature:** It uses Go's `embed` package to bundle the entire frontend (the `web/dist` directory) directly into the Go binary. This means no separate static file hosting is needed.

**Frontend:**
- **Framework:** Vue.js (version 3)
- **Build Tool:** Vite
- **Functionality:** A standard single-page application that acts as the user interface for the application.

**Architecture:**
The key architectural choice is the monolithic, single-binary deployment. The `build.sh` script first builds the Vue.js application, creating a set of static HTML, CSS, and JavaScript files in the `web/dist` directory. Then, it compiles the Go application. The Go application, in turn, embeds the contents of `web/dist` and serves them from memory. This results in a self-contained application that is easy to distribute and run.

## Building and Running

The primary method for building the application is the `build.sh` script.

**To build the application for Windows:**
```bash
./build.sh
```
This command will:
1.  Install npm dependencies for the frontend (if needed).
2.  Build the Vue.js application and place the output in `web/dist`.
3.  Build the Go application, embedding the frontend assets.
4.  Create a Windows executable named `video-collector-windows.exe` in the project root.

**To run the application:**
1.  After building, run the generated executable (e.g., `video-collector-windows.exe` on Windows).
2.  Open a web browser and navigate to `http://localhost:8080`.

**To run in development mode:**
- **Frontend:**
  ```bash
  cd web
  npm install
  npm run dev
  ```

项目中的Vue模块通常遵循以下结构:
```
frontend/src/views/{模块}/
├── index.vue              # 模块主入口
├── components/            # 页面专属组件目录
│   ├── componentA/             # 子组件
│   │   ├── index.vue      # 组件入口文件
│   │   └── components/    # 子组件页面专属组件目录
│   │       ├── a1/      # a1组件
│   │       │   └── index.vue
│   │       ├── a2/    # a2组件
│   │       │   └── index.vue
```

- **Backend:**
  ```bash
  go run main.go
  ```
Note: In development, the Go backend will not serve the frontend assets unless you have a `web/dist` directory from a previous build.

## Development Conventions

- The backend and frontend code are kept in separate directories (`server` and `web`, respectively, although `main.go` was moved to the root).
- The `build.sh` script is the single source of truth for building the combined application.
- The application is configured to be built for a Windows target, but this can be easily changed in the `build.sh` script by modifying the `GOOS` and `GOARCH` environment variables.
