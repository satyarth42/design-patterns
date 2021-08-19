package head_first

import "fmt"

type Amplifier struct{}
type Tuner struct{}
type StreamingPlayer struct{}
type Projector struct{}
type TheaterLight struct{}
type Screen struct{}
type PopcornPopper struct{}

type HomeTheatreFacade struct {
	amp Amplifier
	tuner Tuner
	projector Projector
	theaterLight TheaterLight
	screen Screen
	popcornPopper PopcornPopper
}

func (f *HomeTheatreFacade) watchMovie(movie string){
	fmt.Println("Starting: ",movie)

	//f.popcornPopper.on()
	//f.popcornPopper.pop()
	//f.light.dim()
	//f.screen.down()
	//f.projector.on()
	//....
}

func (f *HomeTheatreFacade) endMovie() {
	//f.popcornPopper.off()
	//f.light.on()
	//f.screen.up()
	//f.projector.off()
	//....
}

func getHomeTheatreFacade(amp Amplifier, tuner Tuner, projector Projector, theaterLight TheaterLight, screen Screen, popper PopcornPopper) *HomeTheatreFacade {
	return &HomeTheatreFacade{
		amp: amp,
		tuner: tuner,
		projector: projector,
		theaterLight: theaterLight,
		screen: screen,
		popcornPopper: popper,
	}
}