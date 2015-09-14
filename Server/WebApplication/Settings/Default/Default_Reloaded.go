package Default

import (
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Settings"
)

func (d *defaultSettings) ConfigReloaded() {
	for _, h := range d.subscribersOnReload {
		h.OnConfigReloaded()
	}
}

func (d *defaultSettings) SubscribeConfigReloaded(handler ConfigReloadHandler) {
	d.reloadEventNotifyLock.Lock()
	defer d.reloadEventNotifyLock.Unlock()
	d.subscribersOnReload = append(d.subscribersOnReload, handler)
}

func (d *defaultSettings) UnsubscribeConfigReloaded(handler ConfigReloadHandler) {
	d.reloadEventNotifyLock.Lock()
	defer d.reloadEventNotifyLock.Unlock()

	newSubscriberList := []ConfigReloadHandler{}
	for ind, _ := range d.subscribersOnReload {
		if d.subscribersOnReload[ind] != handler {
			newSubscriberList = append(newSubscriberList, handler)
		}
	}
	d.subscribersOnReload = newSubscriberList
}
