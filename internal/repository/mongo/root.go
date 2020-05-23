package mongo

import (
	"log"
	"time"

	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sarulabs/di"
)

const MONGO = "mongo"
const DB_NAME = "todo"

func init() {
	config.AddDef(di.Def{
		Name:  MONGO,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(config.CONFIG).(*config.Config)
			s, err := mgo.Dial(cfg.Mongo.Address)
			if err != nil {
				log.Fatal(err)
			}
			return s, nil
		},
	})
}

type Creatable interface {
	SetID(string)
	SetCreated(time.Time)
}

type mongo struct {
	session        *mgo.Session
	db             string
	collectionName string
}

func (m *mongo) getCollection() (*mgo.Collection, func()) {
	session := m.session.Copy()
	return session.DB(m.db).C(m.collectionName), session.Close
}

func (m *mongo) findByFilter(filter interface{}, result interface{}) error {
	collection, close := m.getCollection()
	defer close()

	return collection.Find(filter).One(result)
}

func (m *mongo) findAllByFilter(filter interface{}, result interface{}) error {
	collection, close := m.getCollection()
	defer close()

	return collection.Find(filter).All(result)

}

func (m *mongo) insert(data interface{}) error {
	collection, close := m.getCollection()
	defer close()

	if d, ok := data.(Creatable); ok {
		d.SetID(bson.NewObjectId().Hex())
		d.SetCreated(time.Now())
	}
	return collection.Insert(data)
}
