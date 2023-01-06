# Purpose
This is a small lab for demonstrating and understanding Kubernetes CPU requests and limits.

## Pre-Requisitites

### Kubernetes
- Make sure you have access to a Kubernetes cluster
-- Minikube 

### Tilt
- This lab uses Tilt so that we can move fast!
-- Install Tilt - https://docs.tilt.dev/install.html

## Whats in the Lab
- We're attempting to show how cGroups and Kubernetes Limits interact
- Two nearly identical Go apps
    - maxprocs
        - imports Ubers go.uber.org/automaxprocs
    - nomaxprocs
        - doens't import Ubers package
- Provides 3 deployments
    - A deployment of maxprocs with a limit set
        - This shows that the amount of usable CPUs matches the limits set for the container
    - A deployment of nomaxprocs without a limit set
        - This shows that there will be no limits to how much CPU the container will use
    - A deployment of nomaxprocs with a limit set
        - This will help in showing how CPU throttling occurs
- Deploys Prometheus so that you can see how the two different apps with limits behave
    - Use the query `container_cpu_cfs_throttled_periods_total{pod=~"cpu.+"}/container_cpu_cfs_periods_total{pod=~"cpu-.+"}` to see the different profiles 
- A very simple locust script for generating requests to test the different
    - Do a test run using localhost:8000 to test nomaxprocs without limits
    - Do a test run using localhost:8001 to test nomaxprocs with a 2 core limit
    - Do a test run using localhost:8002 to test maxprocs with a 2 core limit

## What Else?
- Play with the limits on the pods to see what the relative impact to CPU throttling is

