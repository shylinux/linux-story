Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) {
		can.db.plugs = {}, can.run({}, [ctx.RUN, "web.code.system.plugs"], function(_msg) { _msg.Table(function(value) { can.db.plugs[value.path] = value })
			can.db.favor = {}, can.run({}, [ctx.RUN, "web.code.system.favor"], function(_msg) { _msg.Table(function(value) { can.db.favor[value.path] = value })
				can.ui = can.onappend.layout(can), can.onimport._toolkit(can), can.onimport._project(can, msg)
			})
		})
	},
	_toolkit: function(can) {
		can.onimport.item(can, {icon: "bi bi-tools", nick: "tookits", _select: can.db.hash[0]&&can.db.hash[0].indexOf("/") == -1}, function(event, item, show, target) {
			show == undefined && can.onimport.itemlist(can, can.core.List([
				{index: "web.code.system.whois", icon: "bi bi-pin-map"},
				{index: "web.code.system.port", icon: "bi bi-bricks"},
				{index: "web.code.system.proc", icon: "bi bi-box"},
				{index: "web.code.system.service", icon: "bi bi-stack"},
				{index: "web.code.system.yum", icon: "bi bi-arrow-down-square"},
				{index: "web.code.system.disk", icon: "bi bi-disc"},
				{index: "web.code.system.user", icon: "bi bi-person-gear"},
				{index: "web.code.system.group", icon: "bi bi-people"},
				{index: "web.code.system.favor", icon: "bi bi-star"},
				{index: "web.code.system.unicode", icon: "bi bi-sort-alpha-down"},
				{index: "web.code.xterm", args: "sh", style: html.OUTPUT, icon: "bi bi-terminal"},
			], function(value) { value.nick = value.nick||value.index.split(".").pop()
				value._select = can.db.hash[0] == value.index
				return value
			}), function(event, item, show, target) {
				var msg = can.request(event); msg.Push({index: item.index, args: item.args, style: item.style, nick: item.nick, icon: item.icon})
				show == undefined && can.onimport._content(can, msg, item, target)
			}, function() {}, target)
		})
	},
	_project: function(can, msg, target) {
		can.onimport.itemlist(can, msg.Table(function(value) { value.icon = can.base.endWith(value.path, nfs.PS)? icon.path: icon.file
			value.nick = can.core.Split(value.path, nfs.PS).pop()+` <span style="color:var(--disable-fg-color)">(${value.size})</span>`
				+(can.db.plugs[value.path]? " <span style='color:blue'>●</span>": "")
				+(can.db.favor[value.path]? " <span style='color:red'>●</span>": "")
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