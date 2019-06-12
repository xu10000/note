package help

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

type Temp struct {
	Name    string
	CapName string
}

var routerTemplate = `
'use strict';

/**
 * @param {Egg.Application} app - egg application
 */
module.exports = app => {
    const { router, controller } = app;
    // 添加 删除 编辑 列表 单点查询
    router.resources('/api/v1/{{.Name}}', controller.{{.Name}});
};
`
var ruleString = `
{
    "create": {

    },
    "update": {
        
    },
    "createMany": {
        "objs": {
            "type": "array",
            "rule": {
                
            }
        }
    },
    "deleteMany": {

    },
    "updateMany": {
        "fields":{

        },
        "where":{

        }
    }
}
`
var controllerTemplate = `
'use strict';

const Controller = require('egg').Controller;
const createRule = require('../rules/test.json').create;
const createManyRule = require('../rules/test.json').createMany;
const deleteManyRule = require('../rules/test.json').deleteMany;
const updateManyRule = require('../rules/test.json').updateMany;
const updateRule = require('../rules/test.json').update;

class {{.CapName}}Controller extends Controller {
    /**
     * 创建
     */
    async create() {
        const { ctx, service } = this;
        try {
            ctx.validate(createRule);
            await service.{{.Name}}.create(ctx.request.body);
            ctx.successful('创建成功', { success: true, isData: false });
        } catch(err) {
            err = '创建失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
    }
    /**
     * 删除
     */
    async destroy() {
        const { ctx, service } = this;
        try {
            await service.{{.Name}}.destroy(ctx.params.id);
            ctx.successful('删除成功', { isData: false });
        } catch(err) {
            err = '删除失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
    /**
     * 修改
     */
    async update() {
        const { ctx, service } = this;
        try {
            ctx.validate(updateRule);
            await service.{{.Name}}.update(ctx.params.id, ctx.request.body);
            ctx.successful('更新成功', { isData: false });
        } catch(err) {
            err = '更新失败 ' + JSON.stringify(err);
            ctx.failed({ message: err });
        }
        
    }
    /**
     * 根据id查询记录
     */
    async show() {
        const { ctx, service } = this;
        
        try {
            const data = await service.{{.Name}}.show(ctx.params.id);
            ctx.successful(data);
        } catch(err) {
            err = '查询失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
    }
    /**
     * 查询所有
     */
    async index() {
        const { ctx, service } = this;
        try {
            const { page, size } = ctx.query;
            const data = await service.{{.Name}}.index(page, size);

            ctx.successful(data);
        } catch(err) {
            err = '查询所有失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
    /**
     * 批量创建
     */
    async createMany() {
        const { ctx, service } = this;
        try {
            ctx.validate(createManyRule);
            await service.{{.Name}}.createMany(ctx.request.body.objs);
            ctx.successful('批量创建成功', { isData: false });
        } catch(err) {
            err = '批量创建失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
    /**
     * 批量删除
     */
    async deleteMany() {
        const { ctx, service } = this;
        try {
            ctx.validate(deleteManyRule);
        await service.{{.Name}}.deleteMany(ctx.request.body);
        ctx.successful('批量删除成功', { isData: false });
        } catch(err) {
            err = '批量删除失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
    /**
     * 批量更新
     */
    async updateMany() {
        const { ctx, service } = this;
        try {
            ctx.validate(updateManyRule);
            const { fields, where } = ctx.request.body;
            await service.{{.Name}}.updateMany(fields, where);
            ctx.successful('批量更新成功', { isData: false });
        } catch(err) {
            err = '批量更新失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
    /**
     * 条件查询，返回一条记录
     */
    async findOne() {
        const { ctx, service } = this;
        try {
            const data = await service.{{.Name}}.findOne(ctx.request.body);
            ctx.successful(data);
        } catch(err) {
            err = '条件查询findOne失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
      
    }
    /**
     * 条件查询,返回所有符合条件的记录
     */
    async findByExample() {
        const { ctx, service } = this;
        try {
            const { page, size } = ctx.query;
            const data = await service.{{.Name}}.findByExample(page, size, ctx.request.body);
            ctx.successful(data);
        } catch(err) {
            err = '条件查询失败 ' + JSON.stringify(err);
            ctx.failed({message: err });
        }
        
    }
}

module.exports = {{.CapName}}Controller;
`

