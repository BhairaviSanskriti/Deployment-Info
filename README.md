# Deployment-Info
It is a command-line porgram written in Go that retrieves deployment data from a Kubernetes cluster. It displays the deployments in a particular namespace along with the number of healthy and unhealthy pods for each deployment in a JSON format.

## Requirements
Deployment Data requires the following software to be installed:
- Go version 1.15 or later
- The cobra and kubernetes packages for Go

# Installation
1. Clone the repository:
 ```
git clone https://github.com/BhairaviSanskriti/Deployment-Info.git
```
2. Navigate to the cloned repository:
```
cd Deployment-Info/ 
```
3. Build the binary
```
go build -o k8s-info
```

## Usage

To get the configuration of the kubernetes cluster, run the `kubectl config view` commamd. Store the output in a file named `kubeconfig`. You can name it whatever you want.
```
kubectl config view > kubeconfig
```

To run Deployment Data, execute the following command:
```
./k8s-info --namespace <namespace> --kubeconfig <path/to/kubeconfig>
```
or
```
./k8s-info -n <namespace> -k <path/to/kubeconfig>
```

The program will display the deployment data in a JSON format, including the deployment name, the number of healthy replicas, and the number of unhealthy replicas.

Please make sure to replace `<path/to/kubeconfig>` with the actual path of your configuration file.
The `--namespace` flag specifies the namespace to retrieve deployments from, and `--kubeconfig` flag specifies the absolute path to the kubeconfig file.

By default, Deployment Info will use the `default` namespace.

If the `kubeconfig` file is not specified, the tool will check for the file in the following locations, in order:

1. The path specified by the `$KUBECONFIG` environment variable
2. The `$HOME/.kube/config` file
3. The kubeconfig file in the current working directory
    - If you have saved your `kubeconfig` file in the current working directory, please name the file `kubeconfig`. This will allow the repository to work properly without any additional configuration.
    - If you have named your configuration file something else like `my-config-file`, you will need to use the `--kubeconfig ./my-config-file` or `-k ./my-config-file` flag when running Kubernetes commands.
If the kubeconfig file is not found in any of these locations, the tool will attempt to use the in-cluster configuration.


## Example

To display data for the deployments in the namespace `kube-system`  with the cluster configuration stored in the `~/.kube/config` file, run the below command:

```
./k8s-info --namespace kube-system --kubeconfig ~/.kube/config
```
![image](https://user-images.githubusercontent.com/106534693/221357458-f0c420d1-e717-4fcc-b5d7-6463907f99b2.png)

You can use `jq` for parsing, manipulating and formatting JSON data. Make sure to have `jq` installed already.
```
./k8s-info --namespace kube-system --kubeconfig ~/.kube/config | jq
```
![image](https://user-images.githubusercontent.com/106534693/221357492-e162528d-464e-483e-a5c1-90feb9270017.png)
## Conclusion
This code provides a simple and effective way to retrieve deployment data from a Kubernetes cluster. By using this tool, you can quickly and easily monitor the health of your deployments and troubleshoot any issues that arise.
