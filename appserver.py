# -*- coding: utf-8 -*-
import os,sys
for w in os.environ.get("PATH").split(";"):
	for v in w.split(":"):
		if v.find("google-cloud-sdk")>0 and v.find("bin"):
			sys.path.append(v)
			execfile(os.path.join(v,"dev_appserver.py"))
			break