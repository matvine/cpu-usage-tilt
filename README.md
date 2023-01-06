# Purpose
This is a small lab for demonstrating and understanding Kubernetes CPU requests and limits.

## Pre-Requisitites

### Kubernetes
- Make sure you have access to a Kubernetes cluster
-- Minikube 

### Tilt
- This lab uses Tilt so that we can move fast!
-- Install Tilt - https://docs.tilt.dev/install.html

## Why is this Lab needed?
- How Kubernetes limits are implemented and enforced aren't widely understood
    - This can have massive impacts on performance of workloads
    - It can also mislead people about how their services are being implemented
- The Completely Fair Scheduler (CFS) does things that aren't completely intuitive to people used to looking at CPU usage graphs
    - Important reading: 
        -https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-cpu#doc-wrapper

## A Little More Info
- Most people *do* understand that the CPU requests a pod makes represent a guaranteed amount of resources for a Kubernetes pod
- Most people *do* understand that the CPU limits represent the maximum amount of CPU that a Kubernetes pod can burst to
- Most people *do not* understand that the CPU limits are a quota calculated on a set interval (typically 100ms) and measured across all cores
- In a typical, default setting, a CPU limit of "2" means that for any given `cfs_period`, the container can use 200ms of cpu time.... across all available cores!
- Once the container has reached that, it will be throttled for the remainder of the `cfs_period`
- Go, by default, will try to use as many cores as it can discover via `runtime.NumCPU()`
    - `runtime.NumCPU()` will return the amount of cores available to the underlying node - which is ofter a few more than the limits we set
    - As we noted, the 200ms quoata we get setting a CPU limit of "2" will be measured across all cores
    - Workloads that spawn lots of Goroutines / threads can end up spreading work across all of these cores, resulting in increased rates of throttling

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
- Deploys cAdvisor to 
- Deploys Prometheus so that you can see how the two different apps with limits behave
    - Use the query `container_cpu_cfs_throttled_periods_total{pod=~"cpu.+"}/container_cpu_cfs_periods_total{pod=~"cpu-.+"}` to see the different profiles 
- A very simple locust script for generating requests to test the different
    - Do a test run using localhost:8000 to test nomaxprocs without limits
    - Do a test run using localhost:8001 to test nomaxprocs with a 2 core limit
    - Do a test run using localhost:8002 to test maxprocs with a 2 core limit

## What You'll See
- You'll see how 

## What Else?
- Play with the limits on the pods to see what the relative impact to CPU throttling is

