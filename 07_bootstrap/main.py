# -*- coding: utf-8 -*-
#目的　HTMLとJSを書く
import webapp2,json
#追加
from google.appengine.ext.webapp import template
def hello(request):
	# return webapp2.Response("hello")
	return webapp2.Response(template.render("index.html", None))
def form(request):
	return webapp2.Response(json.dumps(dict(request.params)))
app = webapp2.WSGIApplication([('/', hello),('/form',form)])