var serviceTemplate = `
'use strict';

const Service = require('egg').Service;

class {{.CapName}}Service extends Service {
    /**
     * 创建
     * @param {object} obj 项目实体
     */
    async create(obj) {
        
        await this.ctx.model.{{.CapName}}.create(obj);
    }
    /**
     * 删除
     * @param {string} id id
     */
    async destroy(id) {
        const obj = await this.show(id);
        await obj.destroy();
    }
    /**
     * 修改
     * @param {string} id id
     * @param {object} new_obj 新记录
     */
    async update(id, new_obj) {
        const old_obj = await this.show(id);
        await old_obj.update(new_obj);
    }
    /**
     * 根据id查询记录
     * @return {object} 对象实体
     * @param {string} id id
     */
    async show(id) {
        id = parseInt(id)
        console.log("asd" + this.ctx.model.Test)

        const obj = await this.ctx.model.{{.CapName}}.findOne({where: {id}});

        if (!obj) {
            throw "不存在id为" + id + "的记录";
        }
        return obj;
    }
    /**
     * 查询所有
     * @return {object} 多个对象实体数组
     * @param {int} page 页码
     * @param {int} size 条数限制
     */
    async index(page, size) {
        page = parseInt(page);
        size = parseInt(size);
        return await this.ctx.model.{{.CapName}}.findAndCountAll({
            offset: (page - 1) * size,
            limit: size,
            raw: true
        });
    }
    /**
     * 批量创建
     * @param {Array} objs 实体对象数组
     */
    async createMany(objs) {
        objs.forEach(async item => {
            await this.ctx.model.{{.CapName}}.create(item);
        });
    }
    /**
     * 批量删除
     * @param {object} where 条件
     */
    async deleteMany(where) {
        await this.ctx.model.{{.CapName}}.destroy({
            where
        });
    }
    /**
     * 批量更新
     * @param {object} fields 字段信息
     * @param {*} where 条件
     */
    async updateMany(fields, where) {
        await this.ctx.model.{{.CapName}}.update(fields, {
            where
        });
    }
    /**
     * 根据条件查询，返回首条记录
     * @return {object} 实体对象
     * @param {object} where 条件
     */
    async findOne(where) {
        const obj = await this.ctx.model.{{.CapName}}.findOne({
            where,
            raw: true
        });
        if (!obj) {
            throw  "不存在条件为" + JSON.stringify(where) + "的记录";
        }
        return obj;
    }
    /**
     * 根据条件查询，返回所有符合条件的记录
     * @return {object} 包含多个实体对象
     * @param {int} page 页码
     * @param {int} size 条数限制
     * @param {*} where 条件
     */
    async findByExample(page, size, where) {
        page = parseInt(page);
        size = parseInt(size);
        return await this.ctx.model.{{.CapName}}.findAndCountAll({
            offset: (page - 1) < 0 ? 0 : (page - 1) * size,
            limit: size,
            where,
            raw: true
        });
    }
}

module.exports = {{.CapName}}Service;
`

