package handler

import (
	"context"
	"goodsman/config"
	"goodsman/db"
	"goodsman/model"
	"goodsman/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func AddManager(c *gin.Context) {
	addManagerReq := model.ModifyManagerRequest{}
	if err := c.BindJSON(&addManagerReq); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	if addManagerReq.Token != config.App.AddManagerToken {
		logrus.Error("Wrong AddManagerToken")
		response.Error(c, response.AUTH_ERROR)
		return
	}

	manager := model.Manager{}
	ctx := context.TODO()
	filter := bson.D{{"employee_id", addManagerReq.EmployeeId}}
	err := db.MongoDB.ManagerColl.FindOne(ctx, filter).Decode(&manager)

	if err == nil {
		logrus.Error("This user has been a manager, new manager will cover old record")
		if err = delManager(addManagerReq.EmployeeId); err != nil {
			logrus.Error("error happened in database & ", err.Error())
			response.Error(c, response.DATABASE_ERROR)
			return
		}
	} else if err.Error() != "mongo: no documents in result" {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	manager = model.Manager{
		Employee_id: addManagerReq.EmployeeId,
		Name:        addManagerReq.Name,
	}
	if err = addNewManager(manager); err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	response.Success(c, manager)
}

func DeleteManager(c *gin.Context) {
	delManagerReq := model.ModifyManagerRequest{}
	if err := c.BindJSON(&delManagerReq); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	if delManagerReq.Token != config.App.AddManagerToken {
		logrus.Error("Wrong AddManagerToken")
		response.Error(c, response.AUTH_ERROR)
		return
	}

	if err := delManager(delManagerReq.EmployeeId); err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	response.Success(c, "")
}

func QueryManagers() (ret []string, err error) {
	ctx := context.TODO()
	result := []struct {
		ID string `bson:"employee_id"`
	}{}
	cursor, _ := db.MongoDB.ManagerColl.Find(ctx, bson.D{})
	err = cursor.All(ctx, &result)
	if err != nil {
		return
	}
	for _, item := range result {
		ret = append(ret, item.ID)
	}
	return
}

func addNewManager(newman model.Manager) error {
	_, err := db.MongoDB.ManagerColl.InsertOne(context.TODO(), &newman)
	return err
}

func delManager(empID string) error {
	ctx := context.TODO()
	filter := bson.D{{"employee_id", empID}}
	_, err := db.MongoDB.ManagerColl.DeleteOne(ctx, filter)
	return err
}
