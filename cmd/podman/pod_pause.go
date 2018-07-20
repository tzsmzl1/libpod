package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/projectatomic/libpod/cmd/podman/libpodruntime"
	"github.com/projectatomic/libpod/libpod"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	podPauseFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "pause all running pods",
		},
		LatestFlag,
	}
	podPauseDescription = `
   podman pod pause
   Pauses one or more pods.  The pod name or ID can be used.
`

	podPauseCommand = cli.Command{
		Name:                   "pause",
		Usage:                  "Pause one or more pods",
		Description:            podPauseDescription,
		Flags:                  podPauseFlags,
		Action:                 podPauseCmd,
		ArgsUsage:              "POD-NAME [POD-NAME ...]",
		UseShortOptionHandling: true,
	}
)

func podPauseCmd(c *cli.Context) error {
	if err := checkMutuallyExclusiveFlags(c); err != nil {
		return err
	}

	runtime, err := libpodruntime.GetRuntime(c)
	if err != nil {
		return errors.Wrapf(err, "error creating libpod runtime")
	}
	defer runtime.Shutdown(false)

	args := c.Args()
	var pods []*libpod.Pod
	var lastError error

	if c.Bool("all") {
		pods, err = runtime.Pods()
		if err != nil {
			return errors.Wrapf(err, "unable to get pods")
		}
	}
	if c.Bool("latest") {
		pod, err := runtime.GetLatestPod()
		if err != nil {
			return errors.Wrapf(err, "unable to get latest pod")
		}
		pods = append(pods, pod)
	}
	for _, i := range args {
		pod, err := runtime.LookupPod(i)
		if err != nil {
			if lastError != nil {
				logrus.Errorf("%q", lastError)
			}
			lastError = errors.Wrapf(err, "unable to find pod %s", i)
			continue
		}
		pods = append(pods, pod)
	}

	for _, pod := range pods {
		ctr_errs, err := pod.Pause()
		if err != nil {
			if lastError != nil {
				logrus.Errorf("%q", lastError)
			}
			lastError = errors.Wrapf(err, "unable to pause pod %q", pod.ID())
			continue
		} else if ctr_errs != nil {
			for ctr, err := range ctr_errs {
				if lastError != nil {
					logrus.Errorf("%q", lastError)
				}
				lastError = errors.Wrapf(err, "unable to pause container %q on pod %q", ctr, pod.ID())
			}
			continue
		}
		fmt.Println(pod.ID())
	}

	return lastError
}
