
$(document).ready(function() {
	$('.datepicker').datepicker();
	if (window.location.search.match(/saveok/)) {
		alert('保存成功！');
	}
	if (window.location.search.match(/loginfail/)) {
		alert('登录失败！');
		window.location.href = "/login";
	}
});
