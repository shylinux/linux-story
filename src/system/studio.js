Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) { can.ui = can.onappend.layout(can), can.onimport._project(can, msg) },
	_project: function(can, msg, target) {
		can.onimport.itemlist(can, msg.Table(function(value) {
			value.icon = can.base.endWith(value.path, nfs.PS)? icon.path: icon.file
			value.nick = can.core.Split(value.path, nfs.PS).pop()+` <span style="color:var(--disable-fg-color)">(${value.size})</span>`
			value._select = (can.db.hash[0]||"").indexOf(value.path) == 0
			return value
		}), function(event, value, show, target) {
			show == undefined && can.run(event, [value.path], function(msg) {
				can.base.endWith(value.path, nfs.PS)? can.onimport._project(can, msg, target): can.onimport._content(can, msg, target, value)
			})
		}, function() {}, target)
	},
	_content: function(can, msg, target, value) { if (value._tabs) { return value._tabs.click() }
		value._tabs = can.onimport.tabs(can, [value], function() { can.db.current = value, can.onexport.hash(can, [value.path]), can.Status(value), target.click()
			if (can.onmotion.cache(can, function(save, load) {
				save({
					content_plugins: can.ui._content_plugins,
					display_plugins: can.ui._display_plugins,
				})
				can.ui._content_plugins = []
				can.ui._display_plugins = []
				load(value.path, function(bak) {
					can.ui._content_plugins = bak.content_plugins||[]
					can.ui._display_plugins = bak.display_plugins||[]
				}); return value.path
			}, can.ui.content, can.ui.display)) { return can.onimport.layout(can) }
			if (msg.Append(ctx.INDEX)) {
				msg.Table(function(value, index) {
					index == 0 && can.onappend.plugin(can, value, function(sub) {
						can.ui._content_plugins.push(sub), can.onimport.layout(can)
					}, can.ui.content)
					index == 1 && can.onappend.plugin(can, value, function(sub) {
						can.ui._display_plugins.push(sub), can.onimport.layout(can)
					}, can.ui.display)
				})
			} else {
				can.onappend.table(can, msg, null, can.ui.content)
			}
		}, function() { delete(value._tabs), can.onmotion.cacheClear(can, value.path) })
	},
	layout: function(can) {
		can.ui.display && can.onmotion.toggle(can, can.ui.display, can.ui._display_plugins && can.ui._display_plugins.length > 0)
		can.ui.layout(can.ConfHeight(), can.ConfWidth(), 0, function(height, width) {
			can.ui._display_plugins && can.ui._display_plugins.length > 0 && (height = can.ConfHeight()/2-1)
			can.core.List(can.ui._content_plugins, function(sub) { sub.onimport.size(sub, height, width, false) })
			can.core.List(can.ui._display_plugins, function(sub) { sub.onimport.size(sub, height, width, false) })
			can.page.style(can, can.ui.content, html.HEIGHT, height)
		})
	},
})
Volcanos(chat.ONACTION, {
	// upload: function(event, can) { can.request(event, {path: can.db.current.path}), can.user.upload(event, can) },
})