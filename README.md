# 🐝 HiveBedrock AutoConnector

![Go](https://img.shields.io/badge/go-1.20+-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg)

A lightweight, configurable CLI tool written in Go that automatically connects to the Hive Bedrock Network, stays online for a configurable duration, disconnects cleanly, and schedules the next connection with a randomized delay. Designed for reliability and ease of use in open-source and self-hosted scenarios.

---

## Table of Contents

- [Features](#features)  
- [Prerequisites](#prerequisites)  
- [Installation](#installation)  
- [Usage](#usage)  
  - [Command-Line Flags](#command-line-flags)  
  - [Interactive Mode](#interactive-mode)  
- [Examples](#examples)  
- [Configuration](#configuration)  
- [Contributing](#contributing)  
- [License](#license)  

---

## Features

- 🔒 **Device-Flow Authentication**  
  Authenticate via Microsoft’s OAuth Device Flow in your browser.

- 🌐 **Multi-Region Support**  
  Choose between NA, EU, or Asia Hive Bedrock endpoints.

- ⏱️ **Customizable Durations**  
  Specify how long to stay connected on each cycle.

- 🎲 **Randomized Delay**  
  Wait 23 hours + 0–60 minutes (or your custom interval) before reconnecting.

- 📊 **Clean Logging**  
  Colorized console output for INFO, SUCCESS, ERROR, DISCONNECT, and NEXT schedules.

---

## Prerequisites

- Go **1.20** or newer  
- A valid Microsoft account to authenticate on the Hive network  
- Internet access to connect to `hivebedrock.network:19132`

---

## Installation

```bash
# Clone the repository
https://github.com/Daniel-Ric/Hive-Bedrock-AutoLogin.git
cd Hive-Bedrock-AutoLogin

# Build the binary
go build -o hb-connector ./cmd/connector
