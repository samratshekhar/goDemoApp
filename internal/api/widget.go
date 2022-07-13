package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"goDemoApp/internal/datastore/dao"
	"goDemoApp/internal/models/widget"
	"io/ioutil"
	"net/http"
)

type WidgetHandler struct {
	widgetDAO dao.WidgetDAO
}

func NewWidgetHandler(widgetDAO dao.WidgetDAO) WidgetHandler {
	return WidgetHandler{widgetDAO}
}

func (w WidgetHandler) CreateWidget(wr http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(wr, "Method not allowed", 405)
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	userId, err := getUserFromHeader(r)

	if err != nil {
		http.Error(wr, err.Error(), 401)
		return
	}

	widget := new(widget.Widget)
	err = json.Unmarshal(body, widget)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}
	err = ValidateWidgetData(widget)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}
	widgetExist, err := w.widgetDAO.GetWidgetJsonById(widget.Id, widget.TemplateType.String())
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	if widgetExist.Id != "" {
		http.Error(wr, "widget Already Exist. Can't create widget", 405)
		return
	}

	err = w.widgetDAO.Create(body, widget, userId)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}
	wr.Write([]byte("successfully created widget"))
}

func (w WidgetHandler) PutWidget(wr http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(wr, "Method not allowed", 405)
	}
	params := mux.Vars(r)
	widgetId := params["id"]
	templateType := params["templateType"]
	widgetExist, err := w.widgetDAO.GetWidgetJsonById(widgetId, templateType)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}
	if widgetExist.Id != widgetId {
		http.Error(wr, "widget Doesn't Exist.", 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	userId, err := getUserFromHeader(r)

	if err != nil {
		http.Error(wr, err.Error(), 401)
		return
	}

	widget := new(widget.Widget)
	err = json.Unmarshal(body, widget)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}
	err = ValidateWidgetData(widget)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}
	if widget.Id != widgetId {
		http.Error(wr, "ID mismatch, Param ID and ID in widget json doesn't match", 400)
		return
	}
	comments := gjson.Get(string(body), "comments").String()
	// to delete comment property from json, since it is not required in widget
	result, _ := sjson.Delete(string(body), "comments")
	body = []byte(result)
	err = w.widgetDAO.SaveWidgetAndAudit(body, widget, nil, comments, userId)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}
	wr.Write([]byte("successfully modified widget"))
}

func ValidateWidgetData(widget *widget.Widget) error {
	return nil
}

func getUserFromHeader(r *http.Request) (string, error) {
	userId := r.Header.Get("userId")
	if userId == "" {
		return "", errors.New("invalid user")
	}
	return userId, nil
}
