package str

type Str interface {
	IsExit(v interface{})bool
	Insert(v interface{})
}


func StrIs(s Str,value interface{})bool{
	return s.IsExit(value)
}
