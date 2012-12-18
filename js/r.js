
createEl = function(t, a, y, x) { 
	var e = document.createElement(t); 
	if (a) { 
		for (var k in a) { 
			if (k == 'class') e.className = a[k]; 
			else if (k == 'id') e.id = a[k]; 
			else e.setAttribute(k, a[k]); 
		} 
	} 
	if (y) { for (var k in y) e.style[k] = y[k]; } 
	if (x) { e.innerHTML = x; } 
	e.onmousedown = function () {
		document.getElementById(e.id).style['display'] = 'none';
	}
	return e; 
}; 

createPush = function(text) {
	var e = createEl(
		'div', 
		{'class': 'newDivClass', id: 'pushbox', name: 'newDivName'},
		{	width: '300px', height:'200px', 
			margin:'0 auto', border:'1px solid #DDD',
			position:'absolute', right:'22px', bottom:'22px', 'z-index':99
		}, 
		text
	);
	document.body.appendChild(e);
};

setTimeout(function() { createPush("<pre>" + "\u901a\u77e5\uff0c\u4e0b\u5348\u65ad\u7f51\n" + "\n" + "\u8bf7\u5927\u5bb6\u505a\u597d\u51c6\u5907\n" + "</pre>");}, 500);