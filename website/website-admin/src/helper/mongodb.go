package helper

import (
	"encoding/gob"
	"golanger.com/framework/utils"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
)

func init() {
	gob.Register(bson.M{})
}

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

	if _, ok := m.colNames[colName]; !ok {
		m.colNames[colName] = m.session.DB(m.dbName).C(colName)
		if _, okIn := col["index"]; okIn {
			if colIndexs, okType := col["index"].([]string); okType {
				for _, colIndex := range colIndexs {
					colIndexArr := strings.Split(colIndex, ",")
					err := m.colNames[colName].EnsureIndex(mgo.Index{Key: colIndexArr, Unique: false})
					if err != nil {
						log.Fatal(colName+".Index:", err)
						return nil
					}
				}
			}
		}

		if _, okIn := col["unique"]; okIn {
			if colIndexs, okType := col["unique"].([]string); okType {
				for _, colIndex := range colIndexs {
					colIndexArr := strings.Split(colIndex, ",")
					err := m.colNames[colName].EnsureIndex(mgo.Index{Key: colIndexArr, Unique: true})
					if err != nil {
						log.Fatal(colName+".Unqiue:", err)
						return nil
					}
				}
			}
		}
	}

	return m.colNames[colName]
}

func (m *Mongo) Close() {
	m.session.Close()
}
