# -*- coding: utf-8 -*-
#目的　HTMLとJSを書く
import webapp2,time
#追加
from google.appengine.ext.webapp import template
from google.appengine.ext import blobstore, ndb

class candidate(ndb.Model):
	born = ndb.DateTimeProperty(auto_now_add=True)
	vote = ndb.IntegerProperty()
	name = ndb.StringProperty()

def index(request):
	keyword=request.get("keyword")
	if keyword:
		candidatelist=candidate.query(candidate.name==keyword)
	else:
		candidatelist=candidate.query()
	return webapp2.Response(template.render("index.html", {"candidate":candidatelist.order(-candidate.vote).fetch(),"keyword":keyword}))
def command(request):
	id=int(request.get("id") or 0) #idを取得
	action=request.get("action") #指定された動作を文字列で受けとる
	if None:
		pass # 以下で動作を分岐
	elif action == "upvote": #得票数を増やす
		m=candidate.get_by_id(id) #idからエンティティを取得
		m.vote=(m.vote or 0)+1
		m.put()
	elif action == "downvote": #得票数を減らす
		m=candidate.get_by_id(id) #idからエンティティを取得
		m.vote=(m.vote or 0)-1
		m.put()
		if m.vote < -10:
			m.key.delete()
	elif action == "edit": #編集
		m=(id and candidate.get_by_id(id)) or candidate() #idからエンティティを取得
		m.populate(vote=int(request.get("vote") or 0),name=request.get("name"))
		m.put()
	elif action == "delete":
		candidate.get_by_id(id).key.delete()
	time.sleep(0.5)
	return webapp2.redirect("/")
app = webapp2.WSGIApplication([('/', index),('/command',command)])