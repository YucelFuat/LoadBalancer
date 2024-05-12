# TCP Load Balancer

This is a simple TCP load balancer written in Go. It listens for incoming connections on a specified port and forwards them to one of the backend servers in a round-robin fashion.

## Usage

1. Clone the repository:

git clone <repository_url>

2. Navigate to the project directory:

cd tcp-load-balancer

3. Build the project:

go build

4. Run the executable:


5. The load balancer will start listening for incoming connections on `localhost:8080` by default. You can modify the `listenAddr` variable in the code to change the listening address.

6. You can create demo servers in your Linux Machine via openning different terminals and running these commands in each

- npx http-server -p 5001
- npx http-server -p 5002
- npx http-server -p 5003

8. Backend servers are specified in the `servers` slice in the code. You can add or remove backend servers as needed.

## Configuration

You can configure the load balancer by modifying the following variables in the code:

- `listenAddr`: The address and port on which the load balancer listens for incoming connections.
- `servers`: The list of backend servers to which the load balancer forwards incoming connections.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

