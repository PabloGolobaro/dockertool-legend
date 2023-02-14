package containerStreamer

import (
	"context"
	"github.com/docker/docker/api/types"
	"time"
)

func (c *ContainerStreamer) startNewContainersController(errCh chan error) (chan containerStatsChannel, chan string) {
	newContainers := make(chan containerStatsChannel)
	delContainers := make(chan string)

	c.log.Debug("Start NewContainersController")

	c.Add(1)

	go func(errCh chan error) {
		ctx, cancel := context.WithCancel(context.Background())

		defer c.Done()
		defer close(newContainers)
		defer cancel()

		for {
			select {
			case id, ok := <-delContainers:
				if !ok {
					c.log.Debug("Stop NewContainersController")
					return
				}
				c.log.Debugw("Delete Container from Controller", "ID", id)

				delete(c.containerMap, id)
			default:
				containerList, err := c.cli.ContainerList(ctx, types.ContainerListOptions{})
				if err != nil {

					c.log.Errorw("Stop NewContainersController", "error", err)

					errCh <- err

					return
				}
				for _, container := range containerList {
					_, ok := c.containerMap[container.ID]
					if !ok {

						containerChan := c.streamContainer(ctx, container)

						newContainers <- containerStatsChannel{statsCh: containerChan, ID: container.ID}

						c.containerMap[container.ID] = struct{}{}
					}
				}
				time.Sleep(time.Second)
			}
		}

	}(errCh)

	return newContainers, delContainers
}
