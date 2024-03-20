package utility

import (
	"github.com/charliego3/assistant/enums"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type Observer func()

type ObserverObj struct {
	obs map[foundation.NotificationName][]Observer
}

var observerObj = new(ObserverObj)

func AddAppearanceObserver(observer Observer) {
	observerObj.AddAppearanceObserver(observer)
}

func (c *ObserverObj) AddAppearanceObserver(observer Observer) {
	observer()
	if c.fill(enums.AppearanceChangedNotification, observer) {
		c.Start(enums.AppearanceChangedNotification)
	}
}

func (c *ObserverObj) fill(types foundation.NotificationName, observer Observer) bool {
	if c.obs == nil {
		c.obs = make(map[foundation.NotificationName][]Observer)
	}

	chain, ok := c.obs[types]
	c.obs[types] = append(chain, observer)
	return !ok
}

func (c *ObserverObj) Start(types foundation.NotificationName) {
	target, selector := action.Wrap(func(objc.Object) {
		dispatch.MainQueue().DispatchAsync(func() {
			if chain, ok := c.obs[types]; ok {
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
