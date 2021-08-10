package head_first

type IObserver interface {
	Update(data int)
	GetID() int
}

type Observer struct {
	ID int
	Data int
}

func (o *Observer) Update(data int) {
	o.Data = data
}

func (o *Observer) GetID() int {
	return o.ID
}

type ISubject interface {
	Register(observer IObserver)
	Unregister(observer IObserver)
	NotifyObservers()
}

type Subject struct {
	Data int
	Observers map[int]IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.Observers[observer.GetID()] = observer
}

func (s *Subject) Unregister(observer IObserver) {
	delete(s.Observers,observer.GetID())
}

func (s *Subject) NotifyObservers() {
	for _,obs := range s.Observers {
		obs.Update(s.Data)
	}
}
