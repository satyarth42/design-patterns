package appendix

type S1 struct {
	F1 string
	F2 struct{
		F3 int
		F4 bool
	}
}

type CompositeStruct struct {
	field1 S1
	field2 string
}

type IBuilder interface {
	setF1(s string)
	setF3(i int)
	setF4(val bool)
	setField2(s string)
	Build() *CompositeStruct
}

type Builder struct {
	F1 string
	F3 int
	F4 bool
	field2 string
}

func (b *Builder) setF1(s string) {
	b.F1 = s
}

func (b *Builder) setF3(i int) {
	b.F3 = i
}

func (b *Builder) setF4(val bool) {
	b.F4 = val
}

func (b *Builder) setField2(s string) {
	b.field2 = s
}

func (b *Builder) Build() *CompositeStruct {
	return &CompositeStruct{
		field1: S1{
			F1: b.F1,
			F2: struct {
				F3 int
				F4 bool
			}{F3: b.F3, F4: b.F4},
		},
		field2: b.field2,
	}
}

func (b *Builder) GetBuilder() IBuilder {
	return &Builder{}
}

