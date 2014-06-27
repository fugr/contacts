package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	qprintable "github.com/sloonz/go-qprintable"
)

type Contact struct {
	Id        int64
	Name      string
	Telephone string
	Address   string
}

type Contacts []*Contact

func NewContacts() Contacts {
	/*
		names := []string{
			"",
			"陈晖", "陈仁凤", "陈欣", "陈玉钗", "邓志华",
			"付光荣", "郭亮", "郭燕琴", "何林波", "何迎盛",
			"蒋珩", "金新春", "瞿琦", "康斯海", "兰斌",
			"黎亭亭", "李乐源", "李涛", "李瑶", "李园园",
			"廖丽平", "廖逸", "廖振宇", "刘娜", "刘青",
			"刘小峰", "刘业棕", "刘银华", "柳春富", "卢陶",
			"鲁静", "倪能", "聂小波", "彭欣荣", "彭雄",
			"邵玉莲", "粟蛟君", "汤朋", "唐鸿飞", "王剑",
			"王莺", "吴平", "伍星", "习先星", "夏云",
			"肖俊浩", "谢启新", "谢文汉", "许浒", "叶瑞斌",
			"张飞军", "张文标", "张燕群", "张耀平", "张育莲",
			"赵双", "周凤婷", "周星", "邹强"}

		var con [60]Contact
		cons := make([]*Contact, 60)

		for i, _ := range names {
			con[i].Id = int64(i)
			con[i].Name = names[i]
			cons[i] = &con[i]
		}

		return cons
	*/
	filename := "logs/contacts.log"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var list Contacts
	err = decoder.Decode(&list)
	if err != nil {
		panic("decoder " + err.Error())
	}
	fmt.Println("read contacts init done!")
	return list
}

var contactsList = NewContacts()

func (this *Contacts) Post(c *Contact) {
	if int(c.Id) <= len(*this) {
		i := c.Id
		dst := (*this)[i]
		if c.Name != "" {
			dst.Name = c.Name
		}
		if c.Telephone != "" {
			dst.Telephone = c.Telephone
		}
		if c.Address != "" {
			dst.Address = c.Address
		}
		go SaveContacts(*this)
		return
	}
	fmt.Println("new contact ")
	c.Id = int64(len(*this))
	contactsList = append(*this, c)
	go SaveContacts(contactsList)
}

func (this Contacts) Delete(id int64) {
	contact := Get(id)
	if contact != nil {
		contact.Name = ""
		contact.Telephone = ""
		contact.Address = ""
	}
}

func GetAll() Contacts {
	return contactsList
}

func Get(id int64) *Contact {
	if int(id) <= len(contactsList) {
		return contactsList[id]
	}

	return nil
}

func SaveContacts(list Contacts) {
	filename := "logs/contacts.log"
	file, err := os.OpenFile(filename, os.O_WRONLY, 777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(list)
}

func WriteToVCF() *bytes.Buffer {
	var vcard string
	const vcf = `BEGIN:VCARD
VERSION:2.1
N;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:;%s;;;
FN;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:%s
TEL;CELL;PREF:%s
END:VCARD`

	for _, contact := range contactsList {
		if contact.Telephone == "" {
			continue
		}
		encBuf := bytes.NewBuffer(nil)
		enc := qprintable.WindowsTextEncoding
		encoder := qprintable.NewEncoder(enc, encBuf)
		_, err := encoder.Write([]byte(contact.Name))
		if err != nil {
			fmt.Println("encoder error:", err)
			continue
		}
		vcard += fmt.Sprintf(vcf, encBuf, encBuf, contact.Telephone)
	}
	return bytes.NewBufferString(vcard)
}
