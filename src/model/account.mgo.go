// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: account.proto

/*
It has these top-level messages:
	Account
*/

package model

import "fmt"
import "encoding/json"
import "sync"
import "github.com/trist725/myleaf/db/mongodb"
import "gopkg.in/mgo.v2"

var _ = fmt.Sprintf
var _ = json.Marshal
var _ *sync.Pool
var _ *mongodb.DialContext
var _ *mgo.DBRef

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [Account] begin

/// 帐号表 @collection
type Account struct {
	/// mongodb默认主键_id做账号id @bson=_id
	ID int64 `bson:"_id"`
	/// 客户端传来的唯一id,如微信unionID
	UID string `bson:"UID"`
	/// 密码 string Password = 4; / 注册时间
	RegisterTime int64 `bson:"RegisterTime"`
	///token校验    string accessToken = 6;    string refreshToken = 7; /登陆地理位置
	Location string `bson:"Location"`
	///密码
	Password string `bson:"Password"`
	/// 帐号状态 1=游客, 2=注册, 3=绑定 E_AccountState State = 6; / 密钥 string Token = 7; / 上次服务器ID int32 LastLoginServerID = 8; / 登录过服务器ID列表 repeated int32 LoginList = 9; / 渠道名 string ChannelName = 10; / 渠道帐号 string ChannelAccount = 11; / 封号标记,1为被封
	Ban int32 `bson:"Ban"`
}

func New_Account() *Account {
	m := &Account{}
	return m
}

func (m Account) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"Account\":%s}", string(ba))
}

func (m *Account) Reset() {
	m.ID = 0
	m.UID = ""
	m.RegisterTime = 0
	m.Location = ""
	m.Password = ""
	m.Ban = 0

}

func (m Account) Clone() *Account {
	n, ok := g_Account_Pool.Get().(*Account)
	if !ok || n == nil {
		n = &Account{}
	}

	n.ID = m.ID
	n.UID = m.UID
	n.RegisterTime = m.RegisterTime
	n.Location = m.Location
	n.Password = m.Password
	n.Ban = m.Ban

	return n
}

func Clone_Account_Slice(dst []*Account, src []*Account) []*Account {
	for _, i := range dst {
		Put_Account(i)
	}
	dst = []*Account{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func FindOne_Account(session *mongodb.Session, query interface{}) (one *Account, err error) {
	one = Get_Account()
	err = session.DB(dbName).C(TblAccount).Find(query).One(one)
	if err != nil {
		Put_Account(one)
		return nil, err
	}
	return
}

func FindSome_Account(session *mongodb.Session, query interface{}) (some []*Account, err error) {
	some = []*Account{}
	err = session.DB(dbName).C(TblAccount).Find(query).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func UpdateSome_Account(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblAccount).UpdateAll(selector, update)
	return
}

func Upsert_Account(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblAccount).Upsert(selector, update)
	return
}

func UpsertID_Account(session *mongodb.Session, id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblAccount).UpsertId(id, update)
	return
}

func (m Account) Insert(session *mongodb.Session) error {
	return session.DB(dbName).C(TblAccount).Insert(m)
}

func (m Account) Update(session *mongodb.Session, selector interface{}, update interface{}) error {
	return session.DB(dbName).C(TblAccount).Update(selector, update)
}

func (m Account) UpdateByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblAccount).UpdateId(m.ID, m)
}

func (m Account) RemoveByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblAccount).RemoveId(m.ID)
}

var g_Account_Pool = sync.Pool{}

func Get_Account() *Account {
	m, ok := g_Account_Pool.Get().(*Account)
	if !ok {
		m = New_Account()
	} else {
		if m == nil {
			m = New_Account()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_Account(i interface{}) {
	if m, ok := i.(*Account); ok && m != nil {
		g_Account_Pool.Put(i)
	}
}

// collection [Account] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
