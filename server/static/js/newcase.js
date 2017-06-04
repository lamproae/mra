var d=1;
$("#add_dut").click(function(){
	$('#row'+d).html("<td>"+ (d+1) +"</td><td><input name='dut"+d+"' type='text' placeholder='DUT1' class='form-control input-md'  /> </td><td><input  name='device"+d+"' type='text' placeholder='V8500_SFU'  class='form-control input-md'></td>");

	$('#dut_tab').append('<tr id="row'+(d+1)+'"></tr>');
	d++; 
});
$("#del_dut").click(function(){
	if(d>1){
		$("#row"+(d-1)).html('');
		d--;
	}
});
