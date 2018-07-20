% podman-pod-pause "1"

## NAME
podman\-pod\-pause - Pause one or more pods

## SYNOPSIS
**podman pod pause** [*options*] *pod* ...

## DESCRIPTION
Pauses all the processes in the containers associated with a pod.  You may use pod IDs or names as input.

## OPTIONS

**--all, -a**

Pauses all pods

**--latest, -l**

Instead of providing the pod name or ID, pause the last created pod.

## EXAMPLE

podman pod pause mywebserver

podman pod pause 860a4b23

podman pod pause -l

podman pod pause --all

## SEE ALSO
podman-pod(1), podman-pod-unpause(1)

## HISTORY
July 2018, Originally compiled by Peter Hunt <pehunt@redhat.com>
