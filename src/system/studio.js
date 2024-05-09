Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) {
		can.db.favor = {}, can.run({}, [ctx.RUN, "web.code.system.favor"], function(_msg) {
			_msg.Table(function(value) { can.db.favor[value.path] = value })
			can.ui = can.onappend.layout(can), can.onimport._project(can, msg)
		})
	},
	_project: function(can, msg, target) {
		can.onimport.itemlist(can, msg.Table(function(value) { value.icon = can.base.endWith(value.path, nfs.PS)? icon.path: icon.file
			value.nick = can.core.Split(value.path, nfs.PS).pop()+` <span style="color:var(--disable-fg-color)">(${value.size})</span>`+(can.db.favor[value.path]? " <span style='color:red'>●</span>": "")
			value._select = (can.db.hash[0]||"").indexOf(value.path) == 0
			return value
		}), function(event, value, show, target) {
			show == undefined && can.run(event, [value.path], function(msg) {
				can.base.endWith(value.path, nfs.PS)? can.onimport._project(can, msg, target): can.onimport._content(can, msg, value, target)
			})
		}, function() {}, target)
	},
	_content: function(can, msg, value, target) { can.onimport.tabsCache(can, msg, value.path, value, target) },
})