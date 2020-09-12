# -*- coding: utf-8 -*-
#目的　HTMLとJSを書く
import webapp2
#追加
from google.appengine.ext.webapp import template
def hello(request):
	# return webapp2.Response("hello")
	return webapp2.Response(template.render("index.html", None))
app = webapp2.WSGIApplication([('/', hello)])