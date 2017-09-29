package controller

import (
	"net/http"
	"encoding/json"
	"../domain"
	"../repository"
	"math"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"strconv"
)

func HealthFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(domain.Status{"UP"})
}




func GetLastCalls(w http.ResponseWriter, r *http.Request){
	session := repository.GetSession();
	variables := mux.Vars(r)
	Result := []domain.Call{};
	MyRegister := domain.Register{}

	errQuery := session.DB("calls").C("registers").Find( bson.M{"key": variables["key"] }).One(&MyRegister);
	find :=bson.M{}
	if MyRegister.Key!=""{
		find = bson.M{"_id": bson.M{"$gt":bson.ObjectIdHex(MyRegister.LastId)}}
	}

	total,errTotal :=session.DB("calls").C("records").Find(find).Count()
	if(errTotal!=nil){
		panic(errTotal)
	}
	limit,_ :=strconv.Atoi(variables["limit"])
	pageable:=domain.SimplePageable{total,limit,0,0}
	errQuery = session.DB("calls").C("records").Find(find).Limit(pageable.Limit).Sort("_id").All(&Result);

	if errQuery!=nil{ panic(errQuery) }
	w.Header().Set("Content-Type","application/json")
	if(len(Result)==0) {
		json.NewEncoder(w).Encode(domain.CallVO{});
	}else {

		pages:=int( math.Ceil(float64(pageable.Total)/float64(pageable.Limit)) )
		pageable.Pages=pages
		pageable.Remain=pageable.Total-len(Result)
		json.NewEncoder(w).Encode(domain.CallVO{Result, Result[len(Result)-1].Id,pageable});
	}
}

func SetLastCall(w http.ResponseWriter , r *http.Request){
	session := repository.GetSession();
	decoder:=json.NewDecoder(r.Body)
	MyReg := domain.Register{}
	MyQuery := domain.Register{}
	decoder.Decode(&MyReg)
	session.DB("calls").C("registers").Find(bson.M{"key":MyReg.Key}).One(&MyQuery)
	if(MyQuery.Key!="") {
		MyReg.Id = MyQuery.Id
	}else{
		MyReg.Id=bson.NewObjectId()
	}
	session.DB("calls").C("registers").Insert(&MyReg)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(MyReg);
}

