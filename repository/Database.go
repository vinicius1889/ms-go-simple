package repository

import "gopkg.in/mgo.v2"

func GetSession() *mgo.Session{
	var sessionMaster *mgo.Session=nil;
	var erro error=nil;
	sessionMaster,erro=mgo.Dial("mongo-test-ms.icarros.com.br")
	if erro!=nil{
		panic(erro)
	}
	sessionMaster.SetMode(mgo.Monotonic,true)
	return sessionMaster
}

