
$(document).ready(function() {
	$('.datepicker').datepicker();
	if (window.location.search.match(/saveok/)) {
		alert('保存成功！');
	}
});
