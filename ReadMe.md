# Solace PubSub+ CLI API Abstraction Tool

## Overview

This project provides a simple Command Line Interface (CLI) tool for interacting with Solace PubSub+ Event Broker's API. This CLI tool abstracts the underlying API calls, making it easier to manage Solace.

## Features

- List all VPNs or Queues
- Create a new VPN or Queue
- Delete a VPN or Queue

## Getting Started

Prerequisites
* Go (v1.15 or higher)
* Solace PubSub+ Event Broker instance

## Installation
Clone this repository:

```
git clone https://github.com/yourusername/solcli.git
```

Navigate to the project directory and build the project:

```
cd solcli
go build
```

## Usage
List VPNs

To list all VPNs:

```
solcli vpn list -u <username> -p <password> -U <url>
```

Create a VPN
To create a new VPN:

```
solcli vpn create [vpnName] -u <username> -p <password> -U <url>
```

## Code Structure
* solcli.go: Entry point for the CLI application.
* api/: Contains API handler for making API requests.
* models/: Contains structs that map to the API's JSON responses and requests.
* printer/: Utility to pretty-print the API responses.


## Contributing
Feel free to submit an PRs and discuss any enhancement requests.
