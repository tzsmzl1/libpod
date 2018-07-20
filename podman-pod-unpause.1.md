% podman-pod-unpause "1"

## NAME
podman\-pod\-unpause - Unpause one or more pods

## SYNOPSIS
**podman pod unpause** [*options*] *pod* ...

## DESCRIPTION
Unpauses all the processes in the containers associated with a pod.  You may use pod IDs or names as input.

## OPTIONS

**--all, -a**

Unpauses all pods

**--latest, -l**

Instead of providing the pod name or ID, unpause the last created pod.

## EXAMPLE

podman pod unpause mywebserver

podman pod unpause 860a4b23

podman pod unpause -l

podman pod unpause --all

## SEE ALSO
podman-pod(1), podman-pod-unpause(1)

## HISTORY
July 2018, Originally compiled by Peter Hunt <pehunt@redhat.com>
