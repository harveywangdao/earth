package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	_DB_NAME        = "data/beego.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         int64 `orm:"index"`
	Views           int64 `orm:"index"`
	TopicTime       int64 `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	TopicCategory   string
	Labels          string
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	beego.Trace(name)
	o := orm.NewOrm()
	category := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(category)
	if err == nil {
		return err
	}

	_, err = o.Insert(category)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	categorys := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&categorys)

	return categorys, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	category := &Category{Id: cid}
	_, err = o.Delete(category)
	if err != nil {
		return err
	}

	return nil
}

func AddTopic(title, content, label, topicCategory, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	o := orm.NewOrm()
	topic := &Topic{
		Title:         title,
		Content:       content,
		Labels:        label,
		Created:       time.Now(),
		Updated:       time.Now(),
		ReplyTime:     time.Now(),
		Attachment:    attachment,
		TopicCategory: topicCategory,
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	beego.Trace(title)

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", topicCategory).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

func GetAllTopics(cate, label string, isDesc bool) ([]*Topic, error) {
	var err error
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qt := o.QueryTable("topic")

	if isDesc {
		if len(cate) > 0 {
			qt = qt.Filter("TopicCategory", cate)
		}

		if len(label) > 0 {
			qt = qt.Filter("labels__contains", "$"+label+"#")
		}

		_, err = qt.OrderBy("-created").All(&topics)
	} else {
		_, err = qt.All(&topics)
	}

	return topics, err
}

func GetOneTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)

	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)

	return topic, err
}

func ModifyTopic(tid, title, content, label, topicCategory, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	var oldCate, oldAttach string

	tidNum, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.TopicCategory
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Content = content
		topic.Labels = label
		topic.TopicCategory = topicCategory
		topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	beego.Trace("oldCate:" + oldCate)
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	beego.Trace("Old Attachment:" + oldAttach)
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	beego.Trace("New Cate:" + topicCategory)
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", topicCategory).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

func DeleteOneTopic(id string) error {
	var oldCate string

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: cid}
	if o.Read(topic) == nil {
		oldCate = topic.TopicCategory
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	return nil
}

func AddReply(tid, nickname, replycontent string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: replycontent,
		Created: time.Now(),
	}

	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	beego.Trace(tidNum, nickname)

	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err == nil {
		topic.ReplyCount++
		topic.ReplyTime = time.Now()
		_, err = o.Update(topic)
	}

	return err
}

func GetAllReplies(tid string) ([]*Comment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("Tid", tidNum).All(&replies)

	return replies, err
}

func DeleteOneReply(id string) error {
	var tid int64
	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	reply := &Comment{Id: rid}
	if o.Read(reply) == nil {
		tid = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tid).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tid}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}

	return nil
}
