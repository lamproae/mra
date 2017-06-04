var next = 1;
$(".add-more").click(function(e){
	e.preventDefault();
	var addto = "#field" + next;
	next = next + 1;
	var newIn = '<br /><br /><input autocomplete="off" class="span3" id="field' + next + '" name="field' + next + '" type="text" data-provide="typeahead" data-items="8">';
	var newInput = $(newIn);
	$(addto).after(newInput);
	$("#field" + next).attr('data-source',$(addto).attr('data-source'));
	$("#count").val(next);  
});

$("#remove").click(function(){
	$("#table").remove();
});

$("#add").click(function(){
	$("#table").clone().appendTo("#div1");
});

var i=1;
$("#add_row").click(function(){
	$('#addr'+i).html("<td align = 'center'><img src = 'http://www.placehold.it/200x200' alt = '...'></img></td><td><input name='emotion"+i+"' type='text' placeholder='Emotion' class='form-control input-md'/> </td><td><input  name='source"+i+"' type='text' placeholder='Source'  class='form-control input-md'></td>");

	$('#tab_logic').append('<tr id="addr'+(i+1)+'"></tr>');
	i++; 
});

$("#delete_row").click(function(){
	if(i>1){
		$("#addr"+(i-1)).html('');
		i--;
	}
});

var i=1;
$("#add_row1").click(function(){
	$('#addr'+i).html("<td>"+ (i+1) +"</td><td><input name='name"+i+"' type='text' placeholder='Name' class='form-control input-md'  /> </td><td><input  name='mail"+i+"' type='text' placeholder='Mail'  class='form-control input-md'></td><td><input  name='mobile"+i+"' type='text' placeholder='Mobile'  class='form-control input-md'></td>");

	$('#tab_logic').append('<tr id="addr'+(i+1)+'"></tr>');
	i++; 
});
$("#delete_row1").click(function(){
	if(i>1){
		$("#addr"+(i-1)).html('');
		i--;
	}
});

