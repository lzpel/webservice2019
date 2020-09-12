# -*- coding: utf-8 -*-
import webapp2
#追加
from google.appengine.ext.webapp import template
def hello(request):
	# 変更 return webapp2.Response("hello")
	return webapp2.Response(template.render("index.html", None))
app = webapp2.WSGIApplication([('/', hello)])