var TicCrud = cli.Command{
	Name:        "crud",
	Usage:       "usage: create the duplicate file",
	Description: "description: create the duplicate file",
	Flags:       []cli.Flag{},
	Action: func(c *cli.Context) error {

		fmt.Println("begin create crud api, you must be in egg project first catalog!")
		// 判断是否在第一层目录以是否有package.json为准
		_, err := os.Stat("./package.json")

		if err != nil {
			panic("you must be in egg project first catalog!")
		}

		modelCatalog, err := filepath.Glob("./app/model/*")

		if err != nil {
			panic("there is not catalog name model!")
		}

		// 创建通用目录
		createDir()

		// 文件是否存在下划线
		fixReg := regexp.MustCompile(`_`)
		// 遍历每个文件
		reg := regexp.MustCompile(`/\S*/(\S*).js`)
		for _, modelPath := range modelCatalog {
			// 判断命名是否规范
			flag := len(fixReg.Find([]byte(modelPath)))
			if flag != 0 {
				panic("文件命名不允许有下划线: " + modelPath)
			}
			//获取文件名和路径
			fName := reg.FindStringSubmatch(modelPath)[1]
			routerPath := "./app/router/" + fName + ".js"
			rulePath := "./app/rules/" + fName + ".json"
			controllerPath := "./app/controller/" + fName + ".js"
			servicePath := "./app/service/" + fName + ".js"

			// 判断router文件是否存在
			exist, _ := filepath.Glob(routerPath)
			// router文件不存在就创建router, controller, service三个文件
			if length := len(exist); length == 0 {
				createRouter(fName, routerPath)
				createRule(rulePath)
				createController(fName, controllerPath)
				createService(fName, servicePath)
				fmt.Println("创建router controller service文件: ", fName, " 成功")
			}

		}

		return nil
	},
}

func createRouter(fName string, path string) {
	// 创建模板
	t, err := template.New(fName).Parse(routerTemplate)

	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("创建模板失败： " + string(errByte))
	}
	// 生成模板的具体内容
	reg, _ := regexp.Compile("^[a-z]")
	capName := reg.ReplaceAllStringFunc(fName, strings.ToUpper)
	r := Temp{fName, capName}
	d1 := &bytes.Buffer{}
	t.Execute(d1, r)
	// 将模板写入到文件
	write := []byte(d1.String())
	err = ioutil.WriteFile(path, write, 0777)
	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("写入router文件出错： " + string(errByte))
	}
}

func createRule(path string) {
	// 将模板写入到文件
	write := []byte(ruleString)
	err := ioutil.WriteFile(path, write, 0777)
	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("写入rule文件出错： " + string(errByte))
	}
}

func createController(fName string, path string) {
	// 创建模板
	t, err := template.New(fName).Parse(controllerTemplate)

	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("创建模板失败： " + string(errByte))
	}
	// 生成模板的具体内容
	reg, _ := regexp.Compile("^[a-z]")
	capName := reg.ReplaceAllStringFunc(fName, strings.ToUpper)
	r := Temp{fName, capName}
	d1 := &bytes.Buffer{}
	t.Execute(d1, r)
	// 将模板写入到文件
	write := []byte(d1.String())
	err = ioutil.WriteFile(path, write, 0777)
	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("写入controller文件出错： " + string(errByte))
	}
}

func createService(fName string, path string) {
	// 创建模板
	t, err := template.New(fName).Parse(serviceTemplate)

	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("创建模板失败： " + string(errByte))
	}
	// 生成模板的具体内容
	reg, _ := regexp.Compile("^[a-z]")
	capName := reg.ReplaceAllStringFunc(fName, strings.ToUpper)
	r := Temp{fName, capName}
	d1 := &bytes.Buffer{}
	t.Execute(d1, r)
	// 将模板写入到文件
	write := []byte(d1.String())
	err = ioutil.WriteFile(path, write, 0777)
	if err != nil {
		errByte, _ := json.Marshal(err)
		panic("写入service文件出错： " + string(errByte))
	}
}

/**
	如果存在则不创建
**/
func createDir() {
	dirs := []string{"router", "rules", "service"}

	for _, dir := range dirs {
		path := "./app/" + dir
		catalog, _ := filepath.Glob(path)
		if length := len(catalog); length == 0 {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				panic("创建文件夹失败～")
			}
		}
	}

}
