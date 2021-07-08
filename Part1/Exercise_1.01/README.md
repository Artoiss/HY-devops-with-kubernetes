## Commands to run

k3d cluster create -a 2

kubectl create deployment hashprogram --image=artoiss/hashprogram

## To see hash print

kubectl get pods

kubectl logs "pod_name"