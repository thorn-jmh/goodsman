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

func GetManagersByAuth(c *gin.Context) {
	auth := c.GetInt("auth")
	if (auth != -1) && (auth != 1) && (auth != 2) {
		logrus.Error("Wrong request without auth")
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	managers, err := queryManagersByAuth(auth)
	if err != nil {
		logrus.Error("QueryManagersByAuth error", err)
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	logrus.Info("QueryManagersByAuth...Done!")
	response.Success(c, managers)
	return
}

func UpdateManagerAuthByEid(c *gin.Context) {
	updManagerReq := model.ChangeManagerStateRequest{}
	if err := c.BindJSON(&updManagerReq); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	if newAuth := updManagerReq.NewAuth; newAuth != -1 && newAuth != 1 && newAuth != 2 {
		logrus.Error("invalid new auth")
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	superAdmin, err := queryManagerByEid(updManagerReq.SuperEid)
	if err != nil {
		logrus.Error("can't find the super admin", err)
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	if superAdmin.Auth != 2 {
		logrus.Error("this \"super admin\" is not real \"super admin\"")
		response.Error(c, response.AUTH_ERROR)
		return
	}
	_, err = queryManagerByEid(updManagerReq.ManagerEid)
	if err != nil {
		logrus.Error("can't find the manager", err)
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	err = updateManagerAuthByEid(updManagerReq.ManagerEid, updManagerReq.NewAuth)
	if err != nil {
		logrus.Error("error happen when update manager", err)
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	response.Success(c, response.Success)
}

func QueryManagers() (ret []string, err error) {
	ctx := context.TODO()
	filter := bson.D{{}}
	result := []struct {
		ID string `bson:"employee_id"`
	}{}
	cursor, _ := db.MongoDB.ManagerColl.Find(ctx, filter)
	err = cursor.All(ctx, &result)
	if err != nil {
		return
	}
	for _, item := range result {
		ret = append(ret, item.ID)
	}
	return
}

func queryManagersByAuth(auth int) (ret []model.Manager, err error) {
	ctx := context.TODO()
	filter := bson.D{{"auth", auth}}
	result := []model.Manager{}
	cursor, _ := db.MongoDB.ManagerColl.Find(ctx, filter)
	err = cursor.All(ctx, &result)
	if err != nil {
		return
	}
	for _, item := range result {
		ret = append(ret, item)
	}
	return
}

func queryManagerByEid(eid string) (ret model.Manager, err error) {
	ctx := context.TODO()
	filter := bson.D{{"employee_id", eid}}
	result := model.Manager{}
	err = db.MongoDB.ManagerColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return
	}
	ret = result
	return
}

func updateManagerAuthByEid(eid string, newAuth int) error {
	ctx := context.TODO()
	filter := bson.M{"employee_id": eid}
	update := bson.M{"$set": bson.M{"auth": newAuth}}
	updresult, err := db.MongoDB.GoodsColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	logrus.Info(updresult)
	return nil
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
