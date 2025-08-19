# Quill

<img src="quill-logo-cropped.png" alt="Quill Logo" width="200" />

Quill is a fast, modern static site generator and SPA blog engine written in Go. It features live reload, markdown support, and easy configuration for rapid blogging and publishing.

## Features

- **Markdown-based posts**: Write your content in markdown files.
- **Single Page Application**: Fast navigation and dynamic content loading.
- **Live reload**: Instant updates during development.
- **Configurable build and server**: Customize output and dev experience.
- **Easy initialization**: Quickly scaffold a new site with one command.
- **Syntax highlighting**: Prism.js integration for code blocks.
- **Responsive design**: Mobile-friendly out of the box.

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js (optional, for advanced JS/CSS customization)

### Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/bim-dev-tools/quill.git
cd quill
go build -o quill
```

### Initialize a New Site

Run the following command in your target directory:

```sh
./quill init
```

This will create:

- `index.html` (SPA entry point)
- `styles.css` (default styles)
- `.gitignore`
- `posts/0001_hello_world.md` (example post)

### Development Server

Start the live-reload development server:

```sh
./quill server
```

- Visit `http://localhost:8080` (default port) in your browser.
- Edit markdown files in `posts/` and see changes instantly.

### Build for Production

Generate the static site for deployment:

```sh
./quill build
```

Output will be in the `build/` directory (default).

## Configuration

Edit `.quill.config.yml` to customize:

```yaml
build_dir: build
html_entry_point: index.html
watch_files:
  - posts/*.md
  - styles.css
  - index.html
server:
  port: 8080
```

## Commands

- `init` — Scaffold a new site in the current directory.
- `server` — Start the live-reload development server.
- `build` — Build the static site for deployment.

## Contributing

Pull requests and issues are welcome! Please see the code in `cmd/`, `transpiler/`, `server/`, and `utils/` for entry points and architecture.

## License

MIT
