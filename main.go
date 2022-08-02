package main

import (
	"flag"
	"log"

	"github.com/behh/locadora/api"
	"github.com/kardianos/service"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Rodando no terminal.")
	} else {
		logger.Info("Running pelo gerenciador de servicos.")
	}
	p.exit = make(chan struct{})

	go p.run()
	return nil
}
func (p *program) run() error {
	logger.Infof("Rodando na Plataforma %v.", service.Platform())
	api.InitAPI()
	return nil
}
func (p *program) Stop(s service.Service) error {
	logger.Info("Parando Servico")
	close(p.exit)
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "Locadora",
		DisplayName: "Locadora",
		Description: "Esse servico faz a interface da Locadora de Veículos.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
