package trylock

// 排外锁的原理很简单，使用一个channel，如果能获得这个channel就是说明，可以或得到锁，不能就是不能返回false即可。
type Trylock struct {
	data chan struct{}
}

func NewTrylock() *Trylock {
	data := make(chan struct{}, 1)
	data <- struct{}{}
	return &Trylock{
		data: data,
	}
}

// 如果可以把数据取出来，那么就是证明是true，否则就是false
func (t *Trylock) Lock() bool {
	select {
	case <-t.data:
		return true
	default:
	}
	return false
}

// 为了解锁，就是把这个channle重新重入数据即可。
func (t *Trylock) Unlock() {
	select {
	case t.data <- struct{}{}:
	default:
		panic("no lock yet.")
	}
}
