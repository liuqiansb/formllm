package genrator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
)

var DatabaseName = "test"
var sql2StructUrl = "https://www.devtool.com/api/sql2go"
var tableName = "user"

var STRUCT_NAME string
var LOW_STRUCT_NAME string
var columns = make([]*Column, 0)

const MODEL_NAME = "#MODEL_NAME#"
const LOW_MODEL_NAME = "#LOW_MODEL_NAME#"
const FIELD_NAME = "#FIELD_NAME#"
const LOW_FIELD_NAME = "#LOW_FIELD_NAME#"
const FIELD_TYPE = "#FIELD_TYPE#"

type Column struct {
	ColumnName             string `gorm:"column:COLUMN_NAME"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT"`
	DataType               string `gorm:"column:DATA_TYPE"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
}

func (c *Column) TableName() string {
	return "information_schema.columns"
}

func TestG(t *testing.T) {
	var ddl string
	var structString string
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3309)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	if ddl, err = createDDL(db); err != nil {
		panic(err)
	}
	if structString, err = getGoStruct(ddl); err != nil {
		panic(err)
	}
	if err = createModel(structString); err != nil {
		panic(err)
	}
}

func createDDL(db *gorm.DB) (ddl string, err error) {
	if err = db.Debug().Where("table_schema= ? and table_name = ?", DatabaseName, tableName).
		Find(&columns).
		Error; err != nil {
		return
	}
	prefix := "CREATE TABLE `" + tableName + "` (\n"
	subfix := "PRIMARY KEY (`id`)\n) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"

	ddl += prefix
	for _, v := range columns {
		ddl += "`" + v.ColumnName + "`"
		if v.CharacterMaximumLength != "" {
			ddl += " " + v.DataType + "(" + v.CharacterMaximumLength + ")"
		} else {
			ddl += " " + v.DataType
		}
		if v.ColumnName == "id" {
			ddl += " AUTO_INCREMENT"
		}
		ddl += " COMMENT '" + v.ColumnComment + "',\n"
	}
	ddl += subfix
	return
}

type SqlToStructResp struct {
	Data []string `json:"data"`
	Msg  string   `json:"msg"`
	Ok   int64    `json:"ok"`
}

func getGoStruct(ddl string) (structString string, err error) {
	client := &http.Client{}
	sqlToStructResp := new(SqlToStructResp)
	var request = struct {
		JsonTag          string `json:"json_tag"`
		SQL              string `json:"sql"`
		WithColumnPrefix string `json:"with_column_prefix"`
		WithTablePrefix  string `json:"with_table_prefix"`
	}{
		JsonTag:          "true",
		SQL:              ddl,
		WithColumnPrefix: "",
		WithTablePrefix:  "",
	}
	bytesData, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", sql2StructUrl, bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBytes, sqlToStructResp)
	if err != nil {
		return
	}
	structString = sqlToStructResp.Data[0]
	return
}

func createModel(structString string) (err error) {
	var packageString = "package model"
	var importString = "import (\n\t\"time\"\n)"
	end := strings.Index(structString, "struct")
	start := strings.Index(structString, "type") + len("type")
	STRUCT_NAME = structString[start:end]
	STRUCT_NAME = strings.Trim(STRUCT_NAME, " ")
	LOW_STRUCT_NAME = strings.ToLower(STRUCT_NAME[:1]) + STRUCT_NAME[1:]
	var modelTableFunction = "func (t *" + STRUCT_NAME + ") TableName() string {\n\t return \"" + tableName + "\"\n}"
	var builderAndCondition = createBuilderAndCondition(structString)
	var fileString = packageString + "\n" + importString + "\n" + structString + "\n" + modelTableFunction + "\n" + builderAndCondition
	file, err := os.Create("../model/" + tableName + ".go")
	if err != nil {
		return
	}
	_, err = file.Write([]byte(fileString))
	if err != nil {
		return
	}
	fmt.Println("写入" + tableName + "模型成功!")
	return
}

func createBuilderAndCondition(structString string) (builderAndCondition string) {

	builderString := strings.Replace(Builder, MODEL_NAME, STRUCT_NAME, -1)
	builderString = strings.Replace(builderString, LOW_MODEL_NAME, LOW_STRUCT_NAME, -1)
	conditionString := strings.Replace(CONDITION_STRUCT, MODEL_NAME, STRUCT_NAME, -1)
	fieldsStructString := "var " + STRUCT_NAME + "Field = &struct "
	fieldColumnsMapString := "var " + STRUCT_NAME + "FieldColumn = map[string]string{\n"
	var fieldsStructDefine string
	var fieldsStructValue string
	start := strings.Index(structString, "{")
	end := strings.Index(structString, "}")
	fieldStrings := structString[start+1 : end]
	filedStringArr := strings.Split(strings.Trim(fieldStrings, "\n"), "\n")
	for k, v := range filedStringArr {
		var re1 = regexp.MustCompile(`\t+| +`)
		v = re1.ReplaceAllString(v[:strings.Index(v, "gorm")-1], "|")
		field := strings.Split(v, "|")
		fieldName := field[1]
		lowFieldName := strings.ToLower(fieldName[:1]) + fieldName[1:]
		fieldType := field[2]
		builder := strings.Replace(BuilderField, MODEL_NAME, STRUCT_NAME, -1)
		builder = strings.Replace(builder, LOW_MODEL_NAME, LOW_STRUCT_NAME, -1)
		builder = strings.Replace(builder, FIELD_NAME, fieldName, -1)
		builder = strings.Replace(builder, LOW_FIELD_NAME, lowFieldName, -1)
		builder = strings.Replace(builder, FIELD_TYPE, fieldType, -1)

		condition := strings.Replace(ConditionWhere, MODEL_NAME, STRUCT_NAME, -1)
		condition = strings.Replace(condition, LOW_MODEL_NAME, LOW_STRUCT_NAME, -1)
		condition = strings.Replace(condition, FIELD_NAME, fieldName, -1)
		condition = strings.Replace(condition, LOW_FIELD_NAME, lowFieldName, -1)
		condition = strings.Replace(condition, FIELD_TYPE, fieldType, -1)
		builderString += builder
		conditionString += condition
		fieldsStructDefine += "\t" + fieldName + " string\n"
		fieldsStructValue += "\t" + fieldName + " : \"" + fieldName + "\",\n"
		fieldColumnsMapString += "\t\"" + fieldName + "\":" + "\"" + columns[k].ColumnName + "\",\n"
	}
	fieldsStructString += "{\n" + fieldsStructDefine + "}{\n" + fieldsStructValue + "}\n"
	fieldColumnsMapString += "}\n"
	conditionString += strings.Replace(ConditionOther, MODEL_NAME, STRUCT_NAME, -1)
	var Selective = `
func (t *` + STRUCT_NAME + `) Selective(fields []string) (selective map[string]interface{}) {
	selective = make(map[string]interface{})
	for _, v := range fields {
		selective[` + STRUCT_NAME + `FieldColumn[v]] = reflect.ValueOf(t).Elem().FieldByName(v).Interface()
	}
	return selective
}
`
	builderAndCondition = builderString + conditionString + fieldsStructString + fieldColumnsMapString + Selective

	return
}

var Builder = `
func (t #MODEL_NAME#) Builder() *#MODEL_NAME#Builder {
	return &#MODEL_NAME#Builder{#LOW_MODEL_NAME#: &#MODEL_NAME#{}}
}
type #MODEL_NAME#Builder struct {
	#LOW_MODEL_NAME# *#MODEL_NAME#
}
func (t *#MODEL_NAME#Builder) Build() *#MODEL_NAME# {
	return t.#LOW_MODEL_NAME#
}
func (t *#MODEL_NAME#Builder) Flush() *#MODEL_NAME#Builder {
	t.#LOW_MODEL_NAME# = &#MODEL_NAME#{}
	return t
}
`

var BuilderField = `
func (t *#MODEL_NAME#Builder) #FIELD_NAME#(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Builder {
	t.#LOW_MODEL_NAME#.#FIELD_NAME# = #LOW_FIELD_NAME#
	return t
}`

var CONDITION_STRUCT = `
type #MODEL_NAME#Condition struct {
	Where  string
	Order  string
	Offset int64
	Limit  int64
	Params []interface{}
}
`
var ConditionWhere = `
func (t *#MODEL_NAME#Condition) And#FIELD_NAME#Eq(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# = ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# = ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#Bt(start #FIELD_TYPE#, end #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# between ? and ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#Gt(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# > ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# > ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#Lt(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# < ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# < ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#GtAndEq(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# >= ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# >= ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#LtAndEq(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# <= ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# <= ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#Like(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# like ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# like ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#In(#LOW_FIELD_NAME#s []#FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# in(?)"
	} else {
		t.Where += " #LOW_FIELD_NAME# in(?)"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#s)
	return t
}

