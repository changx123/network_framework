package websocket_route

import (
	"encoding/binary"
	"github.com/changx123/websocket-sync"
	"bytes"
)

//钩子常量
const (
	//新连接通知
	HOOK_NEW_CONN = "new conn"

	//连接closed 通知
	HOOK_CLOSED = "closed"

	//错误通知
	HOOK_ERROR = "error"

	//路由寻址不存在
	HOOK_NOT_MODULE = "not module"

	//发送消息通知
	HOOK_WRITE_MESSAGE = "write message"

	//接收消息解包
	HOOK_UN_PACKING = "read message un packing"

	//发送消息封包
	HOOK_PACKET = "send message packing"
)

type HandlerFunc func(conn *websocket.Conn, msg *Message, route *interface{}) error

//路由批量注册结构指针
type HandlerGroup struct {
	//除钩子外 对象计数
	count int
	//路由单指针寻址
	HandlersRoute []HandlerFunc
	//路由寻址
	ModuleRoutes map[interface{}][]int
}

type Storage interface {
	Use(...HandlerFunc)
	Route(interface{}, ...HandlerFunc)
	Group(interface{}, func() *HandlerGroup)
	Hook(string, HandlerFunc)
}

type Message struct {
	MessageType int
	P           []byte
	Err         error
}

//端绪设置 默认大端序
var Endian = binary.BigEndian

//存放路由主结构
type StorageGroup struct {
	HandlerGroup
	//中间件路由寻址
	HandlersUse []int
	//钩子注册存放
	HandlerHook map[string]HandlerFunc
	//端绪设置
}

//检查是否实现所有接口方法
var _ Storage = &StorageGroup{}

//获取一个新的路由对象
func NewRouter() *StorageGroup {
	out := StorageGroup{}
	out.ModuleRoutes = make(map[interface{}][]int)
	out.HandlerHook = make(map[string]HandlerFunc)
	out.count = 0
	//HOOK_UN_PACKING
	out.Hook(HOOK_UN_PACKING,hookUnPacking)
	//HOOK_PACKET
	out.Hook(HOOK_PACKET,hookPacket)
	return &out
}

//注册中间件
func (r *StorageGroup) Use(funcs ...HandlerFunc) {
	//记录每一个路由地址id
	indexs := make([]int, len(funcs))
	for k, v := range funcs {
		//注册到总路由寻址列表
		r.HandlersRoute = append(r.HandlersRoute, v)
		indexs[k] = r.count
		r.count++
	}
	for _, v := range indexs {
		r.HandlersUse = append(r.HandlersUse, v)
	}
}

//单个注册路由列表
func (r *StorageGroup) Route(m interface{}, funcs ...HandlerFunc) {
	indexs := make([]int, len(funcs))
	for k, v := range funcs {
		r.HandlersRoute = append(r.HandlersRoute, v)
		indexs[k] = r.count
		r.count++
	}
	for _, v := range indexs {
		r.ModuleRoutes[m] = append(r.ModuleRoutes[m], v)
	}
}

//批量分组注册路由列表
func (r *StorageGroup) Group(m interface{}, f func() *HandlerGroup) {
	group := f()
	for k, v := range group.ModuleRoutes {
		indexs := make([]int, len(v))
		for key, val := range v {
			r.HandlersRoute = append(r.HandlersRoute, group.HandlersRoute[val])
			indexs[key] = r.count
			r.count++
		}
		for _, val := range indexs {
			kname := []interface{}{m, k}
			r.ModuleRoutes[kname] = append(r.ModuleRoutes[kname], val)
		}
	}
}

//注册钩子
func (r *StorageGroup) Hook(s string, f HandlerFunc) {
	r.HandlerHook[s] = f
}

//获取批量分组注册路由对象
func NewHandlerGroup() *HandlerGroup {
	out := HandlerGroup{}
	out.ModuleRoutes = make(map[interface{}][]int)
	out.count = 0
	return &out
}

//增加注册一个路由
func (h *HandlerGroup) Add(a interface{}, funcs ...HandlerFunc) {
	indexs := make([]int, len(funcs))
	for k, v := range funcs {
		h.HandlersRoute = append(h.HandlersRoute, v)
		indexs[k] = h.count
		h.count++
	}
	for _, v := range indexs {
		h.ModuleRoutes[a] = append(h.ModuleRoutes[a], v)
	}
}

