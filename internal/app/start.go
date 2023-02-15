package app

import (
	"context"
	"sync"
)

var once sync.Once

func (a *statsApp) start(ctx context.Context) {
	withCancel, cancelFunc := context.WithCancel(ctx)
	a.log.Info("App starts collecting stats")
	streamChannel := a.containerStreamer.StartStreaming(withCancel, a.errCh)

	a.streamsPool = make([]portStream, 0)
	a.newPortStreamChannel = make(chan portStream)

	for {
		select {
		case stats := <-streamChannel:

			select {
			case stream := <-a.newPortStreamChannel:
				a.streamsPool = append(a.streamsPool, stream)

			default:
			}

			toDelete := []int{}

			for i, c := range a.streamsPool {
				select {
				case c.statsCh <- stats:
				case <-c.cancelCh:
					toDelete = append(toDelete, i)
				}
			}

			for _, i := range toDelete {
				close(a.streamsPool[i].statsCh)
				a.streamsPool[i] = a.streamsPool[len(a.streamsPool)-1]
				a.streamsPool = a.streamsPool[:len(a.streamsPool)-1]
			}

		case <-ctx.Done():
			cancelFunc()
			a.containerStreamer.Wait()
			a.log.Info("App ends collecting stats")
			return
		}
	}

}

func (a *statsApp) Start(ctx context.Context) {
	a.start(ctx)
}
