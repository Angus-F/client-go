/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package watch

import (
	"sync"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"github.com/Angus-F/client-go/tools/cache"
)

func newEventProcessor(out chan<- watch.Event) *eventProcessor {
	return &eventProcessor{
		out:  out,
		cond: sync.NewCond(&sync.Mutex{}),
		done: make(chan struct{}),
	}
}

// eventProcessor buffers events and writes them to an out chan when a reader
// is waiting. Because of the requirement to buffer events, it synchronizes
// input with a condition, and synchronizes output with a channels. It needs to
// be able to yield while both waiting on an input condition and while blocked
// on writing to the output channel.
type eventProcessor struct {
	out chan<- watch.Event

	cond *sync.Cond
	buff []watch.Event

	done chan struct{}
}

func (e *eventProcessor) run() {
	for {
		batch := e.takeBatch()
		e.writeBatch(batch)
		if e.stopped() {
			return
		}
	}
}

func (e *eventProcessor) takeBatch() []watch.Event {
	e.cond.L.Lock()
	defer e.cond.L.Unlock()

	for len(e.buff) == 0 && !e.stopped() {
		e.cond.Wait()
	}

	batch := e.buff
	e.buff = nil
	return batch
}

func (e *eventProcessor) writeBatch(events []watch.Event) {
	for _, event := range events {
		select {
		case e.out <- event:
		case <-e.done:
			return
		}
	}
}

func (e *eventProcessor) push(event watch.Event) {
	e.cond.L.Lock()
	defer e.cond.L.Unlock()
	defer e.cond.Signal()
	e.buff = append(e.buff, event)
}

func (e *eventProcessor) stopped() bool {
	select {
	case <-e.done:
		return true
	default:
		return false
	}
}

func (e *eventProcessor) stop() {
	close(e.done)
	e.cond.Signal()
}

// NewIndexerInformerWatcher will create an IndexerInformer and wrap it into watch.Interface
// so you can use it anywhere where you'd have used a regular Watcher returned from Watch method.
// it also returns a channel you can use to wait for the informers to fully shutdown.
func NewIndexerInformerWatcher(lw cache.ListerWatcher, objType runtime.Object) (cache.Indexer, cache.Controller, watch.Interface, <-chan struct{}) {
	ch := make(chan watch.Event)
	w := watch.NewProxyWatcher(ch)
	e := newEventProcessor(ch)

	indexer, informer := cache.NewIndexerInformer(lw, objType, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			e.push(watch.Event{
				Type:   watch.Added,
				Object: obj.(runtime.Object),
			})
		},
		UpdateFunc: func(old, new interface{}) {
			e.push(watch.Event{
				Type:   watch.Modified,
				Object: new.(runtime.Object),
			})
		},
		DeleteFunc: func(obj interface{}) {
			staleObj, stale := obj.(cache.DeletedFinalStateUnknown)
			if stale {
				// We have no means of passing the additional information down using
				// watch API based on watch.Event but the caller can filter such
				// objects by checking if metadata.deletionTimestamp is set
				obj = staleObj.Obj
			}

			e.push(watch.Event{
				Type:   watch.Deleted,
				Object: obj.(runtime.Object),
			})
		},
	}, cache.Indexers{})

	go e.run()

	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		defer e.stop()
		informer.Run(w.StopChan())
	}()

	return indexer, informer, w, doneCh
}
