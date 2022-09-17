package main

import (
	"github.com/kardianos/service"
	"log"
	"flag"
)

type Service struct {}

var logger service.Logger

func (*Service) Start(_ service.Service) error {
	if err := StartProcessAsCurrentUser("d:\\temp\\sysproxy.exe", "d:\\temp\\sysproxy.exe global 127.0.0.1:13004 localhost;127.*;10.*;192.168.*", "d:\\temp", false); err != nil {
		return err
	}

	return nil
}

func (*Service) Stop(_ service.Service) error {
	if err := StartProcessAsCurrentUser("d:\\temp\\sysproxy.exe", "d:\\temp\\sysproxy.exe set 1 - - -", "d:\\temp", false); err != nil {
		return err
	}
	return nil
}

var serviceFlag = flag.String("service", "", "Control the service")

func main() {
	svcConfig := &service.Config{
		Name:        "RunAsUserTest",
		DisplayName: "Run As User Test",
		Description: "Service to test launching programs as user from service",
	}

	svc := &Service{}
	s, err := service.New(svc, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if *serviceFlag != "" {
		if err := service.Control(s, *serviceFlag); err != nil {
			log.Fatal(err)
		}

		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}