package harvester

import (
	"context"

	"github.com/coreos/go-systemd/v22/dbus"
	godbusv5 "github.com/godbus/dbus/v5"
	"github.com/sirupsen/logrus"
)

func RestartService(unit string) error {
	ctx := context.Background()
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		logrus.Errorf("Failed to create new connection for systemd. err: %v", err)
		return err
	}
	responseChan := make(chan string, 1)
	if _, err := conn.RestartUnitContext(ctx, unit, "fail", responseChan); err != nil {
		logrus.Errorf("Failed to restart service %s. err: %v", unit, err)
		return err
	}
	return nil
}

func WatchDBusSignal(ctx context.Context, iface, objPath string, handlerFunc func(s *godbusv5.Signal)) {
	conn, err := generateDBUSConnection()
	if err != nil {
		return
	}

	matchInterFace := godbusv5.WithMatchInterface(iface)
	matchObjPath := godbusv5.WithMatchObjectPath(godbusv5.ObjectPath(objPath))
	err = conn.AddMatchSignalContext(ctx, matchObjPath, matchInterFace)
	if err != nil {
		panic(err)
	}

	signals := make(chan *godbusv5.Signal, 2)
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

func generateDBUSConnection() (*godbusv5.Conn, error) {
	conn, err := godbusv5.SystemBus()
	if err != nil {
		logrus.Warnf("Init DBus connection failed. err: %v", err)
		return nil, err
	}

	return conn, nil
}
