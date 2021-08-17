package head_first

//command
type ICommand interface {
	Execute()
	Undo()
}

//receiver
type Light struct {
	state bool
}
func(l *Light) setOn() {
	l.state = true
}
func(l *Light) setOff() {
	l.state = false
}

//receiver
type Stereo struct {
	volume int
	cd string
	radio string
	power bool
}
func (s *Stereo) on() {
	s.power = true
}
func (s *Stereo) off() {
	s.power = false
}
func (s *Stereo) setVolume(volume int) {
	s.volume = volume
}
func (s *Stereo) setCD(cd string) {
	s.cd = cd
}
func (s *Stereo) setRadio(radio string) {
	s.radio = radio
}

//concrete command
type LightOnCommand struct {
	light *Light
}
func (lc *LightOnCommand) Execute(){
	lc.light.setOn()
}
func (lc *LightOnCommand) Undo(){
	lc.light.setOff()
}

//concrete command
type LightOffCommand struct {
	light *Light
}
func (lc *LightOffCommand) Execute(){
	lc.light.setOff()
}
func (lc *LightOffCommand) Undo(){
	lc.light.setOn()
}

//concrete command
type StereoOnWithCD struct {
	stereo *Stereo
}
func (s *StereoOnWithCD) Execute() {
	s.stereo.on()
	s.stereo.setCD("something")
	s.stereo.setVolume(10)
}
func (s *StereoOnWithCD) Undo() {
	s.stereo.off()
}

//concrete command
type StereoOff struct {
	stereo *Stereo
}
func (s *StereoOff) Execute() {
	s.stereo.off()
}
func (s *StereoOff) Undo() {
	s.stereo.on()
	s.stereo.setCD("something")
	s.stereo.setVolume(10)
}

//invoker
type RemoteControl struct {
	onCommands []ICommand
	offCommands []ICommand
	undoCommand ICommand
}
func (rc *RemoteControl) setCommand(slot int, onCommand, offCommand ICommand) {
	rc.onCommands[slot] = onCommand
	rc.offCommands[slot]= offCommand
}
func (rc *RemoteControl) onButtonPushed(slot int) {
	rc.onCommands[slot].Execute()
	rc.undoCommand = rc.onCommands[slot]
}
func (rc *RemoteControl) offButtonPushed(slot int) {
	rc.offCommands[slot].Execute()
	rc.undoCommand = rc.offCommands[slot]
}
func (rc *RemoteControl) undoButtonPushed(){
	rc.undoCommand.Undo()
}

func GetRemoteController() *RemoteControl {
	rc := new(RemoteControl)
	rc.onCommands = make([]ICommand,3)
	rc.offCommands = make([]ICommand,3)

	return rc
}

func initialize() {
	rc := GetRemoteController()

	kitchenLight := new(Light)
	kitchenLightOnCommand := &LightOnCommand{light: kitchenLight}
	kitchenLightOffCommand := &LightOffCommand{light: kitchenLight}
	rc.setCommand(0, kitchenLightOnCommand, kitchenLightOffCommand)

	bedroomLight := new(Light)
	bedroomLightOnCommand := &LightOnCommand{light: bedroomLight}
	bedroomLightOffCommand := &LightOffCommand{light: bedroomLight}
	rc.setCommand(1, bedroomLightOnCommand, bedroomLightOffCommand)

	stereo := new(Stereo)
	steroOnWithCDCommand := &StereoOnWithCD{stereo: stereo}
	steroOffCommand := &StereoOff{stereo: stereo}
	rc.setCommand(2,steroOnWithCDCommand,steroOffCommand)
}