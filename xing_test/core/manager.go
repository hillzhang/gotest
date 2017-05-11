package core

import(
	"sync"
	"github.com/viphxin/xingo/iface"
	"errors"
)

type Mgr struct {
	PlayerNumGen int32
	Players      map[int32]*Player
	sync.RWMutex
}

var WorldMgrObj *Mgr

func init() {
	WorldMgrObj = &Mgr{
		PlayerNumGen:    0,
		Players:         make(map[int32]*Player),
	}
}

func (this *Mgr) AddPlayer(fconn iface.Iconnection) (*Player,error){
	this.Lock()
	defer this.Unlock()
	this.PlayerNumGen ++

	p := NewPlayer(fconn,this.PlayerNumGen)
	this.Players[this.PlayerNumGen] = p
	return p,nil
}

func (this *Mgr) GetPlayNums()int{
	return len(this.Players)
}

func (this *Mgr) GetPlayer(pid int32)(*Player, error){
	this.RLock()
	defer this.RUnlock()
	p, ok := this.Players[pid]
	if ok{
		return p, nil
	}else{
		return nil, errors.New("no player in the world!!!")
	}
}

func (this *Mgr)RemovePlayer(pid int32){
	this.Lock()
	defer this.Unlock()
	//从aoi移除
	delete(this.Players, pid)
}


