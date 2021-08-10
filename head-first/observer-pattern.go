package head_first

type IObserver interface {
	Update(data int)
	GetID() int
}

type Observer struct {
	ID int
	Data int
	Subject ISubject
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


// ~~~****~~~~

type IPullableSubject interface {
	ISubject
	GetData() int
}

type PullableSubject struct {
	Data int
	Observers map[int]IPullableObserver
}

func (s *PullableSubject) Register(observer IPullableObserver) {
	s.Observers[observer.GetID()] = observer
}

func (s *PullableSubject) Unregister(observer IPullableObserver) {
	delete(s.Observers,observer.GetID())
}

func (s *PullableSubject) NotifyObservers() {
	for _,obs := range s.Observers {
		obs.Update()
	}
}

type IPullableObserver interface {
	Update()
	GetID() int
}

type PullableObserver struct {
	Data int
	ID int
	Subject IPullableSubject
}

func (o *PullableObserver) Update() {
	o.Data = o.Subject.GetData()
}

func (o *PullableObserver) GetID() int {
	return o.Data
}