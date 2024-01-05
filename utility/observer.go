package utility

import (
	"DataHarbor/enums"
	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"sync"
)

type Observer func()

type ObserverObj struct {
	obs map[foundation.NotificationName]map[string]Observer
}

var observerObj = new(ObserverObj)

func AddAppearanceObserver(name string, observer Observer) {
	observerObj.AddAppearanceObserver(name, observer)
}

func (c *ObserverObj) AddAppearanceObserver(name string, observer Observer) {
	//dispatch.MainQueue().DispatchSync(observer)
	observer()
	c.fillChain(enums.AppearanceChangedNotification)[name] = observer
	sync.OnceFunc(func() {
		startObserver(enums.AppearanceChangedNotification)
	})()
}

func (c *ObserverObj) fillChain(types foundation.NotificationName) map[string]Observer {
	if c.obs == nil {
		c.obs = make(map[foundation.NotificationName]map[string]Observer)
	}

	chain, ok := c.obs[types]
	if !ok {
		chain = make(map[string]Observer)
		c.obs[types] = chain
	}
	return chain
}

func startObserver(types foundation.NotificationName) {
	target, selector := action.Wrap(func(objc.Object) {
		dispatch.MainQueue().DispatchAsync(func() {
			if chain, ok := observerObj.obs[types]; ok {
				for _, f := range chain {
					f()
				}
			}
		})
	})
	getDefaultNotificationCenter().AddObserverSelectorNameObject(
		target, selector, types, nil)
}

func getDefaultNotificationCenter() foundation.DistributedNotificationCenter {
	return objc.Call[foundation.DistributedNotificationCenter](
		foundation.DistributedNotificationCenterClass,
		objc.Sel("defaultCenter"))
}
