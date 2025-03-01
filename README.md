# RDS Proxy Agent

A lightweight **Go-based TCP proxy** that securely forwards database connections from a **public IP** to an **PostgreSQL instance**. Similar to Google's Cloud SQL Auth Proxy, it enables access control via firewall rules while keeping credentials secure.

## Quick Started
Check for the latest version on the https://cloud-sql-proxy.vercel.app and use the following instructions for your OS and CPU architecture.

<!-- {x-release-please-start-version} -->
<details open>
<summary>Linux amd64</summary>

```sh
# see Releases for other versions
URL="https://cloud-sql-proxy.vercel.app"

curl "$URL/db-proxy-linux-amd64.zip" -o db-proxy-linux-amd64.zip

unzip db-proxy-linux-amd64.zip

chmod +x db-proxy-linux-amd64

export PUBLIC_IP="0.0.0.0:PORT"
export DB_HOST="your-database-endpoint:PORT"

./db-proxy-linux-amd64
```
</details>

<details open>
<summary>Linux arm64</summary>

```sh
# see Releases for other versions
URL="https://cloud-sql-proxy.vercel.app"

curl "$URL/db-proxy-linux-arm64.zip" -o db-proxy-linux-arm64.zip

unzip db-proxy-linux-arm64.zip

chmod +x db-proxy-linux-arm64

export PUBLIC_IP="0.0.0.0:PORT"
export DB_HOST="your-database-endpoint:PORT"

./db-proxy-linux-arm64
```
</details>

<details open>
<summary>Mac (Intel)</summary>

```sh
# see Releases for other versions
URL="https://cloud-sql-proxy.vercel.app"

curl "$URL/db-proxy-darwin-amd64.zip" -o db-proxy-darwin-amd64.zip

unzip db-proxy-darwin-amd64.zip

chmod +x db-proxy-darwin-amd64

export PUBLIC_IP="0.0.0.0:PORT"
export DB_HOST="your-database-endpoint:PORT"

./db-proxy-darwin-amd64
```
</details>

<details>
<summary>Mac (Apple Silicon)</summary>

```sh
# see Releases for other versions
URL="https://cloud-sql-proxy.vercel.app"

curl "$URL/db-proxy-darwin-arm64.zip" -o db-proxy-darwin-arm64.zip

unzip db-proxy-darwin-arm64.zip

chmod +x db-proxy-darwin-arm64

export PUBLIC_IP="0.0.0.0:PORT"
export DB_HOST="your-database-endpoint:PORT"

./db-proxy-darwin-arm64
```
</details>

<details>
<summary>Windows x64</summary>

```sh
# see Releases for other versions
curl https://cloud-sql-proxy.vercel.app/db-proxy-windows-amd64.zip

Unzip and get started

```

</details>


## Features
- **Secure TCP proxy** for database PostgreSQL
- **Public IP forwarding** to an database instance
- **Lightweight & efficient** using Go's concurrency
- **Systemd service support** for easy deployment

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/technomersio/sql-proxy-agent.git
   cd sql-proxy-agent
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the binary:
   ```sh
   go build -o db-proxy ./cmd/main.go
   ```

4. Run the proxy:
   ```sh
   export PUBLIC_IP="0.0.0.0:5432"
   export DB_HOST="your-database-endpoint:5432"
   ./rds-proxy
   ```

***Note:** Your free to change ports lets say your using mysql then port will be 3306

## Deploy as a Systemd Service
1. Copy the service file:
   ```sh
   sudo cp systemd/db-proxy.service /etc/systemd/system/
   ```

***Note:** Update port respectively in systemd/db-proxy.service file

2. Enable and start the service:
   ```sh
   sudo systemctl enable db-proxy
   sudo systemctl start db-proxy
   ```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing
Feel free to open issues or submit pull requests! 🚀