func (t *#MODEL_NAME#Condition) And#FIELD_NAME#NotIn(#LOW_FIELD_NAME#s []#FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# not in(?)"
	} else {
		t.Where += " #LOW_FIELD_NAME# not in(?)"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#s)
	return t
}
func (t *#MODEL_NAME#Condition) And#FIELD_NAME#NotLike(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# not like ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# not like ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}
func (t *#MODEL_NAME#Condition) And#FIELD_NAME#NotBt(start #FIELD_TYPE#, end #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# not between ? and ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *#MODEL_NAME#Condition) And#FIELD_NAME#NotEq(#LOW_FIELD_NAME# #FIELD_TYPE#) *#MODEL_NAME#Condition {
	if t.Where != "" {
		t.Where += " and #LOW_FIELD_NAME# != ?"
	} else {
		t.Where += " #LOW_FIELD_NAME# != ?"
	}
	t.Params = append(t.Params, #LOW_FIELD_NAME#)
	return t
}
`
var ConditionOther = `
func (t *#MODEL_NAME#Condition) SetOrderBy(order string) *#MODEL_NAME#Condition {
	t.Order = order
	return t
}
func (t *#MODEL_NAME#Condition) SetLimit(limit int64) *#MODEL_NAME#Condition {
	t.Limit = limit
	return t
}
func (t *#MODEL_NAME#Condition) setOffset(offset int64) *#MODEL_NAME#Condition {
	t.Offset = offset
	return t
}
`
