package harvester

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/sirupsen/logrus"
)

func WatchDBusSignal(ctx context.Context, iface, objPath string, handlerFunc func(s *dbus.Signal)) {
	conn, err := generateDBUSConnection()
	if err != nil {
		return
	}

	matchInterFace := dbus.WithMatchInterface(iface)
	matchObjPath := dbus.WithMatchObjectPath(dbus.ObjectPath(objPath))
	err = conn.AddMatchSignalContext(ctx, matchObjPath, matchInterFace)
	if err != nil {
		panic(err)
	}

	signals := make(chan *dbus.Signal, 2)
	conn.Signal(signals)

	logrus.Infof("Watch DBus signal...")
	for {
		select {
		case signalContent := <-signals:
			logrus.Debugf("Got signal: %+v", signalContent)
			handlerFunc(signalContent)
		case <-ctx.Done():
			return
		}
	}
}

func generateDBUSConnection() (*dbus.Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		logrus.Warnf("Init DBus connection failed. err: %v", err)
		return nil, err
	}

	return conn, nil
}
