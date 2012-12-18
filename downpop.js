
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

