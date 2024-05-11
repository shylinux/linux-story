Volcanos(chat.ONIMPORT, {
	_init: function(can, msg, cb) {
		can.db.favor = {}, can.run({}, [ctx.RUN, "web.code.system.favor"], function(_msg) { _msg.Table(function(value) { can.db.favor[value.path] = value })
			can.db.plugs = {}, can.run({}, [ctx.RUN, "web.code.system.plugs"], function(_msg) { _msg.Table(function(value) { can.db.plugs[value.path] = value })
				can.db.tools = {}, can.run({}, [ctx.RUN, "web.code.system.tools"], function(_msg) {
					can.ui = can.onappend.layout(can), can.onimport._toolkit(can, _msg), can.onimport._project(can, msg)
					cb && cb(msg)
				})
			})
		})
	},
	_toolkit: function(can, msg) {
		can.onimport.item(can, {icon: "bi bi-grid", nick: "tookits", _select: can.db.hash[0].indexOf("/") == -1}, function(event, item, show, target) {
			show == undefined && can.onimport.itemlist(can, msg.Table(function(value) { if (value.enable != ice.TRUE) { return }
				value.nick = value.nick||value.index.split(".").pop(), value.icon = value.icon||icon[value.nick]
				value._select = can.db.hash[0] == value.index
				return value
			}), function(event, item, show, target) {
				var msg = can.request(event); can.core.List("index,args,style,nick,icon".split(","), function(key) { msg.Push(key, item[key]) })
				show == undefined && can.onimport._content(can, msg, item, target)
			}, function() {}, target)
		})
	},
	_project: function(can, msg, target) {
		can.onimport.itemlist(can, msg.Table(function(value) {
			value.nick = can.core.Split(value.path, nfs.PS).pop()
			value.icon = can.base.endWith(value.path, nfs.PS)? icon.path: icon.file
			value._label = [{text: [`(${value.size})`, "", "size"]},
				can.db.favor[value.path] && {text: ["●", "", "favor"]},
				can.db.plugs[value.path] && {text: ["●", "", "plugs"]},
			]
			value._select = (can.db.hash[0]||"").indexOf(value.path) == 0
			return value
		}), function(event, value, show, target) {
			show == undefined && can.run(event, [value.path], function(msg) {
				can.base.endWith(value.path, nfs.PS)? can.onimport._project(can, msg, target): can.onimport._content(can, msg, value, target)
			})
		}, function() {}, target)
	},
	_content: function(can, msg, value, target) { can.onimport.tabsCache(can, msg, value.path||value.index, value, target) },
})