package main

import (
	"errors"
	"math/rand"

	"github.com/fatih/color"
)

type Dispositivo struct {
	Nombre string
	Estado bool
}

type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya está encendido")
	}
	d.Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo ya está apagado")
	}
	d.Estado = false
	return nil
}

func (d *Dispositivo) EstadoActual() string {
	if d.Estado {
		return "encendido"
	}
	return "apagado"
}

func crearDispositivos() []Dispositivo {
	return []Dispositivo{
		{Nombre: "Lampara", Estado: rand.Intn(2) == 1},
		{Nombre: "Lavarropa", Estado: rand.Intn(2) == 1},
		{Nombre: "Cerradura", Estado: rand.Intn(2) == 1},
		{Nombre: "Reloj", Estado: rand.Intn(2) == 1},
	}
}

func mostrarEstadoDispositivos(dispositivos []Controlable) {
	for _, dispositivo := range dispositivos {
		if rand.Intn(2) == 0 {
			dispositivo.Encender()
		} else {
			dispositivo.Apagar()
		}

		if dispositivo.EstadoActual() == "encendido" {
			color.Green(dispositivo.(*Dispositivo).Nombre)
		} else {
			color.Red(dispositivo.(*Dispositivo).Nombre)
		}
	}
}

func main() {
	dispositivosBase := crearDispositivos()
	var dispositivos []Controlable
	for i := range dispositivosBase {
		dispositivos = append(dispositivos, &dispositivosBase[i])
	}
	mostrarEstadoDispositivos(dispositivos)
}
