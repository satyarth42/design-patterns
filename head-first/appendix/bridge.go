package appendix

type IRemote interface {
	on()
	off()
	setChannel(x int)
}

type TV interface {
	on()
	off()
	tuneChannel(x int)
}
type Sony struct {
	isOn bool
	currentChannel int
}
func (s *Sony) on(){
	s.isOn = true
	s.currentChannel = 1
}
func (s *Sony) off(){
	s.isOn = false
}
func (s *Sony) tuneChannel(x int){
	s.currentChannel = x
}

type Remote struct {
	implementor TV
}
func (r *Remote) on() {
	r.implementor.on()
}
func (r *Remote) off() {
	r.implementor.off()
}
func (r *Remote) setCannel(x int) {
	r.implementor.tuneChannel(x)
}

type ConcreteRemote struct {
	Remote
	currentStation int
}

func (cr *ConcreteRemote) on(){
	cr.Remote.on()
	cr.currentStation = 1
}
func (cr *ConcreteRemote) off(){
	cr.Remote.off()
}
func (cr *ConcreteRemote) setChannel(x int){
	cr.Remote.setCannel(x)
	cr.currentStation = x
}
func (cr *ConcreteRemote) nextChannel(){
	cr.currentStation += 1
	cr.Remote.setCannel(cr.currentStation)
}
func (cr *ConcreteRemote) previousChannel(){
	cr.currentStation -= 1
	cr.Remote.setCannel(cr.currentStation)
}

func NewConcreteRemote() *ConcreteRemote {
	tv := &Sony{
		isOn: false,
		currentChannel: 1,
	}

	remote := &ConcreteRemote{
		Remote:Remote{
			implementor: tv,
		},
	}
	return remote
}