# Kubernetes Event Listener

This Go program is designed to monitor events in a Kubernetes cluster. It allows users to specify filters such as event type, reason, and namespace to target specific events. This tool is particularly useful for developers and system administrators who need to keep an eye on specific occurrences within their Kubernetes environment.

## Features

### Existing Features
- **Event Filtering:** Users can filter events based on their type (e.g., `Normal`, `Warning`), reason, and namespace. This allows for focused monitoring of events that are of specific interest or importance.
- **Custom Kubernetes Context:** The program can use a specified kubeconfig file, allowing it to be used in different Kubernetes environments or contexts.
- **Simple CLI Usage:** Easily run the program with command-line arguments to set your desired filters.

### Future Features/Enhancements
- **Enhanced Filtering Options:** Future versions could include more granular filters, such as filtering by specific resource types or labels.
- **Real-Time Alerts:** Integration with notification systems to send real-time alerts based on certain event triggers.
- **Web Dashboard:** A web-based UI to view and filter events in real-time, enhancing the usability for those who prefer graphical interfaces.
- **Persistent Logging:** Option to log events to a file or a database for historical analysis and auditing purposes.
- **Cluster-Wide Monitoring:** Extend functionality to support monitoring events across multiple clusters simultaneously.

## Getting Started

### Prerequisites
- Go (version 1.x or later)
- Access to a Kubernetes cluster with `kubectl` configured

### Installation
1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Compile the program (optional):
   ```bash
   go build
   ```

Run the program with the desired flags.

## Usage
Run the program using the following command format:

```
go run main.go [--kubeconfig=path/to/kubeconfig] [--namespace=namespace] [--event-type=event_type] [--reason=reason]
```

### Flags
- **--kubeconfig:** Path to the kubeconfig file (optional)
- **--namespace:** Kubernetes namespace to monitor (default: default)
- **--event-type:** Type of event to filter (default: Normal)
- **--reason:** Specific reason to filter events (optional)

Example:

```bash
go run main.go --namespace=my-namespace --event-type=Warning --reason=PodCrashLoopBackOff
```

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your proposed changes or enhancements.

## License

This project is licensed under the MIT License.


