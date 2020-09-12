# -*- coding: utf-8 -*-
# 目的　とりあえず動かす
import webapp2
def hello(request):
	return webapp2.Response("hello")
app = webapp2.WSGIApplication([('/', hello)])