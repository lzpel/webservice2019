# -*- coding: utf-8 -*-
#目的　HTMLとJSを書く
import webapp2,json
#追加
from google.appengine.ext.webapp import template
from google.appengine.ext import blobstore, ndb

class modelbase(ndb.Model):
	# 分類
	born = ndb.DateTimeProperty(auto_now_add=True)
	# 文章検索用
	name = ndb.StringProperty()
	text = ndb.TextProperty()
class comment(modelbase):
	user = ndb.KeyProperty()
class user(modelbase):
	email = ndb.StringProperty()
	password = ndb.StringProperty()
	blob = ndb.BlobProperty()

def hello(request):
	users=user.query().fetch()
	return webapp2.Response(template.render("index.html", {"user":users}))
def signup(request):
	return webapp2.Response(template.render("signup.html", None))
def signuppost(request):
	email=request.get("emailaddress")
	name=request.get("name")
	if user.query(user.email==email).get():
		return webapp2.Response(template.render("signuppost.html", {"success":False}))
	else:
		user(email=email,password=request.get("password"),name=name).put()
		return webapp2.Response(template.render("signuppost.html", {"success":True,"name":name}))
app = webapp2.WSGIApplication([('/', hello),('/signup',signup),('/signuppost',signuppost)])