# HTMX - hx-swap-oob vs hx-on

## Pre-requesites

Before running the project, make sure you have [air](https://github.com/air-verse/air), [node](https://github.com/nvm-sh/nvm) and [templ](https://templ.guide/quick-start/installation) installed on your local machine.


## Set up project

### Clone repository

```bash
git clone https://github.com/webdevfuel/hx-swap-oob-vs-hx-on
```

### Install node packages

```bash
npm install
```

### Copy HTMX minified JS file from CDN to static directory

```bash
mkdir -p static
wget https://unpkg.com/htmx.org@2.0.1/dist/htmx.min.js -O ./static/htmx.min.js
```

### Copy Alpine.js minified JS file from CDN to static directory

```bash
wget https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js -O ./static/alpine.min.js
```

### Run the program

To run the program, use the `air` executable, which is configured in the `.air.toml` file to run tailwindcss, templ and the go server.

```bash
air
```