func (r *StorageGroup) hookRun(n string, conn *websocket.Conn, msg *Message, route *interface{}) error {
	f, ok := r.HandlerHook[n]
	if ok {
		err := f(conn, msg, route)
		return err
	}
	return nil
}

//监听数据
func (r *StorageGroup) Listen(conn *websocket.Conn) error {
	//触发 HOOK_NEW_CONN（新连接通知钩子）
	err := r.hookRun(HOOK_NEW_CONN, conn, nil, nil)
	if err != nil {
		return err
	}
ERROR_BREAK:
	for {
		messageType , p , err := conn.ReadMessage()
		var message *Message
		message.MessageType = messageType
		message.P = p
		message.Err = err
		if err != nil {
			//触发 HOOK_ERROR （错误通知 钩子）
			if err := r.hookRun(HOOK_ERROR, conn, message, nil) ;err != nil {
				return err
			}
			//触发 HOOK_CLOSED （连接closed 通知钩子）
			if err := r.hookRun(HOOK_CLOSED, conn, message, nil);err != nil {
				return err
			}
			break
		}
		var route *interface{}
		if err := r.hookRun(HOOK_UN_PACKING, conn, message, route) ;err != nil {
			return err
		}
		//运行中间件
		useFuncs := r.HandlersUse
	ERROR_CONTINUE:
		for _, v := range useFuncs {
			err := r.HandlersRoute[v](conn, message, route)
			if err != nil {
				switch err {
				case ERROR_STOP:
					//在此终止程序继续往下执行
					continue ERROR_BREAK
				case ERROR_CONTINUE:
					//在此跳过这个中间件继续执行
					continue ERROR_CONTINUE
				case ERROR_BREAK:
					//在此跳过中间件继续执行后续逻辑
					break ERROR_CONTINUE
				default:
					//在此跳过中间件继续执行后续逻辑
					continue ERROR_CONTINUE
				}
			}
		}
		routeFuncs, ok := r.ModuleRoutes[*route]
		if ok {
		ERROR_CONTINUE_TOW:
			for _, v := range routeFuncs {
				err := r.HandlersRoute[v](conn, message, route)
				if err != nil {
					switch err {
					case ERROR_STOP:
						//在此终止程序继续往下执行
						continue ERROR_BREAK
					case ERROR_CONTINUE:
						//在此继续执行下一个路由函数
						continue ERROR_CONTINUE_TOW
					case ERROR_BREAK:
						//在此跳出路由函数执行继续后面逻辑
						break ERROR_CONTINUE_TOW
					default:
						//在此跳出路由函数执行继续后面逻辑
						break ERROR_CONTINUE_TOW
					}
				}
			}
		} else {
			//触发 HOOK_NOT_MODULE （路由寻址不存在 通知钩子）
			if err := r.hookRun(HOOK_NOT_MODULE, conn, message, route) ;err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

//HOOK_UN_PACKING
func hookUnPacking(conn *websocket.Conn, msg *Message, route *interface{}) error {
	bytesBuffer := bytes.NewBuffer(msg.P)
	var module, action uint16
	//uint 16  Module
	binary.Read(bytesBuffer, Endian, &module)
	//uint 16  Action
	binary.Read(bytesBuffer, Endian, &action)
	msg.P = bytesBuffer.Bytes()
	var kname interface{}
	kname = []uint16{module, action}
	route = &kname
	return nil
}

//HOOK_PACKET
func hookPacket(conn *websocket.Conn, msg *Message, route *interface{}) error {
	//触发 HOOK_WRITE_MESSAGE （发送详细前 通知钩子 便于开发调试）
	bytesBuffer := bytes.NewBuffer([]byte{})
	b := *route
	rou := b.([]uint16)
	//uint 16  Module
	binary.Write(bytesBuffer, Endian, &rou[0])
	//uint 16  Action
	binary.Write(bytesBuffer, Endian, &rou[1])
	//[]byte MessageValue
	binary.Write(bytesBuffer, Endian, &msg.P)
	msg.P = bytesBuffer.Bytes()
	return nil
}
