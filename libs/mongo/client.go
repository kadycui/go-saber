package mongo

import (
	"gostart/libs/dict"
	"time"

	"gopkg.in/mgo.v2"
)

type Client struct {
	session *mgo.Session
	db      string
}

func NewClient(mongoConfig map[string]interface{}) (*Client, error) {
	addr := dict.GetString(mongoConfig, "host") + ":" + dict.GetString(mongoConfig, "port")
	user := dict.GetString(mongoConfig, "user")
	pwd := dict.GetString(mongoConfig, "password")
	db := dict.GetString(mongoConfig, "db")

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{addr},
		Username: user,
		Password: pwd,
		Timeout:  5 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	//返回数据
	client := &Client{
		session: session,
		db:      db,
	}
	return client, nil
}

func (client *Client) connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := client.session.Copy()
	c := s.DB(client.db).C(collection)
	return s, c
}

func (client *Client) Insert(collection string, docs ...interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func (client *Client) FindOne(collection string, query, selector, result interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func (client *Client) FindAll(collection string, query, selector, result interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func (client *Client) Update(collection string, query, update interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}

func (client *Client) UpdateAll(collection string, query, update interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	_, err := c.UpdateAll(query, update)
	return err
}

func (client *Client) Remove(collection string, query interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	return c.Remove(query)
}

func (client *Client) RemoveAll(collection string, query interface{}) error {
	ms, c := client.connect(collection)
	defer ms.Close()
	_, err := c.RemoveAll(query)
	return err
}
