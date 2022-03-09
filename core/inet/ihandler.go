package inet

type IHandler interface {
	PreHandle(req *Obj)
	Handle(req *Obj)
	PostHandle(req *Obj)
}
