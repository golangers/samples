package helper

import (
	"golanger.com/utils"
	"labix.org/v2/mgo"
	"log"
	"strings"
)

type Mongo struct {
	session  *mgo.Session
	dbName   string
	colNames map[string]*mgo.Collection
}

func NewMongo(mgoDns string) *Mongo {
	DbPos := strings.LastIndex(mgoDns, "/")
	mgoAddr := mgoDns[:DbPos]
	dbName := mgoDns[DbPos+1:]
	mgoSession, err := mgo.Dial(mgoAddr)
	if err != nil {
		log.Fatal(err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return &Mongo{
		session:  mgoSession,
		dbName:   dbName,
		colNames: map[string]*mgo.Collection{},
	}
}

func (m *Mongo) C(col utils.M) *mgo.Collection {
	colName := col["name"].(string)
	var colIndex []string

	if _, ok := m.colNames[colName]; !ok {
		m.colNames[colName] = m.session.DB(m.dbName).C(colName)
		if _, okIn := col["index"]; okIn {
			colIndex = col["index"].([]string)
			err := m.colNames[colName].EnsureIndex(mgo.Index{Key: colIndex, Unique: true})
			if err != nil {
				log.Fatal(colName+".EnsureIndex:", err)
				return nil
			}
		}
	}

	return m.colNames[colName]
}

func (m *Mongo) Close() {
	m.session.Close()
}
