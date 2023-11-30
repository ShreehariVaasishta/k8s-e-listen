# Kubernetes Event Listener in Go

This project provides a simple Kubernetes event listener written in Go. It watches for events in a specified Kubernetes namespace and can filter these events based on type, reason, and context.

## Prerequisites

- Go (1.13 or higher)
- Access to a Kubernetes cluster
- kubectl configured with access to your cluster

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/k8s-event-listener.git
cd k8s-event-listener
```

## Usage

Run the program using the following command:


```
go run main.go
```

You can specify a Kubernetes context (if different from your current context) by using the -context flag:

```
go run main.go -context your_context_name
```

## Configuration

- namespace: Set to default. Modify in the code to watch events in a different namespace.
- eventType: Set to Normal as the default event type. Change this to watch for different types of events.
- eventReason: Set to an empty string, meaning it will accept all event reasons. Specify a reason to filter events further.

## Features

- [x] Real-time monitoring of Kubernetes events.
- [x] Filters events by type, reason, and namespace.
- [x] Optional flag to specify a Kubernetes context.
- [x]  Basic error handling.

## Future Enhancements

- [ ] Enhanced error handling and logging.
- [ ] Configurable filters via command-line arguments or a configuration file.
- [ ] Support for multiple namespaces.
- [ ] Integration with notification systems (e.g., email, Slack).
- [ ] Deployment as a Kubernetes pod.
- [ ] User authentication and authorization for secure access.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

Support
If you have any issues or feature requests, please file an issue on the GitHub repository.

Acknowledgments
Thanks to the Kubernetes and Go communities for their excellent resources and documentation.
