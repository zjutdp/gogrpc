Followed https://1byte.io/developer-guide-to-docker-and-kubernetes for quickly going over Docker and Kubernetes basics. See original.md if the link is gone.

## Notes
* To enable minikube working, you need use version: 0.25.2, otherwise it will hang at "start cluster components" in Dell office network with Dell Mac OS image versino 10.11
* To enable Kubernetes on docker-for-desktop you need get VPN connected for crossing the GFW at home network environment, so that all of the k8s images are downloadable through, no problem found in Dell office network. Besides, the "Reset to factory defaults" is a powerful tool for starting over (Recommended approach)
* You need clear all the ENV vars with leading DOCKER in current shell, which was set by minikube previously
* You could download minikube-v1.0.6.iso directly to ~/.minikube/cache/iso using ignore certificate options. For example: curl -k -O https://storage.googleapis.com/minikube/iso/minikube-v1.0.6.iso, and start minikube again
* Finally, run "kubectl delete deploy k8s-demo-deployment" to delete the deployment and clean up the K8S env
