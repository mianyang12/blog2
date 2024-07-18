package dao

import (
	"goBlog/models"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// 导入存储库中的模型
	_ "goBlog/models"
)

type MockDB struct {
	mock.Mock
}

// 编写模拟结构体的函数
func (m *MockDB) GetAllCategorys() ([]models.Category, error) {
	// 根据调用情况配置返回值
	// 在这里，我们返回一个空切片和一个 nil 错误来模拟成功的查询
	return []models.Category{}, nil
}
func TestGetAllCategory(t *testing.T) {
	// 创建模拟数据库对象
	mockDB := new(MockDB)

	// 调用模拟方法的存根
	mockDB.On("Query", "select * from blog_category").Return(nil, nil)

	DB = mockDB

	// 调用目标函数
	categorys, err := GetAllCategory()

	// 断言结果
	assert.NoError(t, err)
	assert.NotEmpty(t, categorys)
}

func TestMain(m *testing.M) {
	// 创建 gorm 实例，连接到测试数据库
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 初始化模拟数据库对象
	DB = db

	// 确保在测试结束时，所有未处理的事务都将回滚
	db.Callback().Create().Before("gorm:create").Register("noop", func(*gorm.Scope) {})
	db.Callback().Update().Before("gorm:update").Register("noop", func(*gorm.Scope) {})
	db.Callback().Delete().Before("gorm:delete").Register("noop", func(*gorm.Scope) {})

	// 运行测试
	m.Run()
}
