package ext

type Dict[K comparable, V any] map[K]V

func Dict_[K comparable, V any](cap int) Dict[K, V] {
	return make(map[K]V, cap)
}

type KV[KT comparable, VT any] struct {
	K KT
	V VT
}

func KV_[K comparable, V any](k K, v V) KV[K, V] {
	return KV[K, V]{k, v}
}

// DictOf 将外部map转为Dict
func DictOf[K comparable, V any](m map[K]V) Dict[K, V] {
	return m
}

func (d Dict[K, V]) ForEach(fn func(KV[K, V])) {
	for k, v := range d {
		fn(KV_(k, v))
	}
}

// Len 求出dict长度
func (d Dict[K, V]) Len() int {
	return len(d)
}

// Empty 判断dict是否为空
func (d Dict[K, V]) Empty() bool {
	return len(d) == 0
}

// Load 判断key是否存在,并求出对应的值
func (d Dict[K, V]) Load(key K) (V, bool) {
	v, b := d[key]
	return v, b
}

// LoadOr 通过key获取v，获取不到将返回or
func (d Dict[K, V]) LoadOr(key K, or V) V {
	if v, b := d[key]; b {
		return v
	}
	return or
}

// LoadElse 通过key获取v，获取不到将通过函数返回值
func (d Dict[K, V]) LoadElse(key K, fn func() V) V {
	if v, b := d[key]; b {
		return v
	}
	return fn()
}

// Store 添加键值对
func (d Dict[K, V]) Store(key K, value V) {
	d[key] = value
}

// LoadOrStore 向Dict中添加键值对，如果key存在，则直接返回
func (d Dict[K, V]) LoadOrStore(key K, value V) (V, bool) {
	if v, b := d[key]; b {
		return v, b
	}
	d[key] = value
	return value, false
}

// LoadAndDelete 通过key删除键值对，并且返回v和b，如果key不存在则返回nil，false
func (d Dict[K, V]) LoadAndDelete(key K) (V, bool) {
	v, b := d[key]
	if b {
		delete(d, key)
	}
	return v, b
}

// Delete 删除键值对
func (d Dict[K, V]) Delete(key K) {
	delete(d, key)
}

// ToVec 将dict转为Vec
func (d Dict[K, V]) ToVec() Vec[KV[K, V]] {
	vec := Vec_[KV[K, V]](len(d))
	for k, v := range d {
		vec.Append(KV_(k, v))
	}
	return vec
}

// Keys 获取所有的key
func (d Dict[K, V]) Keys() Vec[K] {
	vec := Vec_[K](len(d))
	for k, _ := range d {
		vec.Append(k)
	}
	return vec
}

// Values 获取所有的Values
func (d Dict[K, V]) Values() Vec[V] {
	vec := Vec_[V](len(d))
	for _, v := range d {
		vec.Append(v)
	}
	return vec
}

func (d Dict[K, V]) Clear() {
	clear(d)
}

func (d Dict[K, V]) append_(kv KV[K, V]) Dict[K, V] {
	d[kv.K] = kv.V
	return d
}
